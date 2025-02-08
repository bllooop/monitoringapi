package delivery

import (
	"github.com/bllooop/monitoringapi/backend/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecases *usecase.Usecase
}

func NewHandler(usecases *usecase.Usecase) *Handler {
	return &Handler{usecases: usecases}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	api := router.Group("/api")
	{
		data := api.Group("/data")
		{
			data.POST("/create", h.createData)
			data.GET("/get", h.getData)
		}
	}
	return router
}
