package controller

import (
	"github.com/devanfer02/nosudes-be/domain"
	resp "github.com/devanfer02/nosudes-be/utils/response"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	artSvc domain.ArticleService
}

func NewArticleController(artSvc domain.ArticleService) *ArticleController {
	return &ArticleController{artSvc}
}

func (c *ArticleController) FetchAll(ctx *gin.Context) {
	articles, err := c.artSvc.FetchAll(ctx.Request.Context())
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to fetch data", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully fetch data", articles, nil)
}

func (c *ArticleController) FetchByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	article, err := c.artSvc.FetchByID(ctx.Request.Context(), idParam)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to fetch data", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully fetch data by id", article, nil)
}

func (c *ArticleController) CreateArticle(ctx *gin.Context) {
	article := domain.ArticlePayload{}

	if err := ctx.ShouldBind(&article); err != nil {
		resp.SendResp(ctx, 400, "failed to update article", nil, err)
		return
	}

	err := c.artSvc.InsertArticle(ctx, &article)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to insert article", nil, err)
		return 
	}
	resp.SendResp(ctx, code, "successfully insert article", nil, nil)
}

func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
	idParam := ctx.Param("id")
	article := domain.ArticlePayload{}

	if err := ctx.ShouldBind(&article); err != nil {
		resp.SendResp(ctx, 400, "failed to update article", nil, err)
		return
	}

	article.ID = idParam
	err := c.artSvc.UpdateArticle(ctx.Request.Context(), &article)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to delete article", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully delete article", nil, nil)
}

func (c *ArticleController) DeleteArticle(ctx *gin.Context) {
	idParam := ctx.Param("id")

	err := c.artSvc.DeleteArticle(ctx.Request.Context(), idParam)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to delete article", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully delete article", nil, nil)
}
