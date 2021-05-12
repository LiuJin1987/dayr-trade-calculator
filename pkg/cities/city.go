package cities

import (
	"fmt"
	"github.com/dayr-trade-calculator/pkg/goods"
)

const (
	Arzamas = "Arzamas"
	Bryansk = "Bryansk"
)

type City struct {
	name      string
	catalogue []goods.Good
}

var cities []City

func init() {
	cities = []City{
		City{Arzamas, goods.MakeCatalogue(goods.Alcohol)},
	}
}

func GetCityList() []City {
	return cities
}

func New(name string, catalogue []goods.Good) *City {
	return &City{name, catalogue}
}

func (city *City) GetCatalogue() []goods.Good {
	return city.catalogue
}

func (city *City) PrintCatalogue() {
	for _, good := range city.GetCatalogue() {
		fmt.Println(good)
	}
}
