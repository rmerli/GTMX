// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	ID        int64
	Name      string
	SectionID pgtype.Int8
}

type Product struct {
	ID         int64
	Name       string
	Abbr       string
	Price      pgtype.Numeric
	CategoryID pgtype.Int8
}

type ProductsVariant struct {
	ProductID int64
	VariantID int64
}

type Section struct {
	ID   int64
	Name string
}

type Variant struct {
	ID    int64
	Name  string
	Price pgtype.Numeric
}
