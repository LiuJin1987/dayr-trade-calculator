package price

import "fmt"

type Price struct {
	saleGoodName string
	buyGoodName  string
	saleAmount   int64
	buyAmount    int64
}

func New(saleGoodName string,
	buyGoodName string,
	saleAmount int64,
	buyAmount int64) *Price {
	return &Price{
		saleGoodName,
		buyGoodName,
		saleAmount,
		buyAmount}
}

func (p Price) PrintPrice() {
	fmt.Printf("\ttrade %3d %-20s for %3d %-20s\n",
		p.saleAmount, p.saleGoodName, p.buyAmount, p.buyGoodName)
}
