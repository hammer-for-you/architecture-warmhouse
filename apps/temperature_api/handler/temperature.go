package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

// GetTemperature handles /temperature?location=:location&sensorID=:sensorId
func GetTemperature(c *gin.Context) {
	location := c.Query("location")
	sensorId := c.Query("sensorID")

	if sensorId == "" && location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No location or sensor id"})
		return
	}

	if sensorId == "" {
		switch location {
		case "Living Room":
			sensorId = "1"
		case "Bedroom":
			sensorId = "2"
		case "Kitchen":
			sensorId = "3"
		default:
			sensorId = "0"
		}
	}

	if location == "" {
		location = findLocation(sensorId)
	}

	c.JSON(http.StatusOK, createResponse(location, sensorId))
}

// GetTemperatureBySensorId handles GET /temperature/:sensorId
func GetTemperatureBySensorId(c *gin.Context) {
	sensorId := c.Param("sensorId")
	if sensorId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sensor id required"})
		return
	}
	location := findLocation(sensorId)
	c.JSON(http.StatusOK, createResponse(location, sensorId))
}

func findLocation(sensorId string) string {
	switch sensorId {
	case "1":
		return "Living Room"
	case "2":
		return "Bedroom"
	case "3":
		return "Kitchen"
	default:
		return "Unknown"
	}

}

func createResponse(location string, sensorId string) TemperatureResponse {
	temperature := rand.Float64()*8 + 19
	return TemperatureResponse{
		Value:       floorToOneDecimal(temperature),
		Unit:        "C",
		Timestamp:   time.Now().UTC(),
		Location:    location,
		SensorID:    sensorId,
		Status:      "OK",
		SensorType:  "Temperature",
		Description: fmt.Sprintf("Sensor in '%s'", location),
	}
}

func floorToOneDecimal(value float64) float64 {
	return math.Floor(value*10) / 10
}
