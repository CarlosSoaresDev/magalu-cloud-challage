package router

import (
	"net/http"

	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/data/cache"
	coinHandler "github.com/CarlosSoaresDev/magalu-cloud-challage/internal/handlers/coin"
	gatewayHandler "github.com/CarlosSoaresDev/magalu-cloud-challage/internal/handlers/gateway"
	coinService "github.com/CarlosSoaresDev/magalu-cloud-challage/internal/services/coin"
	gatewayService "github.com/CarlosSoaresDev/magalu-cloud-challage/internal/services/gateway"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Initialize(route *gin.Engine, logger *zap.Logger) {
	cacheClient := cache.New()

	coinService := coinService.New(cacheClient)
	coinHandler := coinHandler.New(logger, coinService)

	gatewayService := gatewayService.New(cacheClient)
	gatewayHandler := gatewayHandler.New(logger, gatewayService)

	groupRoute := route.Group("/api/v1")

	subscribeRoute := groupRoute.Group("/coins")
	{
		subscribeRoute.GET("", coinHandler.GetAllCoinsHandler)
	}

	personUserRoute := groupRoute.Group("/subscribers")
	{
		personUserRoute.GET("", gatewayHandler.GetAllGatewaysHandler)
	}

	route.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
}
