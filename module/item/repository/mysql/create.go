package mysql

import (
	"todo-list/module/item/model"

	"github.com/gin-gonic/gin"
)

func (s *mysqlRepository) CreateItem(ctx *gin.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
