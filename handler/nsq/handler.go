package nsq

import (
	"time"

	"gihub.com/gadhittana01/product-project/config"
	"github.com/nsqio/go-nsq"
)

type NSQhandler struct {
	Config  nsq.Config
	Topic   string
	Channel string
	Handler nsq.HandlerFunc
}

type IHandler interface {
	BuildConsumers() []NSQhandler
}

type Handler struct {
	config  *config.GlobalConfig
	product *productHandler
}

type HandlerDependency struct {
	Config *config.GlobalConfig
	PS     ProductService `validate:"nonnil"`
}

func NewHandler(dep HandlerDependency) (IHandler, error) {
	return &Handler{
		config:  dep.Config,
		product: newProductHandler(dep.PS),
	}, nil
}

func (h *Handler) BuildConsumers() []NSQhandler {
	return []NSQhandler{
		{
			Config: nsq.Config{
				MaxAttempts:         10,
				MaxInFlight:         5,
				MaxRequeueDelay:     time.Second * 900,
				DefaultRequeueDelay: time.Second * 0,
			},
			Channel: "channel_product_update_all_product",
			Topic:   "product_update_all_product",
			Handler: h.product.UpdateAllProduct,
		},
	}
}
