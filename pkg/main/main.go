package main

import (
	"github.com/dayr-trade-calculator/pkg/cities"
	"github.com/dayr-trade-calculator/pkg/goods"
)

func main() {
	catalogue := goods.MakeCatalogue(goods.Alcohol)
	cities.New("Arzamas", catalogue).PrintCatalogue()
}
