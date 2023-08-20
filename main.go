package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/janhaans/lenslocked/controller"
	"github.com/janhaans/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tmplHome := views.Must(views.ParseTemplate(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controller.StaticHandler(tmplHome))

	tmplContact := views.Must(views.ParseTemplate(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controller.StaticHandler(tmplContact))

	tmplFaq := views.Must(views.ParseTemplate(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controller.StaticHandler(tmplFaq))

	tmplGalleries := views.Must(views.ParseTemplate(filepath.Join("templates", "galleries.gohtml")))
	r.Get("/galleries/{galleryID}", controller.ShowGalleryHandler(tmplGalleries))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
