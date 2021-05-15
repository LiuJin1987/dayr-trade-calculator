package cities

import (
	"errors"
	"fmt"
	"github.com/LiuJin1987/dayr-trade-calculator/pkg/goods"
	"github.com/LiuJin1987/dayr-trade-calculator/pkg/price"
)

const (
	Arzamas            = "Arzamas"
	Bryansk            = "Bryansk"
	Chelyabinsk        = "Chelyabinsk"
	GorenicjiKiev      = "Gorenicji (Kiev)"
	Irkutsk            = "Irkutsk"
	Krasnoyarsk        = "Krasnoyarsk"
	MagnitogorskFitter = "Magnitogorsk (fitter)"
	MagnitogorskTrader = "Magnitogorsk (trader)"
	Novosibirsk        = "Novosibirsk"
	Omsk               = "Omsk"
	Petropavlovsk      = "Petropavlovsk"
	RostovOnDon        = "Rostov-On-Don"
	Sverdlovsk         = "Sverdlovsk"
	Tula               = "Tula"
	Tver               = "Tver"
	Tyumen             = "Tyumen"
	Voronezh           = "Voronezh"
)

type City struct {
	name      string
	catalogue map[string]price.Price
}

var cities []City

func FindCity(goodWant goods.Good) (City, error) {
	for _, aCity := range cities {
		for goodName := range aCity.GetCatalogue() {
			if goodName == goodWant.GetName() {
				return aCity, nil
			}
		}
	}
	return City{}, errors.New("not found in any city")
}

func PrintCityList() {
	for _, city := range cities {
		fmt.Printf("%s:\n", city.name)
		for _, p := range city.GetCatalogue() {
			p.PrintPrice()
		}
	}
}

func (city *City) GetCatalogue() map[string]price.Price {
	return city.catalogue
}

func (city *City) PrintCatalogue() {
	goods.PrintCatalogue(city.GetCatalogue())
}

