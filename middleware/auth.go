package middleware

import (
	"net/http"
	"strings"
	"time"

	resp "github.com/devanfer02/nosudes-be/utils/response"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		token = strings.Replace(token, "Bearer : ", "", 1)

		claims, err := m.authSvc.VerifyToken(token)

		if err != nil {
			resp.SendResp(ctx, http.StatusUnauthorized, "unauthorized", nil, nil)
			ctx.Abort()
			return
		}

		id, exp, err := m.authSvc.GetIdAndExp(claims)

		if err != nil {
			resp.SendResp(ctx, http.StatusUnauthorized, "unauthorized", nil, err)
			ctx.Abort()
			return
		}

		if time.Now().Unix() >= exp {
			resp.SendResp(ctx, http.StatusUnauthorized, "unauthorized", nil, nil)
			ctx.Abort()
			return
		}

		ctx.Set("user", id)
		ctx.Next()
	}
}

func (m *Middleware) OptAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		token = strings.Replace(token, "Bearer : ", "", 1)

		claims, err := m.authSvc.VerifyToken(token)

		if err != nil {
			ctx.Set("user", "")
			ctx.Next()
			return
		}

		id, exp, err := m.authSvc.GetIdAndExp(claims)

		if err != nil {
			ctx.Set("user", "")
			ctx.Next()
			return
		}

		if time.Now().Unix() >= exp {
			ctx.Set("user", "")
			ctx.Next()
			return
		}

		ctx.Set("user", id)
		ctx.Next()
	}
}