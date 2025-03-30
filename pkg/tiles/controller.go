package tiles

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRouters(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routers := r.Group("/tiles")
	routers.POST("/", h.AddTile)
	routers.GET("/", h.GetTiles)
	routers.GET("/:id", h.GetTile)
	routers.GET("/:id/", h.GetTile)
	routers.PUT("/:id", h.UpdateTile)
	routers.PUT("/:id/", h.UpdateTile)
	routers.DELETE("/:id", h.DeleteTile)
	routers.DELETE("/:id/", h.DeleteTile)
}
