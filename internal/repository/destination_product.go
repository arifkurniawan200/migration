package repository

import (
	"context"
	"database/sql"
	"migration/internal/model"
	"time"
)

type DestinationProduct interface {
	// UpdateProductDestinationTx method to update batch data from source to destination using transactional database
	UpdateProductDestinationTx(sources []model.Source) error
	// GetProductDestination return all data from table destination
	GetProductDestination() ([]model.Destination, error)
}

func NewDestinationRepository(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h Handler) UpdateProductDestinationTx(sources []model.Source) error {
	tx, err := h.db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}

	for _, source := range sources {
		now := time.Now()
		_, err = tx.Exec(`UPDATE destination_product SET qty = $1,
                               selling_price=$2,
                               promo_price=$3,
                               updated_at=$4 
                           where id = $5 and product_name = $6`,
			source.Qty, source.SellingPrice, source.PromoPrice, now, source.ID, source.ProductName)
		if err != nil {
			errx := tx.Rollback()
			if errx != nil {
				return errx
			}
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func (h Handler) GetProductDestination() ([]model.Destination, error) {
	var destinations []model.Destination // Create a slice to store the results

	rows, err := h.db.Query("SELECT id,product_name,qty,selling_price,promo_price,created_at,updated_at FROM destination_product")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Close the rows when done to prevent resource leaks

	for rows.Next() {
		var destination model.Destination // Create a variable to store each row's data
		if err := rows.Scan(&destination.ID, &destination.ProductName,
			&destination.Qty, &destination.SellingPrice, &destination.PromoPrice, &destination.CreatedAt, &destination.UpdatedAt); err != nil {
			return nil, err
		}
		destinations = append(destinations, destination) // Append the row to the slice
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return destinations, nil

}
