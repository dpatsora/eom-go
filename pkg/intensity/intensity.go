package intensity

import (
	"github.com/dpatsora/eom-go/pkg/config"
	"github.com/dpatsora/eom-go/pkg/maintenance"
)

func LaborIntensityOfRegularRepairs(config *config.Config) float64 {
	return float64(config.IndividualValues.MachineAmount * maintenance.AmountOfRegularRepairs(config) * config.IndividualValues.LaborIntensity.Regular)
}

func LaborIntensityOfTo2Repairs(config *config.Config) float64 {
	return float64(config.IndividualValues.MachineAmount * maintenance.AmountOfTo2Repairs(config) * config.IndividualValues.LaborIntensity.To2)
}

func LaborIntensityOfTo1Repairs(config *config.Config) float64 {
	return float64(config.IndividualValues.MachineAmount * maintenance.AmountOfTo1Repairs(config) * config.IndividualValues.LaborIntensity.To1)
}

func LaborIntensityOfSeasonRepairs(config *config.Config) float64 {
	amountOfSeasonRepairs := 2
	return float64(config.IndividualValues.MachineAmount * amountOfSeasonRepairs * config.IndividualValues.LaborIntensity.Season)
}

func LaborIntensityOfEveryShiftRepairs(config *config.Config) float64 {
	intensityOfEveryShiftRepairs := 1
	return float64(config.IndividualValues.MachineAmount * maintenance.AmountOfEveryShiftRepair(config) * intensityOfEveryShiftRepairs)
}

func GeneralLaborIntensity(config *config.Config) float64 {
	return (LaborIntensityOfTo2Repairs(config) + LaborIntensityOfTo1Repairs(config) + LaborIntensityOfSeasonRepairs(config) + LaborIntensityOfEveryShiftRepairs(config))
}
