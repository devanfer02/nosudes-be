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
	app *gin.Engine
	db  *sqlx.DB
}

func InitRouter(app *gin.Engine, db *sqlx.DB) {
	r := router{app, db}

	uRepo := repository.NewMysqlUserRepository(db)
	uSvc := service.NewUserService(uRepo, time.Second*12)
	uCtr := controller.NewUserController(uSvc)
	r.setupUserRoutes(uCtr)

	aCtr := controller.NewAuthController(uSvc)
	r.setupAuthRoutes(aCtr)
}

func (r *router) setupUserRoutes(ctr *controller.UserController) {
	uR := r.app.Group("/users")
	uR.GET("", ctr.FetchAll)
	uR.GET("/:id", ctr.FetchByID)
	uR.PUT("/:id", ctr.UpdateUser)
	uR.DELETE("", ctr.DeleteUser)
}

func (r *router) setupAuthRoutes(ctr *controller.AuthController) {
	aR := r.app.Group("/auth")
	aR.POST("/register", ctr.RegisterUser)
	aR.POST("/login", ctr.LoginUser)
	aR.DELETE("/logout", ctr.LogoutUser)
}
