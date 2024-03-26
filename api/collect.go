package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"meetingBooking/pkg/format"
	"meetingBooking/services"
	"meetingBooking/utils"
	"net/http"
)

func CreateCollectHandler(ctx *gin.Context) {
	createCollectService := services.CollectService{}
	if err := ctx.ShouldBind(&createCollectService); err != nil {
		msg := utils.GetValidMsg(err, &createCollectService)
		ctx.JSON(http.StatusBadRequest, format.RespErrorWithData(errors.New(msg)))
	} else {
		resp, respErr := createCollectService.Create(ctx.Request.Context())
		if respErr != nil {
			ctx.JSON(http.StatusInternalServerError, format.RespErrorWithData(respErr))
		} else {
			ctx.JSON(http.StatusOK, format.RespSuccessWithData(resp))
		}
	}
}
