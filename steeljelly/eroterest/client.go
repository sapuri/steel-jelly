package eroterest

import (
	"github.com/sapuri/steel-jelly/steeljelly/eroterest/getblogs"
	"github.com/sapuri/steel-jelly/steeljelly/eroterest/getlinks"
)

type Client interface {
	GetBlogs(inputFilePath, outputFilePath string) error
	GetLinks(outputFilePath string, pageNum int) error
}

type clientImpl struct{}

func (c *clientImpl) GetBlogs(inputFilePath, outputFilePath string) error {
	return getblogs.NewGetBlogsInteractor(inputFilePath, outputFilePath).Invoke()
}

func (c *clientImpl) GetLinks(outputFilePath string, pageNum int) error {
	return getlinks.NewGetLinksInteractor(outputFilePath, pageNum).Invoke()
}

func NewClient() Client {
	return &clientImpl{}
}
