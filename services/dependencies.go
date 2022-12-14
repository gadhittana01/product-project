package services

import "context"

type (
	ProductResource interface {
		UpdateAllProduct(ctx context.Context) error
	}
)
