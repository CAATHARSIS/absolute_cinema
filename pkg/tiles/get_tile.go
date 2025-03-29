package tiles

import (
	"net/http"

	"github.com/CAATHARSIS/absolute_cinema/go/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTile(c *gin.Context) {
	id := c.Param("id")

	var tile models.Tile

	if result := h.DB.First(&tile, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tile)
}
