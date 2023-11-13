package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/jocelynh1110/http-playground/controllers"
	"github.com/jocelynh1110/http-playground/templates"
	_ "github.com/lib/pq"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, templates.IndexTemplate, nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>聯繫我們</h1>")
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>關於我</h1>")
}

func main() {
	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@127.0.0.1:5432/catalog_dev")
	if err != nil {
		log.Fatalln(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controllers.ProductIndex(db))
	r.Get("/contact", ContactHandler)
	r.Get("/about", AboutUs)
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r) //ListenAndServe裡有個無限迴圈，所以會一直跑到程式結束為止
}
