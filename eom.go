package main

import (
	"fmt"

	"github.com/dpatsora/eom-go/pkg/config"
	"github.com/dpatsora/eom-go/pkg/intensity"
	"github.com/dpatsora/eom-go/pkg/maintenance"
	"github.com/dpatsora/eom-go/pkg/technological"
	"github.com/dpatsora/eom-go/pkg/workmode"
)

func main() {
	config := &config.Config{}
	config.GetConfig()

	fmt.Println("Річний режим роботи:", workmode.AnnualOperationMode(config))
	fmt.Println("Простой:", workmode.GeneralWithoutWork(config))
	fmt.Println("Вихідні і  святкові:", workmode.WeekendsAndHolidays(config))
	fmt.Println("Неспрятливі умови:", workmode.BadWeather(config))
	fmt.Println("Непередбачувані:", workmode.Unpredictable(config))
	fmt.Println("Перевезення на ремонт:", workmode.RelocationToRepair(config))

	fmt.Println("Наробіток на капітальний ремонт:", maintenance.OperatingTimeForMajorRepairs(config))
	fmt.Println("Кількість капітальних ремонтів:", maintenance.AmountOfMajorRepairs(config))
	fmt.Println("Кількість поточних ремонтів:", maintenance.AmountOfRegularRepairs(config))
	fmt.Println("Кількість ТО 2 ремонтів:", maintenance.AmountOfTo2Repairs(config))
	fmt.Println("Кількість ТО 1 ремонтів:", maintenance.AmountOfTo1Repairs(config))
	fmt.Println("Кількість Щозмінних ремонтів:", maintenance.AmountOfEveryShiftRepair(config))
	fmt.Println("Трудомісткість Поточних ремонтів:", intensity.LaborIntensityOfRegularRepairs(config))
	fmt.Println("Трудомісткість ТО1 ремонтів:", intensity.LaborIntensityOfTo1Repairs(config))
	fmt.Println("Трудомісткість ТО2 ремонтів:", intensity.LaborIntensityOfTo2Repairs(config))
	fmt.Println("Трудомісткість Сезонних ремонтів:", intensity.LaborIntensityOfSeasonRepairs(config))
	fmt.Println("Трудомісткість Щозмінних ремонтів:", intensity.LaborIntensityOfEveryShiftRepairs(config))
	fmt.Println("Трудомісткість Загальна:", intensity.GeneralLaborIntensity(config))

	fmt.Println("Номінальний Фонд часу:", technological.NominalTimeFund(config))
	fmt.Println("Реальний Фонд часу:", technological.RealTimeFund(config))
	fmt.Println("Річний Фонд часу Обладнання:", technological.AnnualEquipmentTimeFund(config))

	fmt.Println("Кількість Постів:")
	for key, element := range technological.PostsNumber(config) {
		fmt.Println(key, ": ", element)
	}

	fmt.Println("Штатна кількість робітників:")
	for key, element := range technological.NeededEmployeeNumber(config) {
		fmt.Println(key, ": ", element)
	}

	fmt.Println("Технологічно необхідна кількість робітників:")
	for key, element := range technological.TechnologicalEmployeeNumber(config) {
		fmt.Println(key, ": ", element)
	}

	technological.EmployeeNumberByDepartment(config)
	technological.Area(config)

}
