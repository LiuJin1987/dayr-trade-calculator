package goods

import "fmt"

const (
	ActivatedCharcoal  = "Activated charcoal"
	Alcohol            = "Alcohol"
	Antibiotics        = "Antibiotics"
	Apples             = "Apples"
	AssaultRifleAmmo   = "Assault rifle ammo"
	AssaultRifleParts  = "Assault rifle parts"
	AssaultRifleShells = "Assault rifle shells"
	AutoSpareParts     = "Auto spare parts"
)

type Good string

func MakeCatalogue(names ...string) []Good {
	catalogue := make([]Good, 0)
	var goods []Good
	for _, name := range names {
		goods = append(catalogue, Good(name))
	}
	return goods
}

func PrintCatalogue(goods []Good) {
	for _, good := range goods {
		fmt.Println(good)
	}
}
