package main

import (
	"os"

	"https://github.com/mizuki0414/09_myqi/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "qiitasearch"
	app.Usage = "search qiita articles"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:  "article",
			Usage: "qiita + mine : you get yours articles",
			Action: func(c *cli.Context) error {
				qiitaToken := os.Getenv("QIITA_TOKEN")
				datas := cmd.FetchQiitaData(qiitaToken)
				cmd.OutputQiitaData(datas)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
