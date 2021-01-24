package getblogs

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
	"github.com/sapuri/steel-jelly/steeljelly/internal/eroterest/getlinks"
)

type Blog struct {
	Link     string `csv:"link"`
	SiteName string `csv:"site_name"`
}

type GetBlogsInteractor struct {
	inputFilePath  string
	outputFilePath string
}

func NewGetBlogsInteractor(inputFilePath, outputFilePath string) *GetBlogsInteractor {
	return &GetBlogsInteractor{
		inputFilePath:  inputFilePath,
		outputFilePath: outputFilePath,
	}
}

func (it *GetBlogsInteractor) Invoke() error {
	file, err := os.OpenFile(it.inputFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	var links []*getlinks.Link
	if err := gocsv.UnmarshalFile(file, &links); err != nil {
		return err
	}

	var blogs []*Blog
	for _, link := range links {
		fmt.Println(link.Link)
		blog, err := it.scrape(link.Link, link.SiteName)
		if err != nil {
			return err
		}
		blogs = append(blogs, blog)
		time.Sleep(time.Second)
	}

	if err := it.export(it.outputFilePath, blogs); err != nil {
		return err
	}

	fmt.Println("exported > " + it.outputFilePath)
	return nil
}

func (it *GetBlogsInteractor) scrape(targetURL string, siteName string) (blog *Blog, err error) {
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
			err = ErrPageNotFound
			return
		}
		err = fmt.Errorf("returned %d from targetURL", res.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	a := doc.Find(`a[class="btn btn-primary btn-lg btn-block"]`).First()
	link, exists := a.Attr("href")
	if !exists {
		err = errors.New("href not found")
		return
	}

	return &Blog{
		Link:     link,
		SiteName: siteName,
	}, nil
}

func (it *GetBlogsInteractor) export(filePath string, blogs []*Blog) error {
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer func() {
		_ = file.Close()
	}()

	return gocsv.MarshalFile(&blogs, file)
}