func init() {
	cities = []City{
		{Arzamas,
			map[string]price.Price{
				goods.Salt:           price.New(goods.Alcohol, goods.Salt, 1, 5),
				goods.Coffee:         price.New(goods.Lymph, goods.Coffee, 1, 1),
				goods.Chlorcystamine: price.New(goods.RevolverAmmo, goods.Chlorcystamine, 20, 1),
				goods.HandMill:       price.New(goods.Vodka, goods.HandMill, 3, 1),
			},
		},
		{Bryansk,
			map[string]price.Price{
				goods.RubberPart: price.New(goods.EnergizingPotion, goods.RubberPart, 10, 1),
				goods.Hacksaw:    price.New(goods.Vodka, goods.Hacksaw, 5, 1),
				goods.Nettle:     price.New(goods.Flour, goods.Nettle, 1, 10),
				goods.Pasta:      price.New(goods.Coal, goods.Pasta, 50, 1),
			},
		},
		{Chelyabinsk,
			map[string]price.Price{
				goods.Sulfur:         price.New(goods.BoneGlue, goods.Sulfur, 1, 20),
				goods.Cement:         price.New(goods.PlasticExplosives, goods.Cement, 1, 1),
				goods.AutoSpareParts: price.New(goods.ActivatedCharcoal, goods.AutoSpareParts, 50, 1),
				goods.Brick:          price.New(goods.Flour, goods.Brick, 1, 10),
			},
		},
		{GorenicjiKiev, map[string]price.Price{
			goods.FireBrick: price.New(goods.Cloth, goods.FireBrick, 1, 5),
			goods.Lead:      price.New(goods.IronPipe, goods.Lead, 1, 250),
			goods.Byrocarm:  price.New(goods.TTAmmo, goods.Byrocarm, 20, 1),
			goods.Apples:    price.New(goods.AssaultRifleParts, goods.Apples, 1, 3),
			goods.Cigarette: price.New(goods.Threads, goods.Apples, 3, 1),
		},
		},
		{Irkutsk, map[string]price.Price{
			goods.Chainsaw:   price.New(goods.TannedLeather, goods.Chainsaw, 10, 1),
			goods.CarBattery: price.New(goods.Whiskey, goods.CarBattery, 50, 1),
			goods.Sulfur:     price.New(goods.StewMeat, goods.Sulfur, 1, 150),
			goods.Rice:       price.New(goods.Soap, goods.Rice, 1, 5),
		},
		},
		{Krasnoyarsk, map[string]price.Price{
			goods.ChemistrySet:    price.New(goods.Jam, goods.ChemistrySet, 20, 1),
			goods.Sugar:           price.New(goods.Whiskey, goods.Sugar, 1, 5),
			goods.Pepsi:           price.New(goods.SmockedFatback, goods.Pepsi, 1, 1),
			goods.BuckwheatGrains: price.New(goods.Scrap, goods.BuckwheatGrains, 100, 1),
		},
		},
		{MagnitogorskFitter, map[string]price.Price{
			goods.PMShells:           price.New(goods.Steel, goods.PMShells, 1, 10),
			goods.TTShells:           price.New(goods.Steel, goods.TTShells, 1, 10),
			goods.RevolverShells:     price.New(goods.Steel, goods.RevolverShells, 1, 10),
			goods.AssaultRifleShells: price.New(goods.Steel, goods.AssaultRifleShells, 1, 5),
			goods.RifleShells:        price.New(goods.Steel, goods.RifleShells, 1, 5),
		},
		},
		{MagnitogorskTrader, map[string]price.Price{
			goods.Scrap:      price.New(goods.Chitin, goods.Scrap, 1, 100),
			goods.Electrodes: price.New(goods.Lidiacide34, goods.Electrodes, 5, 1),
			goods.Lead:       price.New(goods.MachinOil, goods.Lead, 100, 500),
			goods.Steel:      price.New(goods.HandmadeRocket, goods.Steel, 1, 100),
		},
		},
		{Novosibirsk, map[string]price.Price{
			goods.GasMaskFilter: price.New(goods.Antibiotics, goods.GasMaskFilter, 10, 1),
			goods.SparkPlug:     price.New(goods.EnergizingPotion, goods.SparkPlug, 1, 1),
			goods.Threads:       price.New(goods.Honey, goods.Threads, 1, 10),
			goods.Lighter:       price.New(goods.PickledVegetables, goods.Lighter, 3, 1),
		},
		},
		{Omsk, map[string]price.Price{
			goods.Battery:      price.New(goods.Primer, goods.Battery, 5, 1),
			goods.ChemicalSuit: price.New(goods.Cement, goods.ChemicalSuit, 5, 1),
			goods.FreshBones:   price.New(goods.HomemadeWine, goods.FreshBones, 1, 25),
			goods.Tires:        price.New(goods.Ir190, goods.Tires, 1, 4),
		},
		},
		{Petropavlovsk, map[string]price.Price{
			goods.Paper: price.New(goods.Primer, goods.Paper, 1, 1),
			goods.Tea:   price.New(goods.Cement, goods.Tea, 20, 1),
			goods.Rags:  price.New(goods.HomemadeWine, goods.Rags, 1, 20),
			goods.Corn:  price.New(goods.Ir190, goods.Corn, 1, 1),
		},
		},
		{RostovOnDon, map[string]price.Price{
			goods.ToolKit:    price.New(goods.GunpowderGrenades, goods.ToolKit, 3, 1),
			goods.Wheat:      price.New(goods.BoneGlue, goods.Wheat, 1, 1),
			goods.Wine:       price.New(goods.RifleAmmo, goods.Wine, 20, 1),
			goods.CannedBeef: price.New(goods.RifleParts, goods.CannedBeef, 1, 1),
			goods.Salt:       price.New(goods.Gasoline, goods.Salt, 500, 10),
		},
		},
		{Sverdlovsk, map[string]price.Price{
			goods.InsulatingTape:  price.New(goods.Primer, goods.InsulatingTape, 1, 3),
			goods.Welder:          price.New(goods.Cement, goods.Welder, 150, 1),
			goods.ElectricalCable: price.New(goods.HomemadeWine, goods.ElectricalCable, 100, 1),
			goods.Vegetables:      price.New(goods.Ir190, goods.Vegetables, 3, 1),
		},
		},
		{Tula, map[string]price.Price{
			goods.Saltpeter: price.New(goods.Soap, goods.Saltpeter, 1, 100),
			goods.TTAmmo:    price.New(goods.HealingSalve, goods.TTAmmo, 1, 3),
			goods.Wire:      price.New(goods.Strawberry, goods.Wire, 5, 1),
			goods.Pan:       price.New(goods.Cloth, goods.Pan, 10, 1),
		},
		},
		{Tver, map[string]price.Price{
			goods.Rope:              price.New(goods.TanningMixture, goods.Rope, 1, 2),
			goods.Crowbar:           price.New(goods.DriedMeat, goods.Crowbar, 50, 1),
			goods.BicycleSpareParts: price.New(goods.FreshBones, goods.BicycleSpareParts, 5, 1),
			goods.Potatoes:          price.New(goods.Poison, goods.Potatoes, 1, 2),
		},
		},
		{Tyumen, map[string]price.Price{
			goods.Gasoline:   price.New(goods.SulfuricAcid, goods.Gasoline, 1, 1000),
			goods.Diesel:     price.New(goods.SulfuricAcid, goods.Diesel, 1, 2000),
			goods.MachineOil: price.New(goods.Fat, goods.MachineOil, 1, 50),
			goods.Coal:       price.New(goods.Salt, goods.Coal, 1, 25),
		},
		},
		{Voronezh, map[string]price.Price{
			goods.Vodka:           price.New(goods.MolotovCocktail, goods.Vodka, 1, 1),
			goods.Metocaine:       price.New(goods.PMAmmo, goods.Metocaine, 20, 1),
			goods.MotorcycleParts: price.New(goods.Scrap, goods.MotorcycleParts, 50, 1),
			goods.CondensedMilk:   price.New(goods.PistolParts, goods.CondensedMilk, 2, 1),
		},
		},
	}

}
