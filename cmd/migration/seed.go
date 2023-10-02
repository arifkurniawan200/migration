package migration

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"migration/config"
	driver "migration/db"
	"time"
)

// temporaryData temporary struct for seeding data
type temporaryData struct {
	ID           int       `json:"id"`
	ProductName  string    `json:"product_name"`
	Qty          int       `json:"qty"`
	SellingPrice float64   `json:"selling_price"`
	PromoPrice   float64   `json:"promo_price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type paramInject struct {
	Data      []temporaryData
	TableName string
}

// SeedingData running seeding data into database
func SeedingData(cfg config.Config) error {
	var datas []temporaryData

	// func to generate random name
	var randomName = func(i int) string {
		names := []string{"tas", "buku", "bolpoint", "pensil", "penghapus"}
		return fmt.Sprintf("%s-%d", names[rand.Intn(len(names)-0)], i)
	}

	var injectDatabase = func(cfg config.Database, param paramInject) error {
		ctx := context.Background()
		db, err := driver.NewDatabase(cfg)
		if err != nil {
			return err
		}

		queryInsert := fmt.Sprintf(`INSERT INTO %s(id,product_name,qty,selling_price,promo_price,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)`, param.TableName)

		tx, err := db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if err != nil {
			return err
		}
		for _, data := range param.Data {
			if param.TableName == "source_product" {
				// get random between 1 - 100
				data.Qty = rand.Intn(100-1) + 1

				// get random between 100 - 1000
				data.PromoPrice = rand.Float64()*900.0 + 100.0

				// get random between 1000 - 5000
				data.SellingPrice = rand.Float64()*5000.0 + 1000.0
			}
			_, err = tx.Exec(queryInsert, data.ID, data.ProductName, data.Qty, data.SellingPrice, data.PromoPrice, data.CreatedAt, data.UpdatedAt)
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

		if err := db.Close(); err != nil {
			return fmt.Errorf("db migrate: failed to close DB: %v\n", err)
		}
		return nil
	}

	for i := 0; i < 500; i++ {
		now := time.Now()
		datas = append(datas, temporaryData{
			ID:          i + 1,
			ProductName: randomName(i + 1),
			CreatedAt:   now,
			UpdatedAt:   now,
		})
	}

	// inject source database
	err := injectDatabase(cfg.SrcDB, paramInject{
		TableName: "source_product",
		Data:      datas,
	})
	if err != nil {
		return err
	}

	// inject destination database
	err = injectDatabase(cfg.DesDB, paramInject{
		TableName: "destination_product",
		Data:      datas,
	})
	if err != nil {
		return err
	}

	log.Println("Success seeding data")
	return nil
}
