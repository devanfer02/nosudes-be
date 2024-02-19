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
			resp.SendResp(ctx, http.StatusUnauthorized, "invalid email or password", nil, nil)
			return
		}

		id, exp, err := m.authSvc.GetIdAndExp(claims) 

		if err != nil {
			resp.SendResp(ctx, http.StatusUnauthorized, "failed to convert", nil, err)
		}

		if  time.Now().Unix() >= exp {
			resp.SendResp(ctx, http.StatusUnauthorized, "token duration exceeded", nil, nil)
		}

		ctx.Set("user", id)
		ctx.Next()
	}
}
