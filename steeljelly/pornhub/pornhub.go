//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package pornhub

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Client interface {
	GetThumbnailURLs(videoURL string) ([]string, error)
}

type clientImpl struct{}

func NewClient() Client {
	return &clientImpl{}
}

func (c *clientImpl) GetThumbnailURLs(videoURL string) ([]string, error) {
	const imageNum = 16

	ret := make([]string, 0, imageNum)

	u, err := c.getThumbnailURL(videoURL)
	if err != nil {
		return nil, err
	}

	spl := strings.Split(u, ".")
	if len(spl) < 2 {
		return nil, errors.New("unexpected URL format")
	}
	ext := spl[len(spl)-1]

	spl = strings.Split(u, ")")
	if len(spl) < 2 {
		return nil, errors.New("unexpected URL format")
	}
	spl = spl[:len(spl)-1]

	for i := 1; i <= imageNum; i++ {
		ret = append(ret, fmt.Sprintf("%s)%d.%s", strings.Join(spl, ")"), i, ext))
	}

	return ret, nil
}

func (c *clientImpl) getThumbnailURL(videoURL string) (string, error) {
	res, err := http.Get(videoURL)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("returned %d from videoURL", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	imageURL, exists := doc.Find(`meta[property="og:image"]`).First().Attr("content")
	if !exists {
		return "", errors.New("could not find image URL")
	}

	return imageURL, nil
}
