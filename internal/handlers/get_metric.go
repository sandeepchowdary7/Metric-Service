package handlers

import (
	"net/http"
	"time"

	"github.com/deepu/ms/internal/models"
	"github.com/gin-gonic/gin"
)

type GetMetric struct {
	Store map[string][]models.Metric
}

func (h GetMetric) Handle(c *gin.Context) {
	key := c.Param("key")

	if _, ok := h.Store[key]; ok {
		var sum int = 0
		for _, v := range h.Store[key] {
			if time.Since(v.Time).Hours() < 1 {
				sum += v.Value
			}
		}

		c.JSON(http.StatusOK, struct {
			Value int `json:"value"`
		}{
			Value: sum,
		})
	} else {
		c.JSON(http.StatusNotFound, struct {
			Error string `json:"error"`
		}{
			Error: "Key not found",
		})
	}
}
