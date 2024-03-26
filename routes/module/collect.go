package module

import (
	"github.com/gin-gonic/gin"
	"meetingBooking/api"
)

func LoadCollectRoute(v1 *gin.RouterGroup) {

	category := v1.Group("/collect")
	category.POST("/", api.CreateCollectHandler)
	//category.GET("/list", api.CollectListHandler)
	//category.DELETE("/delete", api.DeleteCollectHandler)

}
