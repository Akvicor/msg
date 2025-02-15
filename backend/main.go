package main

import (
	"log"
	"msg/cmd/app"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	if err := app.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
