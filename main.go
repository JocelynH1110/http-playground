package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "aloha")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>聯繫我們</h1>")
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>關於我</h1>")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", HomeHandler)
	r.Get("/contact", ContactHandler)
	r.Get("/about", AboutUs)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r) //ListenAndServe裡有個無限迴圈，所以會一直跑到程式結束為止
}
