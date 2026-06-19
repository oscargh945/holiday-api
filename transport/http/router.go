package http

import "github.com/gin-gonic/gin"

func NewRouter(handler *HolidayHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/holidays", handler.List)
	}

	return router
}
