package trade

import (
	"github.com/dayr-trade-calculator/pkg/cities"
	"github.com/dayr-trade-calculator/pkg/goods"
)

type trade interface {
	tradeFor(good goods.Good, amount int64)
}

type Trader struct {
	inventory []goods.Good
}

func New(inventory []goods.Good) *Trader {
	return &Trader{inventory}
}

func (trader *Trader) tradeFor(goodWith goods.Good, amount int64) {
	for _, city := range cities.GetCityList() {
		for _, goodInCity := range city.GetCatalogue() {
			if goodWith.Equals(goodInCity) {

			}
		}
	}
}
