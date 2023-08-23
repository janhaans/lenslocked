package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/janhaans/lenslocked/controller"
	"github.com/janhaans/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controller.StaticHandler(views.Must(views.ParseFS("home.gohtml"))))
	r.Get("/contact", controller.StaticHandler(views.Must(views.ParseFS("contact.gohtml"))))
	r.Get("/faq", controller.StaticHandler(views.Must(views.ParseFS("faq.gohtml"))))
	r.Get("/galleries/{galleryID}", controller.ShowGalleryHandler(views.Must(views.ParseFS("galleries.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
