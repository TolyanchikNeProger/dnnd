package handlers

import (
	"github.com/TolyanchikNeProger/dnnd/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		menus := api.Group("/menus")
		{
			menus.POST("/", h.createMenu)
			menus.GET("/", h.getAllMenus)
			menus.GET("/:menu_id", h.getMenu)
			menus.PUT("/:menu_id", h.updateMenu)
			menus.DELETE("/:menu_id", h.deleteMenu)

			items := menus.Group("/:menu_id/items")
			{
				items.GET("/", h.getMenuItems)
				items.POST("/", h.createMenuItem)
				items.GET("/:item_id", h.getMenuItem)
				items.PUT("/:item_id", h.updateMenuItem)
				items.DELETE("/:item_id", h.deleteMenuItem)
			}
		}
	}

	return router
}
