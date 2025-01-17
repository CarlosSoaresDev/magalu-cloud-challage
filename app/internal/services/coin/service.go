package coin

import (
	"encoding/json"
	"time"

	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/data/cache"
	"github.com/CarlosSoaresDev/magalu-cloud-challage/internal/domain/coin"
)

const (
	getAllCoinsCache = "CoinsCache"
)

type CoinService interface {
	GetAllCoinLanguage() (*[]coin.CoinDto, error)
	Count() error
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

	coinBase, err := p.cache.Get(getAllCoinsCache)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(coinBase, &coins)

	if err != nil {
		return nil, err
	}

	return &coins, nil
}

func (p *coinService) Count() error {
	var coins []coin.CoinDto

	coins = append(coins, coin.CoinDto{Amount: 10, FromCurrency: "seila", ToCurrency: "alguem"})

	coinsSerialized, err := json.Marshal(coins)

	if err != nil {
		return err
	}

	err = p.cache.Set(getAllCoinsCache, coinsSerialized, 5*time.Hour)

	if err != nil {
		return err
	}

	return err
}
