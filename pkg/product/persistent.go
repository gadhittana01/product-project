package product

import (
	"context"
	"database/sql"
	"log"
)

var (
	getAllProduct = `
		SELECT 
			id,
			product_name,
			qty,
			selling_price,
			promo_price
		FROM source_product
	`

	updateProduct = `
		UPDATE 
			destination_product
		SET qty = $1,
		selling_price = $2,
		promo_price = $3
		WHERE id = $4
	`
)

type persistent interface {
	updateAllProduct(ctx context.Context) error
}

type psql struct {
	dbSource *sql.DB
	dbDest   *sql.DB
}

func newPersistent(dbSource *sql.DB, dbDest *sql.DB) persistent {
	return psql{
		dbSource: dbSource,
		dbDest:   dbDest,
	}
}

func (p psql) updateAllProduct(ctx context.Context) error {
	res, err := p.dbSource.Query(getAllProduct)

	if err != nil {
		log.Fatal(err)
	}

	type Product struct {
		ID           int
		ProductName  string
		Qty          int
		SellingPrice float64
		PromoPrice   float64
	}

	for res.Next() {
		var product Product
		if err := res.Scan(&product.ID, &product.ProductName, &product.Qty,
			&product.SellingPrice, &product.PromoPrice); err != nil {
			return err
		}

		if _, err = p.dbDest.Exec(updateProduct, product.Qty, product.SellingPrice,
			product.PromoPrice, product.ID); err != nil {
			return err
		}
	}
	return nil
}
