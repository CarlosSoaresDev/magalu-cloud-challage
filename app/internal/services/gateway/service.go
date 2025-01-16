package gateway

import (
	"encoding/json"

	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/data/cache"
	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/domain/gateway"
)

const (
	getAllGatewaysCache = "GatewayCache"
)

type GatewayService interface {
	GetAllGatewaysLanguage() (*[]gateway.GatewayDto, error)
}

type gatewayService struct {
	cache cache.CacheClient
}

func New(cache cache.CacheClient) *gatewayService {
	return &gatewayService{
		cache: cache,
	}
}

func (p *gatewayService) GetAllGatewaysLanguage() (*[]gateway.GatewayDto, error) {

	var gateways []gateway.GatewayDto

	c, err := p.cache.Get(getAllGatewaysCache)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(c, &gateways)

	if err != nil {
		return nil, err
	}

	return &gateways, nil
}
