package steeljelly_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sapuri/steel-jelly/steeljelly"
	"github.com/sapuri/steel-jelly/steeljelly/internal/pornhub"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		pornhubClient pornhub.Client
	}{
		"default": {},
		"with pornhub client": {
			pornhubClient: pornhub.NewMockPornhub(nil),
		},
	}

	for name, tc := range cases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			client := steeljelly.NewClient(steeljelly.WithPornhubClient(tc.pornhubClient))
			if client == nil {
				t.Errorf("client is nil")
			}
		})
	}
}

func TestGetThumbnailURLs(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		siteType steeljelly.SiteType
	}{
		"pornhub": {
			siteType: steeljelly.SiteTypePornhub,
		},
	}

	for name, tc := range cases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			const videoURL = "https://example.com/video/"

			ctrl := gomock.NewController(t)
			t.Cleanup(func() {
				ctrl.Finish()
			})

			pornhubClient := pornhub.NewMockPornhub(ctrl)
			pornhubClient.EXPECT().GetThumbnailURLs(videoURL).Return([]string{"https://example.com/img/1/"}, nil)

			client := steeljelly.NewClient(steeljelly.WithPornhubClient(pornhubClient))

			got, err := client.GetThumbnailURLs(tc.siteType, videoURL)
			if err != nil {
				t.Fatal(err)
			}

			if len(got) != 1 {
				t.Errorf("len(got) = %v, want 1", len(got))
			}
		})
	}
}
