package handler

import (
	"net/http"

	"github.com/and-period/furumaru/api/internal/gateway/admin/v1/request"
	"github.com/and-period/furumaru/api/internal/gateway/admin/v1/response"
	"github.com/and-period/furumaru/api/internal/gateway/admin/v1/service"
	"github.com/and-period/furumaru/api/internal/gateway/util"
	"github.com/and-period/furumaru/api/internal/user"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) administratorRoutes(rg *gin.RouterGroup) {
	arg := rg.Use(h.authentication())
	arg.GET("", h.ListAdministrators)
	arg.POST("", h.CreateAdministrator)
	arg.GET("/:adminId", h.GetAdministrator)
}

func (h *apiV1Handler) ListAdministrators(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	const (
		defaultLimit  = 20
		defaultOffset = 0
	)

	limit, err := util.GetQueryInt64(ctx, "limit", defaultLimit)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	offset, err := util.GetQueryInt64(ctx, "offset", defaultOffset)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.ListAdministratorsInput{
		Limit:  limit,
		Offset: offset,
	}
	admins, err := h.user.ListAdministrators(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.AdministratorsResponse{
		Administrators: service.NewAdministrators(admins).Response(),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) GetAdministrator(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	in := &user.GetAdministratorInput{
		AdministratorID: util.GetParam(ctx, "adminId"),
	}
	admin, err := h.user.GetAdministrator(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.AdministratorResponse{
		Administrator: service.NewAdministrator(admin).Response(),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) CreateAdministrator(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateAdministratorRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.CreateAdministratorInput{
		Lastname:      req.Lastname,
		Firstname:     req.Firstname,
		LastnameKana:  req.LastnameKana,
		FirstnameKana: req.FirstnameKana,
		Email:         req.Email,
		PhoneNumber:   req.PhoneNumber,
	}
	admin, err := h.user.CreateAdministrator(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.AdministratorResponse{
		Administrator: service.NewAdministrator(admin).Response(),
	}
	ctx.JSON(http.StatusOK, res)
}
