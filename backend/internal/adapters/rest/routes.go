package rest

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	v1 := router.Group("/api/v1")
	v1.Use(newBasicMiddleware(h.log))

	//routes...

	return router
}
