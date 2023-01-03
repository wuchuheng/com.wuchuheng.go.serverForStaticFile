package main

import (
	"embed"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/fs"
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

//go:embed assets/*
var content embed.FS

func runServer(cCtx *cli.Context) error {
	port := cCtx.Int("port")
	contentStatic, _ := fs.Sub(content, "assets")
	http.Handle("/", http.FileServer(http.FS(contentStatic)))

	log.Printf("Listening on :%d...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
