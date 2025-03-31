package server

import (
	httpItemHandler "todo-list/module/item/handler/http"
	itemRepository "todo-list/module/item/repository/mysql"
	itemService "todo-list/module/item/service"
)

func (s *Server) MapRouters() error {
	// Setup dependencies
	//repository
	itemRepo := itemRepository.NewMySQLRepository(s.db)

	//service
	itemSv := itemService.NewItemService(itemRepo)

	//handler
	httpItemHd := httpItemHandler.NewHTTPHandler(itemSv)

	// Setup Routers
	v1 := s.gin.Group("/v1")

	httpItemHandler.MapItemRoutes(v1, httpItemHd)

	return nil
}
