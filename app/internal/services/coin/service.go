package coin

import (
	"encoding/json"

	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/data/cache"
	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/domain/coin"
)

const (
	getAllCoinsCache = "CoinsCache"
)

type CoinService interface {
	GetAllCoinLanguage() (*[]coin.CoinDto, error)
}

type coinService struct {
	cache cache.CacheClient
}

func New(cache cache.CacheClient) *coinService {
	return &coinService{
		cache: cache,
	}
}

func (p *coinService) GetAllCoinLanguage() (*[]coin.CoinDto, error) {

	var coins []coin.CoinDto

	cacheLanguages, err := p.cache.Get(getAllCoinsCache)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(cacheLanguages, &coins)

	if err != nil {
		return nil, err
	}

	return &coins, nil
}
