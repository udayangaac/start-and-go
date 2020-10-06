package main

import "github.com/udayangaac/start-and-go/boot"

func main() {

	boot.Application().Init()
	boot.Application().Start(HTTPServer)
	boot.Application().Stop()

}

func HTTPServer() {

}
