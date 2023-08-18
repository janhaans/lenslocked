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

	tmplHome, err := views.ParseTemplate(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controller.StaticHandler(tmplHome))

	tmplContact, err := views.ParseTemplate(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controller.StaticHandler(tmplContact))

	tmplFaq, err := views.ParseTemplate(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controller.StaticHandler(tmplFaq))

	tmplGalleries, err := views.ParseTemplate(filepath.Join("templates", "galleries.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/galleries/{galleryID}", controller.ShowGalleryHandler(tmplGalleries))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
