package main

import (
	"context"
	"github.com/pterm/pterm"
	conf "p_s_cli/internal/config"
	handler2 "p_s_cli/pkg/handler"
	"p_s_cli/pkg/repository"
)

func main() {
	ctx := context.Background()
	cfg := conf.Load()
	err := conf.LoadToken()
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(ctx, cfg.Url, cfg.AuthUrl, cfg.TimeOut, cfg.Retries, cfg.AppId)
	handler := handler2.NewHandler(repo)

	if err := handler.Init(ctx); err != nil {
		pterm.Error.Printfln("Error: %s", err.Error())
	}
}
