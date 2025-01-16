package main

import (
	"fmt"
	"os"

	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	engine := setupServer(logger)

	port := getPort()
	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		logger.Fatal("Error starting application", zap.Error(err))
	}
}

func setupServer(logger *zap.Logger) *gin.Engine {
	engine := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	engine.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	router.Initialize(engine, logger)
	return engine
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
