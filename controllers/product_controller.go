package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/jocelynh1110/http-playground/models"
	"github.com/jocelynh1110/http-playground/templates"
)

type productIndexAssigns struct {
	Products []models.Product
}

func ProductIndex(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, _ := models.ListProducts(db)
		// TODO: handle errors
		templates.RenderTemplate(w, templates.IndexTemplate, productIndexAssigns{
			Products: products,
		})
	}
}

type showProductAssigns struct {
	Product *models.Product
}

func ShowProduct(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, _ := strconv.ParseInt(id, 10, 64)
		product, _ := models.GetProduct(db, idInt)
		if product == nil {
			templates.HandleNotFound(w)
			return
		}
		templates.RenderTemplate(w, templates.ShowProductTemplate, showProductAssigns{
			Product: product,
		})
	}
}
