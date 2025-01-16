package gateway

import (
	"net/http"

	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/services/gateway"
	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GatewayHandler struct {
	logger         *zap.Logger
	gatewayService gateway.GatewayService
}

func New(logger *zap.Logger, gatewayService gateway.GatewayService) *GatewayHandler {
	return &GatewayHandler{
		logger:         logger,
		gatewayService: gatewayService,
	}
}

func (c *GatewayHandler) GetAllGatewaysHandler(ctx *gin.Context) {

	correlationId := ctx.GetHeader("x-magalu-cloud-correlationId")

	c.logger.Info("[ Started ] - Initialize request to get all gateway", zap.String("CorrelationId", correlationId))

	result, err := c.gatewayService.GetAllGatewaysLanguage()

	if err != nil {
		c.logger.Error("[ Ended ] - Finalize with error request to get gateway", zap.String("CorrelationId", correlationId), zap.Error(err))
		utils.ApiResponse(ctx, http.StatusBadRequest, string("We were unable to process your request, please try later"))
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, result)
	c.logger.Info("[ Ended ] - Finalize request to get gateway", zap.String("CorrelationId", correlationId))
}
