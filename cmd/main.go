package main

import (
	"github.com/pterm/pterm"
	conf "p_s_cli/internal/config"
	handler2 "p_s_cli/pkg/handler"
	"p_s_cli/pkg/repository"
)

const url string = "http://localhost:8000"

func main() {

	err := conf.Load()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(url)
	handler := handler2.NewHandler(repo)

	go func() {
		if err := handler.Init(); err != nil {
			pterm.Error.Printfln("Error: %s", err.Error())
		}
	}()
}
