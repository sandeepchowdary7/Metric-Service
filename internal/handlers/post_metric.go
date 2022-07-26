package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/deepu/ms/internal/models"
	"github.com/gin-gonic/gin"
)

type PostMetric struct {
	Store map[string][]models.Metric
}

func (h PostMetric) Handle(c *gin.Context) {
	key := c.Param("key")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("provide a value")
		return
	}

	value := struct {
		Value int `json:"value"`
	}{}

	err = json.Unmarshal(body, &value)
	if err != nil {
		log.Println("provide a value")
		return
	}

	metric := models.Metric{
		Key:   key,
		Value: value.Value,
	}

	if _, ok := h.Store[key]; ok {
		l := append(h.Store[key], metric)
		h.Store[key] = l

	} else {
		v := []models.Metric{metric}
		h.Store[key] = v
	}

	c.JSON(http.StatusOK, struct{}{})

}
