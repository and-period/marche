package handler

import (
	"net/http"

	"github.com/and-period/furumaru/api/internal/gateway/user/v1/request"
	"github.com/and-period/furumaru/api/internal/gateway/user/v1/response"
	"github.com/and-period/furumaru/api/internal/gateway/user/v1/service"
	"github.com/and-period/furumaru/api/internal/gateway/util"
	"github.com/and-period/furumaru/api/internal/store"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func (h *handler) cartRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/carts")

	r.GET("", h.GetCart)
	r.POST("/-/items", h.AddCartItem)
	r.DELETE("/:number/items/:productId", h.RemoveCartItem)
}

func (h *handler) GetCart(ctx *gin.Context) {
	in := &store.GetCartInput{
		SessionID: h.getSessionID(ctx),
	}
	cart, err := h.store.GetCart(ctx, in)
	if err != nil {
		h.httpError(ctx, err)
		return
	}
	if cart == nil {
		h.notFound(ctx, errNotFoundCart)
		return
	}
	var (
		coordinators service.Coordinators
		products     service.Products
	)
	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		coordinators, err = h.multiGetCoordinators(ectx, cart.Baskets.CoordinatorID())
		return
	})
	eg.Go(func() (err error) {
		products, err = h.multiGetProducts(ectx, cart.Baskets.ProductIDs())
		return
	})
	if err := eg.Wait(); err != nil {
		h.httpError(ctx, err)
		return
	}
	res := &response.CartResponse{
		Carts:        service.NewCarts(cart).Response(),
		Coordinators: coordinators.Response(),
		Products:     products.Response(),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *handler) AddCartItem(ctx *gin.Context) {
	req := &request.AddCartItemRequest{}
	if err := ctx.BindJSON(req); err != nil {
		h.badRequest(ctx, err)
		return
	}
	in := &store.AddCartItemInput{
		SessionID: h.getSessionID(ctx),
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	if err := h.store.AddCartItem(ctx, in); err != nil {
		h.httpError(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *handler) RemoveCartItem(ctx *gin.Context) {
	boxNumber, err := util.GetParamInt64(ctx, "number")
	if err != nil {
		h.badRequest(ctx, err)
		return
	}
	in := &store.RemoveCartItemInput{
		SessionID: h.getSessionID(ctx),
		BoxNumber: boxNumber,
		ProductID: util.GetParam(ctx, "productId"),
	}
	if err := h.store.RemoveCartItem(ctx, in); err != nil {
		h.httpError(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
