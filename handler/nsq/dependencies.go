package nsq

import "context"

type (
	ProductService interface {
		UpdateAllProduct(ctx context.Context) error
	}
)
