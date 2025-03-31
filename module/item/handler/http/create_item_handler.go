package httpHandler

import (
	"net/http"
	"todo-list/common"
	"todo-list/module/item/model"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) CreateItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var itemData model.TodoItemCreation

		if err := ctx.ShouldBind(&itemData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := h.service.CreateItem(ctx, &itemData); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(itemData.Id))
	}
}
