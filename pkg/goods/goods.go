package goods

import (
	"fmt"
	"github.com/LiuJin1987/dayr-trade-calculator/pkg/price"
	"strings"
)

type Good struct {
	name  string
	price price.Price
}

func New(name string, price price.Price) Good {
	return Good{name, price}
}

func (good Good) GetName() string {
	return good.name
}

func (good *Good) SetName(name string) {
	good.name = name
}

func (good Good) GetPrice() price.Price {
	return good.price
}

func (good *Good) SetPrice(priceToSet price.Price) {
	good.price = priceToSet
}

func PrintCatalogue(catalogue map[string]price.Price) {
	for _, v := range catalogue {
		fmt.Println(v)
	}
}

func (good Good) Equals(goodToCompareWith Good) bool {
	return strings.EqualFold(good.name, goodToCompareWith.name)
}

const (
	ActivatedCharcoal  = "Activated charcoal"
	Alcohol            = "Alcohol"
	Antibiotics        = "Antibiotics"
	Apples             = "Apples"
	AssaultRifleAmmo   = "Assault rifle ammo"
	AssaultRifleParts  = "Assault rifle parts"
	AssaultRifleShells = "Assault rifle shells"
	AutoSpareParts     = "Auto spare parts"
	Battery            = "Battery"
	BicycleSpareParts  = "Bicycle spare parts"
	BoneGlue           = "Bone glue"
	Brick              = "Brick"
	BuckwheatGrains    = "Buckwheat grains"
	Byrocarm           = "Byrocarm"
	CannedBeef         = "Canned beef"
	CarBattery         = "Car battery"
	Cement             = "Cement"
	Chainsaw           = "Chainsaw"
	ChemicalSuit       = "Chemical suit"
	ChemistrySet       = "Chemistry set"
	Chitin             = "Chitin"
	Chlorcystamine     = "Chlorcystamine"
	Cigarette          = "Cigarette"
	Cloth              = "Cloth"
	Coal               = "Coal"
	Coffee             = "Coffee"
	CondensedMilk      = "Condensed milk"
	Corn               = "Corn"
	Crowbar            = "Crowbar"
	DetoxifyingPotion  = "Detoxifying potion"
	Diesel             = "Diesel"
	DriedFish          = "Dried fish"
	DriedMeat          = "Dried meat"
	ElectricalCable    = "Electrical cable"
	Electrodes         = "Electrodes"
	EnergizingPotion   = "Energizing potion"
	Fat                = "Fat"
	FireBrick          = "Fire brick"
	Flour              = "Flour"
	FreshBones         = "Fresh bones"
	FreshFish          = "Fresh fish"
	GasMaskFilter      = "Gas mask filter"
	Gasoline           = "Gasoline"
	Gunpowder          = "Gunpowder"
	GunpowderGrenades  = "Gunpowder grenades"
	Hacksaw            = "Hacksaw"
	HandMill           = "Hand mill"
	HandmadeRocket     = "Handmade rocket"
	HealingSalve       = "Healing salve"
	HomemadeWine       = "Homemade wine"
	Honey              = "Honey"
	InsulatingTape     = "Insulating tape"
	Ir190              = "IR-190"
	IronPipe           = "Iron pipe"
	Jam                = "Jam"
	Lead               = "Lead"
	Lidiacide34        = "Lidiacide-34"
	Lighter            = "Lighter"
	Lymph              = "Lymph"
	MachinOil          = "Machin oil"
	MachineOil         = "Machine oil"
	Metocaine          = "Metocaine"
	MolotovCocktail    = "Molotov cocktail"
	MotorcycleParts    = "Motorcycle parts"
	Nail               = "Nail"
	Nettle             = "Nettle"
	Pan                = "Pan"
	Paper              = "Paper"
	Pasta              = "Pasta"
	Pepsi              = "Pepsi"
	PickledVegetables  = "Pickled vegetables"
	PistolParts        = "Pistol parts"
	PlasticExplosives  = "Plastic explosives"
	PMAmmo             = "PM ammo"
	PMShells           = "PM shells"
	Poison             = "Poison"
	Potatoes           = "Potatoes"
	Primer             = "Primer "
	Rags               = "Rags"
	RevolverAmmo       = "Revolver ammo"
	RevolverShells     = "Revolver shells"
	Rice               = "Rice"
	RifleAmmo          = "Rifle ammo"
	RifleParts         = "Rifle parts"
	RifleShells        = "Rifle shells"
	Rope               = "Rope"
	RubberPart         = "Rubber part"
	Salt               = "Salt"
	Saltpeter          = "Saltpeter"
	Scrap              = "Scrap"
	SmockedFatback     = "Smocked fatback"
	Soap               = "Soap"
	SparkPlug          = "Spark plug"
	Steel              = "Steel"
	StewMeat           = "Stew meat"
	Strawberry         = "Strawberry"
	Sugar              = "Sugar"
	Sulfur             = "Sulfur"
	SulfuricAcid       = "Sulfuric acid"
	TannedLeather      = "Tanned leather"
	TanningMixture     = "Tanning mixture"
	Tea                = "Tea"
	Threads            = "Threads"
	Tires              = "Tires"
	ToolKit            = "Tool kit"
	TTAmmo             = "TT Ammo"
	TTShells           = "TT shells"
	Vegetables         = "Vegetables"
	Vodka              = "Vodka"
	Welder             = "Welder"
	Wheat              = "Wheat"
	Whiskey            = "Whiskey"
	Wine               = "Wine"
	Wire               = "Wire"
)
