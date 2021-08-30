package main

import (
	"repository"
	"server"
)

func main() {
	s := new(Server)
	urlData := make(UrlsData)
	s.start("localhost:8080", &urlData)
}
