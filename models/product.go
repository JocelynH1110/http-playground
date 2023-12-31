package models

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const PRODUCT_COLUMNS = "id,name,price,inserted_at,updated_at"

type Product struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Price      int32     `db:"price"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func ListProducts(db *sqlx.DB) ([]Product, error) {
	result := []Product{}
	err := db.Select(&result, "select "+PRODUCT_COLUMNS+" from products order by 1")
	return result, err
}

func GetProduct(db *sqlx.DB, id int64) (*Product, error) {
	result := Product{}
	err := db.Get(&result, "select "+PRODUCT_COLUMNS+" from products where id=$1", id)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func InsertProduct(db *sqlx.DB, name string, price int32) (*Product, error) {
	result := Product{}
	err := db.QueryRowx("insert into products (name,price) values ($1,$2) returning "+PRODUCT_COLUMNS, name, price).StructScan(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &result, nil
}
