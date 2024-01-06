package handlers

import (
	"myapp/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl")
}
