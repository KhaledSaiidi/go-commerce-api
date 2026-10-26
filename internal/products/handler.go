package products

import (
	"net/http"

	"github.com/KhaledSaiidi/go-commerce/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Call the service -> ListProducts()
	// 2. Return JSON in HTTP response
	products := struct {
		Products []string `json:"products"`
	}{}
	json.Write(w, http.StatusOK, products)
}
