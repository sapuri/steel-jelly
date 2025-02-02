package getlinks

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
	eroterest_errs "github.com/sapuri/steel-jelly/steeljelly/eroterest/errors"
	eroterest_types "github.com/sapuri/steel-jelly/steeljelly/eroterest/types"
)

type GetLinksInteractor struct {
	OutputFilePath string
	PageNum        int
}

func NewGetLinksInteractor(outputFilePath string, pageNum int) *GetLinksInteractor {
	return &GetLinksInteractor{
		OutputFilePath: outputFilePath,
		PageNum:        pageNum,
	}
}

func (it *GetLinksInteractor) Invoke() error {
	const baseURL = "https://movie.eroterest.net/site/"
	var output []*eroterest_types.Link

	for i := 1; i <= it.PageNum; i++ {
		targetURL := fmt.Sprintf("%s?site_name=&page=%d", baseURL, i)
		fmt.Println(targetURL)

		links, err := it.scrape(targetURL)
		if err != nil {
			if err == eroterest_errs.ErrPageNotFound {
				break
			}
			return err
		}
		output = append(output, links...)

		time.Sleep(time.Second)
	}

	if err := it.export(it.OutputFilePath, output); err != nil {
		return err
	}

	fmt.Println("exported > " + it.OutputFilePath)
	return nil
}

func (it *GetLinksInteractor) scrape(targetURL string) (links []*eroterest_types.Link, err error) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, targetURL, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			err = eroterest_errs.ErrPageNotFound
			return
		}
		err = fmt.Errorf("returned %d from targetURL", res.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	li := doc.Find(`ul[class="list"]`).First().Find("li")
	if li == nil {
		err = errors.New("could not find ul element")
		return
	}

	li.Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		link, exists := a.Attr("href")
		if exists {
			links = append(links, &eroterest_types.Link{
				Link:     link,
				SiteName: a.Text(),
			})
		}
	})

	return
}

func (it *GetLinksInteractor) export(filePath string, links []*eroterest_types.Link) error {
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer func() {
		_ = file.Close()
	}()

	return gocsv.MarshalFile(&links, file)
}
