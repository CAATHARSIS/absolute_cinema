package tiles

import (
	"net/http"

	"github.com/CAATHARSIS/absolute_cinema/go/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateTileRequestBody struct {
	Title       string `json:"title"`
	Directors   string `json:"directors"`
	Actors      string `json:"actors"`
	Description string `json:"description"`
}

func (h handler) UpdateTile(c *gin.Context) {
	id := c.Param("id")
	body := UpdateTileRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var tile models.Tile

	if result := h.DB.First(&tile, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	tile.Title = body.Title
	tile.Directors = body.Directors
	tile.Actors = body.Actors
	tile.Description = body.Description

	h.DB.Save(&tile)

	c.JSON(http.StatusOK, &tile)
}
