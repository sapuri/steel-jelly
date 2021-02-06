package pornhub

import (
	"fmt"

	"github.com/sapuri/steel-jelly/steeljelly/pornhub"
	"github.com/urfave/cli/v2"
)

func NewPornhubCmd() *cli.Command {
	return &cli.Command{
		Name:  "pornhub",
		Usage: "Pornhub",
		Subcommands: []*cli.Command{
			newGetThumbnailsCmd(),
		},
	}
}

func newGetThumbnailsCmd() *cli.Command {
	return &cli.Command{
		Name:  "get-thumbnails",
		Usage: "Retrieve thumbnail URLs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "video-url",
				Aliases:  []string{"url"},
				Usage:    "The URL of the video",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			videoURL := c.String("video-url")
			res, err := pornhub.NewClient().GetThumbnailURLs(videoURL)
			for _, u := range res {
				fmt.Println(u)
			}
			return err
		},
	}
}
