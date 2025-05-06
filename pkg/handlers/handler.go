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

	api := router.Group("/api", h.signIn) // Требуется только аутентификация
	{
		// Публичные элементы меню (доступны всем)
		menu := api.Group("/menu-items")
		{
			menu.GET("/", h.getAllMenuItems)
			menu.GET("/:item_id", h.getMenuItem)
		}

		// Управление комнатами (доступно всем аутентифицированным пользователям)
		rooms := api.Group("/rooms")
		{
			rooms.POST("/", h.createRoom)           // Создать комнату (стать ГМ)
			rooms.GET("/my", h.getUserRooms)        // Получить свои комнаты
			rooms.GET("/:room_id", h.getRoom)       // Просмотр любой комнаты
			rooms.PUT("/:room_id", h.updateRoom)    // Изменить свою комнату
			rooms.DELETE("/:room_id", h.deleteRoom) // Удалить свою комнату
		}

		// Персональные настройки меню
		user := api.Group("/user")
		{
			user.GET("/menu", h.getUserMenu)
			user.PUT("/menu", h.updateUserMenu)
		}
	}

	return router
}
