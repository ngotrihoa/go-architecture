package service

import (
	"todo-list/module/item/model"

	"github.com/gin-gonic/gin"
)

func (s *itemService) CreateItem(ctx *gin.Context, data *model.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := s.repository.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
