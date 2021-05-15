package price

import (
	"fmt"
)

type Price struct {
	saleGoodName string
	buyGoodName  string
	saleAmount   int64
	buyAmount    int64
}

func New(saleGoodName string,
	buyGoodName string,
	saleAmount int64,
	buyAmount int64) Price {
	return Price{
		saleGoodName,
		buyGoodName,
		saleAmount,
		buyAmount}
}

func (p Price) GetSaleGoodName() string {
	return p.saleGoodName
}

func (p Price) GetBuyGoodName() string {
	return p.buyGoodName
}

func (p Price) GetSaleAmount() int64 {
	return p.saleAmount
}

func (p Price) GetBuyAmount() int64 {
	return p.buyAmount
}

func (p Price) PrintPrice() {
	fmt.Printf("\ttrade %3d %-20s for %3d %-20s\n",
		p.saleAmount, p.saleGoodName, p.buyAmount, p.buyGoodName)
}
