package service

import (
	"todo-list/module/item/model"

	"github.com/gin-gonic/gin"
)

type ItemRepository interface {
	CreateItem(ctx *gin.Context, data *model.TodoItemCreation) error
}

type itemService struct {
	repository ItemRepository
}

func NewItemService(repository ItemRepository) *itemService {
	return &itemService{repository: repository}
}
