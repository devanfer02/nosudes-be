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
	artRepo := repository.NewMysqlArticleRepository(db)
	opHourRepo := repository.NewMysqlOpHoursRepository(db)
	priceAttrRepo := repository.NewMysqlAttractionPricesRepository(db)
	bookmarkRepo := repository.NewMysqlBookmarkRepository(db)
	attrRepo := repository.NewMysqlAttractionRepository(db)
	attrPhotoRepo := repository.NewMysqlAttractionPhotoRepository(db)

	userSvc := service.NewUserService(userRepo, 12 * time.Second)
	authSvc := service.NewAuthService()
	artSvc := service.NewArticleService(artRepo, fileStorage, 20 * time.Second)
	bookmarkSvc := service.NewBookmarkService(
		bookmarkRepo, userRepo, attrRepo, attrPhotoRepo, priceAttrRepo, opHourRepo,
	)
	attrSvc := service.NewAttractionSerivce(
		attrRepo, attrPhotoRepo, priceAttrRepo, opHourRepo, fileStorage, 20 * time.Second,
	)

	userCtr := controller.NewUserController(userSvc)
	authCtr := controller.NewAuthController(userSvc, authSvc)
	artCtr := controller.NewArticleController(artSvc)
	attrCtr := controller.NewAttractionController(attrSvc, bookmarkSvc)

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

	bR := aR.Group("/bookmarks")
	bR.GET("", r.mdlwr.Auth(), ctr.GetBookmarkedByUser)
	bR.POST("", r.mdlwr.Auth(), ctr.BookmarkAttraction)
	bR.DELETE("", r.mdlwr.Auth())
}