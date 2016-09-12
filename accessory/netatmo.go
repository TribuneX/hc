package accessory

import (
	"github.com/brutella/hc/service"
)

type NetatmoSensor struct {
    *Accessory

    TempSensor *service.TemperatureSensor
    CO2Sensor *service.CarbonDioxideLevelSensor
    HumiditySensor *service.HumiditySensor
}

func NewNetatmoSensor(info Info, temp, humidity, co2 float64) *NetatmoSensor {
    acc := NetatmoSensor{}
    acc.Accessory = New(info, TypeThermostat)
    acc.TempSensor = service.NewTemperatureSensor()
    acc.TempSensor.CurrentTemperature.SetValue(temp)
	acc.TempSensor.CurrentTemperature.SetMinValue(0)
	acc.TempSensor.CurrentTemperature.SetMaxValue(100)
	acc.TempSensor.CurrentTemperature.SetStepValue(0.1)

	acc.AddService(acc.TempSensor.Service)

	acc.HumiditySensor = service.NewHumiditySensor()
	acc.HumiditySensor.CurrentRelativeHumidity.SetValue(humidity)
	acc.HumiditySensor.CurrentRelativeHumidity.SetMinValue(0)
	acc.HumiditySensor.CurrentRelativeHumidity.SetMaxValue(100)
	acc.HumiditySensor.CurrentRelativeHumidity.SetStepValue(0.1)

	acc.AddService(acc.HumiditySensor.Service)


	acc.CarbonDioxideLevelSensor = service.NewCarbonDioxideLevelSensor()
	acc.CarbonDioxideLevelSensor.CurrentRelativeHumidity.SetValue(co2)
	acc.CarbonDioxideLevelSensor.CurrentRelativeHumidity.SetMinValue(0)
	acc.CarbonDioxideLevelSensor.CurrentRelativeHumidity.SetMaxValue(100)
	acc.CarbonDioxideLevelSensor.CurrentRelativeHumidity.SetStepValue(0.1)

	acc.AddService(acc.HumiditySensor.Service)

	return &acc

}