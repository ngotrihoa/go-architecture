package httpHandler

import (
	"github.com/gin-gonic/gin"
)

func MapItemRoutes(group *gin.RouterGroup, h *httpHandler) {
	itemGroup := group.Group("/items")
	{
		itemGroup.POST("", h.CreateItem())
	}
}
