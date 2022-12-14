package nsq

import (
	"context"

	"github.com/nsqio/go-nsq"
)

type productHandler struct {
	service ProductService
}

func newProductHandler(service ProductService) *productHandler {
	return &productHandler{
		service: service,
	}
}

func (p *productHandler) UpdateAllProduct(message *nsq.Message) error {
	err := p.service.UpdateAllProduct(context.Background())
	if err != nil {
		message.Requeue(-1)
		return err
	}
	message.Finish()
	return nil
}
