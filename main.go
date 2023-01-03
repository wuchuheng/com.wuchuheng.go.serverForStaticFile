package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "sever",
		Usage: "Start server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Value:       3000,
				Usage:       "Server Listening Port Number",
				DefaultText: "3000",
				Aliases:     []string{"p"},
			},
		},
		Action: runServer,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runServer(cCtx *cli.Context) error {
	port := cCtx.Int("port")
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/", fs)

	log.Printf("Listening on :%d...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
