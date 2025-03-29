package tiles

import (
	"net/http"

	"github.com/CAATHARSIS/absolute_cinema/go/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTiles(c *gin.Context) {
	var tiles []models.Tile

	if result := h.DB.Find(&tiles); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tiles)
}
