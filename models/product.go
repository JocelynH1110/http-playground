package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Price      int32     `db:"price"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func ListProducts(db *sqlx.DB) ([]Product, error) {
	result := []Product{}
	err := db.Select(&result, "select id,name,price,inserted_at,updated_at from products order by 1")
	return result, err
}
