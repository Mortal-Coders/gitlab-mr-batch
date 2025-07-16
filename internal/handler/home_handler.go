package handler

import (
	"github.com/valyala/fasthttp"
	"html/template"
)

type HomeHandler struct {
	tmpl *template.Template
}

func NewHomeHandler() *HomeHandler {
	t := template.Must(template.ParseFiles("./web/templates/index.html"))
	return &HomeHandler{tmpl: t}
}

func (h *HomeHandler) Home(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("text/html; charset=utf-8")
	h.tmpl.Execute(ctx.Response.BodyWriter(), nil)
}
