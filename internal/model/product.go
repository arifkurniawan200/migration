package model

import "time"

// Source struct for table source_product
type Source struct {
	ID           int       `json:"id" db:"id"`
	ProductName  string    `json:"product_name" db:"product_name"`
	Qty          int       `json:"qty" db:"qty"`
	SellingPrice float64   `json:"selling_price" db:"selling_price"`
	PromoPrice   float64   `json:"promo_price" db:"promo_price"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Destination struct for table destination_product
type Destination struct {
	ID           int       `json:"id" db:"id"`
	ProductName  string    `json:"product_name" db:"product_name"`
	Qty          int       `json:"qty" db:"qty"`
	SellingPrice float64   `json:"selling_price" db:"selling_price"`
	PromoPrice   float64   `json:"promo_price" db:"promo_price"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
