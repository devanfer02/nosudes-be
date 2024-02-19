package controller

import (
	"net/http"

	"github.com/devanfer02/nosudes-be/domain"
	resp "github.com/devanfer02/nosudes-be/utils/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userSvc domain.UserService
	authSvc domain.AuthService
}

func NewAuthController(userSvc domain.UserService, authSvc domain.AuthService) *AuthController {
	return &AuthController{
		userSvc,
		authSvc,
	}
}

func (c *AuthController) RegisterUser(ctx *gin.Context) {
	user := domain.UserPayload{}

	if bindFailed(ctx, &user) {
		return
	}

	user.Default()
	err := c.userSvc.InsertUser(ctx.Request.Context(), &user)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to register user", nil, err)
		return
	}

	resp.SendResp(ctx, code, "user successfully registered", nil, nil)
}

func (c *AuthController) LoginUser(ctx *gin.Context) {
	userPayload := domain.UserLogin{}

	if bindFailed(ctx, &userPayload) {
		return
	}

	user, err := c.userSvc.FetchByEmail(ctx.Request.Context(), userPayload.Email)

	if err != nil {
		resp.SendResp(ctx, http.StatusUnauthorized, "wrong email or password", nil, nil)
	}

	if !user.Compare(userPayload.Password) {
		resp.SendResp(ctx, http.StatusUnauthorized, "wrong email or password", nil, nil)
	}

	token, err := c.authSvc.CreateAccessToken(user.ID, user.Fullname)

	if err != nil {
		resp.SendResp(ctx, 500, "internal server error", nil, err)
	}

	resp.SendResp(ctx, 200, "user successfully login", gin.H{"token": token}, nil)
}