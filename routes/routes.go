package routes

import (
	"time"

	"github.com/devanfer02/nosudes-be/controller"
	"github.com/devanfer02/nosudes-be/repository"
	"github.com/devanfer02/nosudes-be/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type router struct {
	app  *gin.Engine
	db *sqlx.DB
}

func InitRouter(app *gin.Engine, db *sqlx.DB) {
	r := router{app, db}

	r.setupUserRoutes()
}

func(r *router) setupUserRoutes() {
	uRepo := repository.NewMysqlUserRepository(r.db)
	uSvc := service.NewUserService(uRepo, time.Second * 12)
	uCtr := controller.NewUserController(uSvc)

	uR := r.app.Group("/users")

	uR.GET("", uCtr.FetchAll)
	uR.PUT("", uCtr.UpdateUser)
	uR.DELETE("", uCtr.DeleteUser)
}