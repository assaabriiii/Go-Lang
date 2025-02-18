package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherResponse struct {
	Temperature string `json:"temperature"`
	Description string `json:"description"`
	Wind        string `json:"wind"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/weather", func(c *gin.Context) {
		city := c.Query("city")
		if city == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "City name is required"})
			return
		}

		url := fmt.Sprintf("https://goweather.herokuapp.com/weather/%s", city)
		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var weather WeatherResponse
		err = json.Unmarshal(body, &weather)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "weather.html", gin.H{
			"city":        city,
			"temperature": weather.Temperature,
			"wind":        weather.Wind,
			"description": weather.Description,
		})
	})

	router.Run(":8080")
}
