package rest

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	v1 := router.Group("/api/v1")
	v1.Use(newBasicMiddleware(h.log))

	v1.POST("/add-yt-video", h.addYTVideoByURL)
	v1.POST("/add-yt-videos/recent", h.addRecentYTVideos)
	v1.POST("/add-llm-analyze/:recordId", h.addLLMContentAnalyze)
	v1.GET("/records", h.takeRecords)

	return router
}
