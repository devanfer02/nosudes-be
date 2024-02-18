package controller

import (
	"github.com/devanfer02/nosudes-be/domain"
	resp "github.com/devanfer02/nosudes-be/utils/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc domain.UserService
}

func NewUserController(svc domain.UserService) *UserController {
	return &UserController{svc}
}

func(c *UserController) FetchAll(ctx *gin.Context) {
	users, err := c.svc.FetchAll(ctx.Request.Context())
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to fetch data", nil, err)
		return 
	}	

	resp.SendResp(ctx, code, "successfully fetch data", users, nil)
}

func(c *UserController) FetchByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	user, err := c.svc.FetchByID(ctx.Request.Context(), idParam)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to fetch data", nil, err)
		return 
	}

	resp.SendResp(ctx, code, "successfully fetch data by id", user, nil)
}

func(c *UserController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	user := domain.UserPayload{}

	if bindFailed(ctx, &user) {
		return 
	}

	user.ID = idParam

	err := c.svc.UpdateUser(ctx.Request.Context(), &user) 
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to update user", nil, err)
		return 
	}

	resp.SendResp(ctx, code, "successfully update user", user, nil)
}

func(c *UserController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	
	err := c.svc.DeleteUser(ctx.Request.Context(), idParam)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to delete user", nil, err)
		return 
	}

	resp.SendResp(ctx, code, "successfully delete user", nil, nil)
}