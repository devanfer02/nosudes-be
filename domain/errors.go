package domain

import (
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var (
	ErrInternalServer 	= errors.New("internal server error")
	ErrConflict		  	= errors.New("requested item conflicted with existing item")
	ErrNotFound		  	= errors.New("item not found")
	ErrInvalidToken   	= errors.New("invalid token")
	ErrInvalidClaimsDT	= errors.New("invalid claims data type")
)

func GetCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch(err) {
	case ErrInternalServer :
		return http.StatusInternalServerError
	case ErrConflict :
 		return http.StatusConflict
	case ErrNotFound :
		return http.StatusNotFound
	case ErrInvalidToken :
		return http.StatusUnauthorized
	default : 
		return http.StatusInternalServerError
	}
}

func IsSQLUniqueViolation(err error) bool {
	sqlerr, ok := err.(*mysql.MySQLError)

	if !ok {
		return false
	}

	return sqlerr.Number == 1062
}
