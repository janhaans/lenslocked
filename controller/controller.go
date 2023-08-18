package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/janhaans/lenslocked/views"
)

func StaticHandler(tmpl *views.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func ShowGalleryHandler(tmpl *views.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		galleryID := chi.URLParam(r, "galleryID")
		tmpl.Execute(w, galleryID)
	}
}
