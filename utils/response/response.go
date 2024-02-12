package resp

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code	 	int  			`json:"code"`
	Message		string			`json:"message"`
	Data		interface{} 	`json:"data,omitempty"`
	Err			string 			`json:"err_message,omitempty"`
}

func SendResp(ctx *gin.Context, code int, message string, data interface{}, err error) {
	errMessage := ""

	if err != nil {
		errMessage = err.Error()
	}

	ctx.JSON(
		code, 
		&Response{code, message, data, errMessage},
	)
}