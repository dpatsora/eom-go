package maintenance

import (
	"github.com/dpatsora/eom-go/pkg/config"
	"github.com/dpatsora/eom-go/pkg/workmode"
	"math"
)

func OperatingTimeForMajorRepairs(config *config.Config) float32 {
	koef_naprac := float32(config.IndividualValues.KoefNaprac)
	capital_repair_periodicity := float32(config.IndividualValues.Periodicity.Capital)
	return capital_repair_periodicity * koef_naprac
}

func AmountOfMajorRepairs(config *config.Config) int {
	amount := float64((OperatingTimeForMajorRepairs(config) + workmode.AnnualOperationMode(config))/float32(config.IndividualValues.Periodicity.Capital))
	return int(math.Round(float64(amount-0.5)))
}

func AmountOfRegularRepairs(config *config.Config) int {
	operatingTime := int(OperatingTimeForMajorRepairs(config)) % config.IndividualValues.Periodicity.Regular
	amount := float64((float32(operatingTime) + workmode.AnnualOperationMode(config))/float32(config.IndividualValues.Periodicity.Regular)) - float64(AmountOfMajorRepairs(config))
	return int(math.Round(float64(amount-0.5)))
}

func AmountOfTo2Repairs(config *config.Config) int {
	operatingTime := int(OperatingTimeForMajorRepairs(config)) % config.IndividualValues.Periodicity.To2
	amount := float64((float32(operatingTime) + workmode.AnnualOperationMode(config))/float32(config.IndividualValues.Periodicity.To2)) - float64(AmountOfMajorRepairs(config)) - float64(AmountOfRegularRepairs(config))
	return int(math.Round(float64(amount-0.5)))
}

func AmountOfTo1Repairs(config *config.Config) int {
	operatingTime := int(OperatingTimeForMajorRepairs(config)) % config.IndividualValues.Periodicity.To1
	amount := float64((float32(operatingTime) + workmode.AnnualOperationMode(config))/float32(config.IndividualValues.Periodicity.To1)) - float64(AmountOfMajorRepairs(config)) - float64(AmountOfRegularRepairs(config)) - float64(AmountOfTo2Repairs(config))
	return int(math.Round(float64(amount-0.5)))
}

func AmountOfEveryShiftRepair(config *config.Config) int {
	d_k := float32(config.SharedValues.CalendarDays)
	d_p := workmode.GeneralWithoutWorkWithRepair(config)
	k_zm := float32(config.IndividualValues.KoefZminnosti)
	amount := (d_k - d_p) * k_zm
	return int(math.Round(float64(amount - 0.5)))
}
