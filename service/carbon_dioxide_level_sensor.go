// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeCarbonDioxideSensor = "97"

type CarbonDioxideLevelSensor struct {
	*Service

	CarbonDioxideLevel *characteristic.CarbonDioxideLevel
}

func NewCarbonDioxideLevelSensor() *CarbonDioxideSensor {
	svc := CarbonDioxideSensor{}
	svc.Service = New(TypeCarbonDioxideSensor)

	svc.CarbonDioxideLevel = characteristic.NewCarbonDioxideLevel()
	svc.AddCharacteristic(svc.CarbonDioxideLevel.Characteristic)

	return &svc
}
