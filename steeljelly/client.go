package steeljelly

import (
	"errors"

	"github.com/sapuri/steel-jelly/steeljelly/internal/pornhub"
)

type Client interface {
	GetThumbnailURLs(siteType SiteType, videoURL string) ([]string, error)
}

type clientImpl struct {
	pornhub pornhub.Client
}

type ClientOptions func(*options)

type options struct {
	pornhubClient pornhub.Client
}

// WithPornhubClient sets Pornhub client as ClientOptions.
func WithPornhubClient(client pornhub.Client) ClientOptions {
	return func(o *options) {
		o.pornhubClient = client
	}
}

// NewClient returns a new Client.
func NewClient(opts ...ClientOptions) Client {
	var o options
	for _, fn := range opts {
		fn(&o)
	}

	var pornhubClient pornhub.Client
	if o.pornhubClient == nil {
		pornhubClient = pornhub.NewClient()
	} else {
		pornhubClient = o.pornhubClient
	}

	return &clientImpl{
		pornhub: pornhubClient,
	}
}

func (c *clientImpl) GetThumbnailURLs(siteType SiteType, videoURL string) ([]string, error) {
	switch siteType {
	case SiteTypePornhub:
		return c.pornhub.GetThumbnailURLs(videoURL)
	default:
		return nil, errors.New("invalid siteType")
	}
}
