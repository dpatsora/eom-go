package main

import (
	"fmt"
	"github.com/dpatsora/eom-go/pkg/config"
	"github.com/dpatsora/eom-go/pkg/maintenance"
	"github.com/dpatsora/eom-go/pkg/workmode"
	"github.com/dpatsora/eom-go/pkg/intensity"
)

func main() {
	config := &config.Config{}
	config.GetConfig()
	//
	fmt.Println("Неспрятливі умови:",workmode.BadWeather(config))
	fmt.Println("Річний режим роботи:",workmode.AnnualOperationMode(config))
	fmt.Println("Наробіток на капітальний ремонт:",maintenance.OperatingTimeForMajorRepairs(config))
	fmt.Println("Кількість капітальних ремонтів:",maintenance.AmountOfMajorRepairs(config))
	fmt.Println("Кількість поточних ремонтів:",maintenance.AmountOfRegularRepairs(config))
	fmt.Println("Кількість ТО 2 ремонтів:",maintenance.AmountOfTo2Repairs(config))
	fmt.Println("Кількість ТО 1 ремонтів:",maintenance.AmountOfTo1Repairs(config))
	fmt.Println("Кількість Щозмінних ремонтів:",maintenance.AmountOfEveryShiftRepair(config))
	fmt.Println("Трудомісткість Поточних ремонтів:",intensity.LaborIntensityOfRegularRepairs(config))

}