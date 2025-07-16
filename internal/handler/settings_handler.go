package handler

import (
	"fmt"
	"github.com/mortal-coders/gitlab-mr-batch/internal/model"
	"github.com/mortal-coders/gitlab-mr-batch/internal/service"
	"github.com/valyala/fasthttp"
	"html/template"
	"log"
)

type SettingsHandler struct {
	tmpl   *template.Template
	client service.GitlabClient
}

func NewSettingsHandler(client service.GitlabClient) *SettingsHandler {
	t := template.Must(template.ParseFiles("./web/templates/settings.html"))
	return &SettingsHandler{tmpl: t, client: client}
}

func (h *SettingsHandler) Settings(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("text/html; charset=utf-8")

	namespaces, err := h.client.ListNamespaces()
	if err != nil {
		ctx.Error(fmt.Sprintf("Failed to fetch namespaces: %v", err), fasthttp.StatusInternalServerError)
		log.Printf("error listing namespaces: %v", err)
		return
	}

	if len(namespaces) == 0 {
		ctx.Error("No namespaces found", fasthttp.StatusNotFound)
		return
	}
	data := Data{
		Namespaces: namespaces,
		Users:      make(map[int][]model.GitlabUser, len(namespaces)),
		Projects:   make(map[int][]model.GitlabProject, len(namespaces)),
	}

	for _, namespace := range namespaces {
		users, err := h.client.ListUsersInGroup(namespace.Id)
		if err != nil {
			log.Printf("error listing users: %v", err)
			continue
		}

		projects, err := h.client.ListProjectsInGroup(namespace.Id)
		if err != nil {
			log.Printf("error listing projects: %v", err)
			continue
		}
		data.Users[namespace.Id] = users
		data.Projects[namespace.Id] = projects
	}

	if err = h.tmpl.Execute(ctx.Response.BodyWriter(), data); err != nil {
		ctx.Error("Template error", fasthttp.StatusInternalServerError)
		log.Printf("error executing template: %v", err)
	}
}

type Data struct {
	Namespaces []model.GitlabNamespace
	Users      map[int][]model.GitlabUser
	Projects   map[int][]model.GitlabProject
}
