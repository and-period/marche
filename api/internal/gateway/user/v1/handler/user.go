package handler

import (
	"net/http"

	"github.com/and-period/marche/api/internal/gateway/user/v1/request"
	"github.com/and-period/marche/api/internal/gateway/user/v1/response"
	"github.com/and-period/marche/api/internal/gateway/util"
	"github.com/and-period/marche/api/proto/user"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) GetUserMe(ctx *gin.Context) {
	// TODO: 詳細の実装
	res := &response.UserMeResponse{}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) CreateUser(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateUserRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.CreateUserRequest{
		Email:                req.Email,
		PhoneNumber:          req.PhoneNumber,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}
	out, err := h.user.CreateUser(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.CreateUserResponse{
		ID: out.UserId,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) VerifyUser(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.VerifyUserRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.VerifyUserRequest{
		UserId:     req.ID,
		VerifyCode: req.VerifyCode,
	}
	_, err := h.user.VerifyUser(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) CreateUserWithOAuth(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	token, err := util.GetAuthToken(ctx)
	if err != nil {
		unauthorized(ctx, err)
		return
	}

	in := &user.CreateUserWithOAuthRequest{
		AccessToken: token,
	}
	out, err := h.user.CreateUserWithOAuth(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.UserMeResponse{
		ID:          out.User.Id,
		Email:       out.User.Email,
		PhoneNumber: out.User.PhoneNumber,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateUserEmail(ctx *gin.Context) {
	// TODO: 詳細の実装
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) VerifyUserEmail(ctx *gin.Context) {
	// TODO: 詳細の実装
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) UpdateUserPassword(ctx *gin.Context) {
	// TODO: 詳細の実装
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) ForgotUserPassword(ctx *gin.Context) {
	// TODO: 詳細の実装
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) ResetUserPassword(ctx *gin.Context) {
	// TODO: 詳細の実装
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) DeleteUser(ctx *gin.Context) {
	// TODO: 詳細の実装
	ctx.JSON(http.StatusNoContent, gin.H{})
}
