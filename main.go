package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Config struct {
	Host string
	Port int
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadConfig() (*Config, error) {
	host := getEnv("HOST", "0.0.0.0")
	portStr := getEnv("PORT", "80")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT format: %v", err)
	}

	return &Config{
		Host: host,
		Port: port,
	}, nil
}

var serveruuid = uuid.New().String()

func uuidHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"uuid": serveruuid,
	})
}

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", uuidHandler)

	router.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
}
