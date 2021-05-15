package cities

import (
	"fmt"
	"github.com/LiuJin1987/dayr-trade-calculator/pkg/goods"
	"github.com/LiuJin1987/dayr-trade-calculator/pkg/price"
)

const (
	Arzamas             = "Arzamas"
	Bryansk             = "Bryansk"
	Chelyabinsk         = "Chelyabinsk"
	GorenicjiKiev       = "Gorenicji (Kiev)"
	Irkutsk             = "Irkutsk"
	Krasnoyarsk         = "Krasnoyarsk"
	Magnitogorsk_Fitter = "Magnitogorsk (fitter)"
	MagnitogorskTrader  = "Magnitogorsk (trader)"
	Novosibirsk         = "Novosibirsk"
	Omsk                = "Omsk"
	Petropavlovsk       = "Petropavlovsk"
	Rostov_On_Don       = "Rostov-On-Don"
	Sverdlovsk          = "Sverdlovsk"
	Tula                = "Tula"
	Tver                = "Tver"
	Tyumen              = "Tyumen"
	Voronezh            = "Voronezh"
)

type City struct {
	name      string
	catalogue []goods.Good
}

var cities []City

func GetCityList() []City {
	return cities
}

func PrintCityList() {
	for _, city := range cities {
		fmt.Printf("%s:\n", city.name)
		for _, good := range city.GetCatalogue() {
			good.GetPrice().PrintPrice()
		}
	}
}

func New(name string, catalogue []goods.Good) *City {
	return &City{name, catalogue}
}

func (city *City) GetCatalogue() []goods.Good {
	return city.catalogue
}

func (city *City) PrintCatalogue() {
	goods.PrintCatalogue(city.GetCatalogue())
}

func init() {
	cities = []City{
		City{Arzamas, []goods.Good{
			goods.New(goods.Salt,
				price.New(goods.Alcohol, goods.Salt, 1, 5)),
			goods.New(goods.Coffee,
				price.New(goods.Lymph, goods.Coffee, 1, 1)),
			goods.New(goods.Chlorcystamine,
				price.New(goods.RevolverAmmo, goods.Chlorcystamine, 20, 1)),
			goods.New(goods.HandMill,
				price.New(goods.Vodka, goods.HandMill, 3, 1))},
		},
		City{Bryansk, []goods.Good{
			goods.New(goods.RubberPart,
				price.New(goods.EnergizingPotion, goods.RubberPart, 10, 1)),
			goods.New(goods.Hacksaw,
				price.New(goods.Vodka, goods.Hacksaw, 5, 1)),
			goods.New(goods.Nettle,
				price.New(goods.Flour, goods.Nettle, 1, 10)),
			goods.New(goods.Pasta,
				price.New(goods.Coal, goods.Pasta, 50, 1))},
		},
		City{Chelyabinsk, []goods.Good{
			goods.New(goods.Sulfur,
				price.New(goods.BoneGlue, goods.Sulfur, 1, 20)),
			goods.New(goods.Cement,
				price.New(goods.PlasticExplosives, goods.Cement, 1, 1)),
			goods.New(goods.AutoSpareParts,
				price.New(goods.ActivatedCharcoal, goods.AutoSpareParts, 50, 1)),
			goods.New(goods.Brick,
				price.New(goods.Flour, goods.Brick, 1, 10))},
		},
		City{GorenicjiKiev, []goods.Good{
			goods.New(goods.FireBrick,
				price.New(goods.Cloth, goods.FireBrick, 1, 5)),
			goods.New(goods.Lead,
				price.New(goods.IronPipe, goods.Lead, 1, 250)),
			goods.New(goods.Byrocarm,
				price.New(goods.TTAmmo, goods.Byrocarm, 20, 1)),
			goods.New(goods.Apples,
				price.New(goods.AssaultRifleParts, goods.Apples, 1, 3)),
			goods.New(goods.Cigarette,
				price.New(goods.Threads, goods.Apples, 3, 1))},
		},

		City{Irkutsk, []goods.Good{
			goods.New(goods.Chainsaw,
				price.New(goods.Cloth, goods.Chainsaw, 1, 5)),
			goods.New(goods.CarBattery,
				price.New(goods.IronPipe, goods.CarBattery, 1, 250)),
			goods.New(goods.Sulfur,
				price.New(goods.TTAmmo, goods.Sulfur, 20, 1)),
			goods.New(goods.Rice,
				price.New(goods.AssaultRifleParts, goods.Rice, 1, 3))},
		},
		/*,
		City{,
			goods.MakeCatalogue(
				goods.,
				goods.,
				goods.,
				goods.)},
		City{Krasnoyarsk,
			goods.MakeCatalogue(
				goods.ChemistrySet,
				goods.Sugar,
				goods.Pepsi,
				goods.BuckwheatGrains)},
		City{Magnitogorsk_Fitter,
			goods.MakeCatalogue(
				goods.PMShells,
				goods.TTShells,
				goods.RevolverShells,
				goods.AssaultRifleShells,
				goods.RifleShells)},
		City{MagnitogorskTrader,
			goods.MakeCatalogue(
				goods.Scrap,
				goods.Electrodes,
				goods.Lead,
				goods.Steel)},
		City{Novosibirsk,
			goods.MakeCatalogue(
				goods.GasMaskFilter,
				goods.SparkPlug,
				goods.Threads,
				goods.Lighter)},
		City{Omsk,
			goods.MakeCatalogue(
				goods.Battery,
				goods.ChemicalSuit,
				goods.FreshBones,
				goods.Tires)},
		City{Petropavlovsk,
			goods.MakeCatalogue(
				goods.Paper,
				goods.Tea,
				goods.Rags,
				goods.Corn)},
		City{Rostov_On_Don,
			goods.MakeCatalogue(
				goods.ToolKit,
				goods.Wheat,
				goods.Wine,
				goods.CannedBeef,
				goods.Salt)},
		City{Sverdlovsk,
			goods.MakeCatalogue(
				goods.InsulatingTape,
				goods.Welder,
				goods.ElectricalCable,
				goods.Vegetables)},
		City{Tula,
			goods.MakeCatalogue(
				goods.Saltpeter,
				goods.TTAmmo,
				goods.Wire,
				goods.Pan)},
		City{Tver,
			goods.MakeCatalogue(
				goods.Rope,
				goods.Crowbar,
				goods.BicycleSpareParts,
				goods.Potatoes)},
		City{Tyumen,
			goods.MakeCatalogue(
				goods.Gasoline,
				goods.Diesel,
				goods.MachineOil,
				goods.Coal)},
		City{Voronezh,
			goods.MakeCatalogue(
				goods.Vodka,
				goods.Metocaine,
				goods.MotorcycleParts,
				goods.CondensedMilk)}*/
	}
}
