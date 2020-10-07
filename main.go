package main

import (
	"github.com/udayangaac/start-and-go/boot"
	"github.com/udayangaac/start-and-go/transport/http/server"
)

func main() {
	app := boot.Application()
	app.Init()
	app.Start(HTTPServer)
	app.Stop()

}

func HTTPServer() {
	server.Init()
}
