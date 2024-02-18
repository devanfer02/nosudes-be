package controller

import (
	"github.com/devanfer02/nosudes-be/domain"
	resp "github.com/devanfer02/nosudes-be/utils/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userSvc domain.UserService
}

func NewAuthController(userSvc domain.UserService) *AuthController {
	return &AuthController{
		userSvc,
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
	user := domain.UserLogin{}

	if bindFailed(ctx, &user) {
		return
	}

}

func (c *AuthController) LogoutUser(ctx *gin.Context) {

}
