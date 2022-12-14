package resthttp

import (
	"context"
	"net/http"
	"time"
)

type productHandler struct {
	service ProductService
}

func newProductHandler(service ProductService) *productHandler {
	return &productHandler{
		service: service,
	}
}
func (p productHandler) UpdateAllProduct(w http.ResponseWriter, r *http.Request) {
	resp := newResponse(time.Now())

	if err := p.service.PublishUpdateAllProduct(context.Background()); err != nil {
		resp.setInternalServerError(err.Error(), w)
		return
	}

	resp.setOK(map[string]interface{}{
		"message": "success",
	}, w)
	return
}
