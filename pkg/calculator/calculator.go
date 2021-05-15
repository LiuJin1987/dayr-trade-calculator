package calculator

import (
	"errors"
	"github.com/LiuJin1987/dayr-trade-calculator/pkg/cities"
	"github.com/LiuJin1987/dayr-trade-calculator/pkg/goods"
	"github.com/LiuJin1987/dayr-trade-calculator/pkg/price"
)

type calculate interface {
	GetHaveTotalPrice(good goods.Good, amount int64) ([]price.Price, error)
}

func GetHaveTotalPrice(goodHave goods.Good, wantAmount int64) ([]price.Price, error) {
	cityHave, err := cities.FindCity(goodHave)
	if nil != err {
		return nil, err
	}
	for goodName, goodPrice := range cityHave.GetCatalogue() {
		if goodName == goodHave.GetName() {
			return calTotalPrice(goodPrice, wantAmount)
		}
	}
	return []price.Price{}, errors.New("error occurred when try to calculate price")
}

func GetWantTotalPrice(goodWant goods.Good, wantAmount int64) ([]price.Price, error) {
	cityHave, err := cities.FindCity(goodWant)
	if nil != err {
		return nil, err
	}
	for goodName, goodPrice := range cityHave.GetCatalogue() {
		if goodName == goodWant.GetName() {
			return calTotalPrice(goodPrice, wantAmount)
		}
	}
	return []price.Price{}, errors.New("error occurred when try to calculate price")
}

func calTotalPrice(goodPrice price.Price, wantAmount int64) ([]price.Price, error) {
	singlePrice := goodPrice
	totalAmount := wantAmount / singlePrice.GetSaleAmount() * singlePrice.GetBuyAmount()
	return []price.Price{price.New(
		singlePrice.GetSaleGoodName(),
		singlePrice.GetBuyGoodName(),
		wantAmount,
		totalAmount)}, nil
}
