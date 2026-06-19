package http

import "github.com/gin-gonic/gin"

func NewRouter(handler *HolidayHandler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	_ = router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Holiday API is running",
			"endpoint": "/api/v1/holidays",
			"openapi":  "/docs/openapi.yaml",
		})
	})

	router.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")

	v1 := router.Group("/api/v1")
	{
		v1.GET("/holidays", handler.List)
	}

	return router
}
