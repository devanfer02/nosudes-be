package controller

import (
	"strings"
	"strconv"

	"github.com/devanfer02/nosudes-be/domain"
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

func getLocQuery(ctx *gin.Context) *domain.LocQuery {
	query := domain.LocQuery{}

	if ctx.Query("user_lat") == "" || ctx.Query("user_lng") == "" {
		return nil
	}

	userlat, err := strconv.ParseFloat(strings.TrimSpace(ctx.Query("user_lat")), 64)
	
	if err != nil {
		return nil
	}

	userlng, err := strconv.ParseFloat(strings.TrimSpace(ctx.Query("user_lng")), 64)
	
	if err != nil {
		return nil
	}

	query.UserLat = userlat 
	query.UserLng = userlng

	return &query
}
