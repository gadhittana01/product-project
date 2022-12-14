package resthttp

import "context"

type (
	ProductService interface {
		PublishUpdateAllProduct(ctx context.Context) error
	}
)
