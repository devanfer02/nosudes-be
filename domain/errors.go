package domain

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ErrInternalServer 		= errors.New("internal server error")
	ErrConflict		  		= errors.New("requested item conflicted with existing item")
	ErrNotFound		  		= errors.New("item not found")
	ErrInvalidToken   		= errors.New("invalid token")
	ErrInvalidClaimsDT		= errors.New("invalid claims data type")
	ErrForbidden 			= errors.New("forbidden to modify resources")
	ErrInvalidFileType 		= errors.New("invalid file type")
	ErrForeignItemNotFound 	= errors.New("foreign item with requested id not found")
)

func GetCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if strings.Contains(err.Error(), "validate") {
		return http.StatusBadRequest
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
	case ErrForbidden :
		return http.StatusForbidden
	case ErrInvalidFileType :
		return http.StatusBadRequest
	default : 
		return http.StatusInternalServerError
	}
}