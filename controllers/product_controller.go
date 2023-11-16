package controllers

import (
	"fmt"
	"log"
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

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, templates.NewProductTemplate, nil)
}

func CreateProduct(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Bad Request")
			return
		}
		name := r.Form.Get("name")
		price := r.Form.Get("price")
		priceInt, _ := strconv.ParseInt(price, 10, 32)
		product, err := models.InsertProduct(db, name, int32(priceInt))
		if err != nil {
			log.Println(err)
		}
		log.Println(product)
		url := fmt.Sprintf("/products/%d", product.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}

}
