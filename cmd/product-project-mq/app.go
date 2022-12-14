package main

import (
	"log"

	"gihub.com/gadhittana01/product-project/config"
	"gihub.com/gadhittana01/product-project/db"
	nsqd "gihub.com/gadhittana01/product-project/handler/nsq"
	"gihub.com/gadhittana01/product-project/pkg/product"
	"gihub.com/gadhittana01/product-project/services"
	"github.com/nsqio/go-nsq"
)

func initApp(c *config.GlobalConfig) error {
	config := nsq.NewConfig()

	//Creating the Producer using NSQD Address
	publisher, err := nsq.NewProducer(c.NSQ.Address, config)
	if err != nil {
		log.Fatal(err)
	}
	productPkg := product.New(db.InitDB("source"), db.InitDB("destination"))

	ps, err := services.NewProductService(services.ProductDependencies{
		Publisher: publisher,
		PR:        productPkg,
	})
	if err != nil {
		return err
	}

	handler, err := nsqd.NewHandler(nsqd.HandlerDependency{
		Config: c,
		PS:     ps,
	})
	if err != nil {
		return err
	}

	return serverConsumers(MQServerDependency{
		c:       c,
		handler: handler,
	})
}
