package controller

import (
	"github.com/devanfer02/nosudes-be/domain"
	resp "github.com/devanfer02/nosudes-be/utils/response"

	"github.com/gin-gonic/gin"
)

type ReviewController struct {
	rvSvc domain.ReviewService
}

func NewReviewController(rvSvc domain.ReviewService) *ReviewController {
	return &ReviewController{rvSvc}
}

func (c *ReviewController) FetchAll(ctx *gin.Context) {
	userId := ctx.GetString("user")
	reviews, err := c.rvSvc.FetchAll(ctx.Request.Context(), userId)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to fetch data", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully fetch data", reviews, nil)
}

func (c *ReviewController) FetchByAttrID(ctx *gin.Context) {
	userId := ctx.GetString("user")
	attractionId := ctx.Param("attractionId")

	reviews, err := c.rvSvc.FetchByAttrID(ctx.Request.Context(), attractionId, userId)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to fetch data by attraction id", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully fetch data", reviews, nil)
}

func (c *ReviewController) FetchByID(ctx *gin.Context) {
	userId := ctx.GetString("user")
	id := ctx.Param("id")

	review, err := c.rvSvc.FetchByID(ctx.Request.Context(), id, userId)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to fetch data by id", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully fetch data", review, nil)
}

func (c *ReviewController) CreateReview(ctx *gin.Context) {
	userId := ctx.GetString("user")
	attractionId := ctx.Param("attractionId")
	review := domain.ReviewPayload{}

	if err := ctx.ShouldBind(&review); err != nil {
		resp.SendResp(ctx, 400, "bad form request", nil, err)
		return
	}

	review.Default(attractionId, userId)

	err := c.rvSvc.InsertReview(ctx.Request.Context(), &review)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to create review", nil, err)
		return
	}

	resp.SendResp(ctx, 201, "successfully create review", review, err)
}

func (c *ReviewController) LikeReview(ctx *gin.Context) {
	reviewId := ctx.Param("reviewId")
	userId := ctx.GetString("user")

	err := c.rvSvc.LikeReview(ctx, reviewId, userId)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to like review", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully like review", nil, nil)
}

func (c *ReviewController) UnlikeReview(ctx *gin.Context) {
	reviewId := ctx.Param("reviewId")
	userId := ctx.GetString("user")

	err := c.rvSvc.UnlikeReview(ctx, reviewId, userId)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to unlike review", nil, nil)
		return
	}

	resp.SendResp(ctx, code, "successfully unlike review", nil, nil)
}

func (c *ReviewController) DeleteReview(ctx *gin.Context) {
	id := ctx.Param("id")
	userId := ctx.GetString("user")

	err := c.rvSvc.DeleteReview(ctx.Request.Context(), id, userId)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to delete review", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully delete reivew", nil, nil)
}
