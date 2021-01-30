package cli

import (
	"github.com/sapuri/steel-jelly/steeljelly/internal/eroterest"
	"github.com/urfave/cli/v2"
)

func newEroterestCmd() *cli.Command {
	return &cli.Command{
		Name:  "eroterest",
		Usage: "エロタレスト用のコマンド",
		Subcommands: []*cli.Command{
			newGetBlogsCmd(),
			newGetLinksCmd(),
		},
	}
}

func newGetBlogsCmd() *cli.Command {
	return &cli.Command{
		Name:  "get-blogs",
		Usage: "ブログ一覧を取得します",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "links-csv",
				Aliases: []string{"l"},
				Usage:   "リンク一覧のCSV",
				Value:   "eroterest_links.csv",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "出力先ファイルパス",
				Value:   "eroterest_blogs.csv",
			},
		},
		Action: func(c *cli.Context) error {
			linksCSV := c.String("links-csv")
			blogsCSV := c.String("output")
			return eroterest.NewClient().GetBlogs(linksCSV, blogsCSV)
		},
	}
}

func newGetLinksCmd() *cli.Command {
	return &cli.Command{
		Name:  "get-links",
		Usage: "リンク一覧を取得します",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "出力先ファイルパス",
				Value:   "eroterest_links.csv",
			},
			&cli.IntFlag{
				Name:    "page-num",
				Aliases: []string{"n"},
				Usage:   "取得するページ数",
				Value:   10,
			},
		},
		Action: func(c *cli.Context) error {
			outputFilePath := c.String("output")
			pageNum := c.Int("page-num")
			return eroterest.NewClient().GetLinks(outputFilePath, pageNum)
		},
	}
}
