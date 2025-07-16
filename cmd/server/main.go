package main

import (
	"github.com/fasthttp/router"
	"github.com/mortal-coders/gitlab-mr-batch/internal/handler"
	"github.com/mortal-coders/gitlab-mr-batch/pkg/gitlab"
	"github.com/valyala/fasthttp"
	"log"

	"github.com/mortal-coders/gitlab-mr-batch/internal/config"
)

func main() {

	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Config load error: %v", err)
	}

	gitlabClient := gitlab.NewGitlabClient(config.Gitlab.Host, config.Gitlab.Token)

	log.Println(config)

	r := router.New()

	fs := &fasthttp.FS{
		Root:       "./web/",
		IndexNames: []string{},
		Compress:   true,
	}
	r.ANY("/{filepath:*}", fs.NewRequestHandler())

	homeHandler := handler.NewHomeHandler()
	settingsHandler := handler.NewSettingsHandler(gitlabClient)

	r.GET("/", homeHandler.Home)
	r.GET("/settings", settingsHandler.Settings)

	if err = fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
