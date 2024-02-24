package routes

import (
	"time"

	"github.com/devanfer02/nosudes-be/bootstrap/firebase"
	"github.com/devanfer02/nosudes-be/controller"
	"github.com/devanfer02/nosudes-be/middleware"
	"github.com/devanfer02/nosudes-be/repository"
	"github.com/devanfer02/nosudes-be/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type router struct {
	app *gin.Engine
	db  *sqlx.DB
	mdlwr *middleware.Middleware
}

func InitRouter(app *gin.Engine, db *sqlx.DB) {
	
	fileStorage := firebase.NewFirebaseStorage()

	userRepo := repository.NewMysqlUserRepository(db)
	userSvc := service.NewUserService(userRepo, 12 * time.Second)
	userCtr := controller.NewUserController(userSvc)

	authSvc := service.NewAuthService()
	authCtr := controller.NewAuthController(userSvc, authSvc)

	artRepo := repository.NewMysqlArticleRepository(db)
	artSvc := service.NewArticleService(artRepo, fileStorage, 20 * time.Second)
	artCtr := controller.NewArticleController(artSvc)

	opHourRepo := repository.NewMysqlOpHoursRepository(db)

	attrRepo := repository.NewMysqlAttractionRepository(db)
	attrPhotoRepo := repository.NewMysqlAttractionPhotoRepository(db)
	attrSvc := service.NewAttractionSerivce(attrRepo, attrPhotoRepo, opHourRepo, fileStorage, 20 * time.Second)
	attrCtr := controller.NewAttractionController(attrSvc)

	r := router{app, db, middleware.NewMiddleware(userSvc, authSvc)}

	r.setupUserRoutes(userCtr)
	r.setupAuthRoutes(authCtr)
	r.setupArticleRoutes(artCtr)
	r.setupAttractionRoutes(attrCtr)
}

func (r *router) setupUserRoutes(ctr *controller.UserController) {
	uR := r.app.Group("/users")
	uR.GET("", ctr.FetchAll)
	uR.GET("/:id", ctr.FetchByID)
	uR.PUT("", r.mdlwr.Auth(), ctr.UpdateUser)
	uR.DELETE("", r.mdlwr.Auth(), ctr.DeleteUser)
}

func (r *router) setupAuthRoutes(ctr *controller.AuthController) {
	aR := r.app.Group("/auth")
	aR.POST("/register", ctr.RegisterUser)
	aR.POST("/login", ctr.LoginUser)
}

func (r *router) setupArticleRoutes(ctr *controller.ArticleController) {
	aR := r.app.Group("/articles")
	aR.GET("", ctr.FetchAll)
	aR.GET("/:id", ctr.FetchByID)
	aR.POST("", ctr.CreateArticle)
	aR.PUT("/:id", ctr.UpdateArticle)
	aR.DELETE("/:id", ctr.DeleteArticle)
}

func (r *router) setupAttractionRoutes(ctr *controller.AttractionController) {
	aR := r.app.Group("/attractions")
	aR.GET("", ctr.FetchAll)
	aR.GET("/:id", ctr.FetchByID)
	aR.POST("", ctr.InsertAttraction)
	aR.PUT("/:id", ctr.UpdateAttraction)
	aR.POST("/:id", ctr.UploadAttractionPhoto)
	aR.DELETE("/:id", ctr.DeleteAttraction)
}