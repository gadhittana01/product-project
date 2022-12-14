package main

import (
	"log"

	"gihub.com/gadhittana01/product-project/db"
	"gihub.com/gadhittana01/product-project/pkg/domain"
	"github.com/jaswdr/faker"
)

var (
	queryInsertSource = `
		INSERT INTO source_product (product_name, qty, selling_price, promo_price)
		VALUES ($1, $2, $3, $4)
	`

	queryInsertDestination = `
		INSERT INTO destination_product (product_name, qty, selling_price, promo_price)
		VALUES ($1, $2, $3, $4)
	`
)

func seed(n int) {
	fake := faker.New()
	connSource := db.InitDB("source")
	destinationSource := db.InitDB("destination")

	defer connSource.Close()
	defer destinationSource.Close()

	for i := 0; i < n; i++ {
		productName := fake.Beer().Name()
		qty := fake.RandomDigitNotNull()
		sellingPrice := fake.Payment().Faker.Float64(2, 1, 200)
		promoPrice := fake.Payment().Faker.Float64(2, 1, 200)

		productSource := domain.Product{
			ProductName:  productName,
			Qty:          qty,
			SellingPrice: sellingPrice,
			PromoPrice:   promoPrice,
		}

		productDestination := domain.Product{
			ProductName:  productName,
			Qty:          0,
			SellingPrice: 0,
			PromoPrice:   0,
		}

		if _, err := connSource.Exec(queryInsertSource,
			productSource.ProductName,
			productSource.Qty,
			productSource.SellingPrice,
			productSource.PromoPrice); err != nil {
			log.Fatal(err)
		}

		if _, err := destinationSource.Exec(queryInsertDestination,
			productDestination.ProductName,
			productDestination.Qty,
			productDestination.SellingPrice,
			productDestination.PromoPrice); err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	n := 500
	seed(n)
}
