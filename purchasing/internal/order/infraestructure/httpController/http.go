package httpController

import (
	"context"
	"net/http"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/application/creator"
	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"github.com/AlfredoPastor/ddd-go/shared/ginhttp"
	"github.com/gin-gonic/gin"
)

type HttpController struct {
	ginhttp.HttpServer
	creator creator.OrderCreatorService
}

func NewHttpController(clientGin ginhttp.HttpServer, creator creator.OrderCreatorService) HttpController {
	server := HttpController{
		HttpServer: clientGin,
		creator:    creator,
	}
	server.registerRoutes()
	return server
}

func (h *HttpController) registerRoutes() {
	h.Engine.PUT("/v1/orders", h.createOrder())
}

func (h *HttpController) Run(ctx context.Context) error {
	return h.HttpServer.Run(ctx)
}

func (h *HttpController) createOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orderAdapter domain.OrderAdapter
		if err := ctx.ShouldBindJSON(&orderAdapter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		orderLines := []domain.OrderLine{}
		for _, orderLineAdapter := range orderAdapter.OrderLines {
			orderLine, err := domain.NewOrderLine(orderLineAdapter.ID, orderLineAdapter.ProductID, orderLineAdapter.Price, orderLineAdapter.Quantity)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			orderLines = append(orderLines, orderLine)
		}
		err := h.creator.Create(ctx, orderAdapter.ID, orderAdapter.ClientID, orderAdapter.AddressShipping, orderLines)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, nil)
	}
}
