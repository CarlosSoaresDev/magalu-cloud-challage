package coin

import (
	"net/http"

	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/services/coin"
	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CoinHandler struct {
	logger      *zap.Logger
	coinService coin.CoinService
}

func New(logger *zap.Logger, coinService coin.CoinService) *CoinHandler {
	return &CoinHandler{
		logger:      logger,
		coinService: coinService,
	}
}

func (c *CoinHandler) GetAllCoinsHandler(ctx *gin.Context) {

	correlationId := ctx.GetHeader("x-magalu-cloud-correlationId")

	c.logger.Info("[ Started ] - Initialize request to get all coin", zap.String("CorrelationId", correlationId))

	result, err := c.coinService.GetAllCoinLanguage()

	if err != nil {
		c.logger.Error("[ Ended ] - Finalize with error request to get coin", zap.String("CorrelationId", correlationId), zap.Error(err))
		utils.ApiResponse(ctx, http.StatusBadRequest, string("We were unable to process your request, please try later"))
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, result)
	c.logger.Info("[ Ended ] - Finalize request to get coin", zap.String("CorrelationId", correlationId))
}
