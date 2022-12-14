package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

const TopicNSQUpdateAllProduct string = "product_update_all_product"

type ProductService interface {
	PublishUpdateAllProduct(ctx context.Context) error
	UpdateAllProduct(ctx context.Context) error
}

type productService struct {
	publisher *nsq.Producer
	pr        ProductResource
}

func NewProductService(dep ProductDependencies) (ProductService, error) {
	return &productService{
		publisher: dep.Publisher,
		pr:        dep.PR,
	}, nil
}

func (p productService) PublishUpdateAllProduct(ctx context.Context) error {
	type Message struct {
	}
	msg := Message{}

	//Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	// This part will bun in the backgrond
	err = p.publisher.Publish(TopicNSQUpdateAllProduct, payload)
	if err != nil {
		return err
	}
	return nil
}

func (p productService) UpdateAllProduct(ctx context.Context) error {
	// This part will bun in the backgrond
	err := p.pr.UpdateAllProduct(context.Background())
	if err != nil {
		return err
	}
	return nil
}
