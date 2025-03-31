package httpHandler

import (
	"todo-list/module/item/model"

	"github.com/gin-gonic/gin"
)

type ItemService interface {
	CreateItem(ctx *gin.Context, data *model.TodoItemCreation) error
}

type httpHandler struct {
	service ItemService
}

func NewHTTPHandler(service ItemService) *httpHandler {
	return &httpHandler{service: service}
}
