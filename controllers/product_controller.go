package controllers

import (
	"net/http"

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
