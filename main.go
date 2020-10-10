package main

import (
	"github.com/udayangaac/start-and-go/boot"
	"github.com/udayangaac/start-and-go/lib/app"
)

func main() {
	app := app.Application()
	app.Init(boot.Init)
	app.Start()
	app.Stop()
}
