package models

import (
	"net/http"

	"github.com/CAATHARSIS/absolute_cinema/go/pkg/models"
	"github.com/gin-gonic/gin"
)

type AddTileRequestBody struct {
	Title       string   `json:"title"`
	Directors   []string `json:"directors"`
	Actors      []string `json:"actors"`
	Description string   `json:"descriptioin"`
}

func (h handler) AddTile(c *gin.Context) {
	body := AddTileRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var tile models.Tile

	tile.Title = body.Title
	tile.Derectors = body.Directors
	tile.Actors = body.Actors
	tile.Description = body.Description

	if result := h.DB.Create(&tile); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &tile)
}
