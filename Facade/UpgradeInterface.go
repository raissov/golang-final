package Facade

import (
	_ "goFinal/Factory"
)
import "goFinal/Builder"

type UpgradeInterface interface{
	UpgradeStats(*Builder.HeroFinal)
}