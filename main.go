package main

import "github.com/udayangaac/start-and-go/boot"

func main() {

	app := boot.Application()
	app.Init()
	app.Start(HTTPServer)
	app.Stop()

}

func HTTPServer() {

}
