package main

import (
	"dictionary/model"
	repository "dictionary/repository"
)

func main() {
	repository := repository.Repository{
		Database: make(map[string]model.Entry),
	}

	repository.Init()

	CLI := CLI{Repository: &repository}
	go CLI.Listen()

	server := Server{Repository: &repository}
	server.Listen()
}
