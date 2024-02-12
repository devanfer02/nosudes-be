package controller

import (
	"github.com/devanfer02/nosudes-be/domain"

	"github.com/gin-gonic/gin"
)

type authController struct {
	userSvc domain.UserService
}

func NewAuthController(userSvc domain.UserService) *authController {
	return &authController{
		userSvc,
	}
}

func (c *authController) RegisterUser(ctx *gin.Context) {

}

func (c *authController) LoginUser(ctx *gin.Context) {

}

func (c *authController) LogoutUser(ctx *gin.Context) {
	
}