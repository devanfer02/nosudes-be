package controller

import (
	"github.com/devanfer02/nosudes-be/domain"
	resp "github.com/devanfer02/nosudes-be/utils/response"

	"github.com/gin-gonic/gin"
)

type AttractionController struct {
	attrSvc domain.AttractionService
}

func NewAttractionController(attrSvc domain.AttractionService) *AttractionController {
	return &AttractionController{attrSvc}
}

func (c *AttractionController) FetchAll(ctx *gin.Context) {
	attractions, err := c.attrSvc.FetchAll(ctx.Request.Context())
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to fetch data", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully fetch data", attractions, nil)
}

func (c *AttractionController) FetchByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	attraction, err := c.attrSvc.FetchByID(ctx.Request.Context(), idParam)
	code := domain.GetCode(err)

	if err != nil {
		if err == domain.ErrFailedFetchOtherAPI {
			resp.SendResp(ctx, code, "successfully fetch data, failed to fetch some data from other API", attraction, nil)
			return 
		}

		resp.SendResp(ctx, code, "failed to fetch data", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully fetch data by id", attraction, nil)
}

func (c *AttractionController) InsertAttraction(ctx *gin.Context) {
	attraction := domain.AttractionPayload{}

	if bindFailed(ctx, &attraction) {
		return
	}

	err := c.attrSvc.InsertAttraction(ctx.Request.Context(), &attraction)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to insert attraction", nil, err)
		return
	}

	resp.SendResp(ctx, 201, "successfully insert attraction", attraction, nil)
}

func (c *AttractionController) UpdateAttraction(ctx *gin.Context) {
	attraction := domain.AttractionPayload{}
	idParam := ctx.Param("id")

	if bindFailed(ctx, &attraction) {
		return
	}

	attraction.ID = idParam
	err := c.attrSvc.UpdateAttraction(ctx.Request.Context(), &attraction)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to update attraction", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully update attraction", attraction, nil)
}

func (c *AttractionController) UploadAttractionPhoto(ctx *gin.Context) {
	attrPhoto := domain.AttractionPhotoPayload{}
	idParam := ctx.Param("id")

	if err := ctx.ShouldBind(&attrPhoto); err != nil {
		resp.SendResp(ctx, 400, "failed to upload photo", nil, err)
		return
	}

	attrPhoto.AttractionID = idParam
	err := c.attrSvc.UploadPhotoByAttID(ctx.Request.Context(), &attrPhoto)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to upload photo", nil, err)
		return
	}

	resp.SendResp(ctx, 201, "successfully upload attraction photo", nil, nil)
}

func (c *AttractionController) DeleteAttraction(ctx *gin.Context) {
	idParam := ctx.Param("id")

	err := c.attrSvc.DeleteAttraction(ctx.Request.Context(), idParam)
	code := domain.GetCode(err)

	if err != nil {
		resp.SendResp(ctx, code, "failed to delete attraction", nil, err)
		return
	}

	resp.SendResp(ctx, code, "successfully delete attraction", nil, nil)
}
