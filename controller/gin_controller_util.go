package controller

import (
	resp "github.com/devanfer02/nosudes-be/utils/response"
	"github.com/gin-gonic/gin"
)

func bindFailed(ctx *gin.Context, data interface{}) bool {
	if err := ctx.ShouldBindJSON(data); err != nil {
		resp.SendResp(ctx, 400, "bad body request", nil, err)
		return true 
	}

	return false
}