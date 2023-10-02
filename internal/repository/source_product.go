package repository

import (
	"database/sql"
	"migration/internal/model"
)

type SourceProduct interface {
	// GetProductSource return all data in table source
	GetProductSource() ([]model.Source, error)
}

type Handler struct {
	db *sql.DB
}

func (h Handler) GetProductSource() ([]model.Source, error) {
	var sources []model.Source // Create a slice to store the results

	rows, err := h.db.Query("SELECT id,product_name,qty,selling_price,promo_price,created_at,updated_at FROM source_product")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Close the rows when done to prevent resource leaks

	for rows.Next() {
		var source model.Source // Create a variable to store each row's data
		if err := rows.Scan(&source.ID, &source.ProductName,
			&source.Qty, &source.SellingPrice, &source.PromoPrice, &source.CreatedAt, &source.UpdatedAt); err != nil {
			return nil, err
		}
		sources = append(sources, source) // Append the row to the slice
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sources, nil
}

func NewSourceRepository(db *sql.DB) *Handler {
	return &Handler{db}
}
