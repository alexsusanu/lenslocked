package main

import (
	"fmt"
	"github.com/alexsusanu/lenslocked/controllers"
	"github.com/alexsusanu/lenslocked/templates"
	"github.com/alexsusanu/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found x", http.StatusNotFound)
	})
	fmt.Println("Listening on port :3000")
	http.ListenAndServe(":3000", r)
}
