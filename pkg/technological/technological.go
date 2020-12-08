package technological

import (
	"fmt"
	"github.com/dpatsora/eom-go/pkg/config"
	"math"
	"strconv"
	"strings"
)

func NominalTimeFund(config *config.Config) float32 {
	d_k := float32(config.SharedValues.CalendarDays)
	holidays := float32(config.SharedValues.Holidays)
	t_zm := float32(config.SharedValues.ShiftDuration)

	return (d_k - holidays) * t_zm
}

func RealTimeFund(config *config.Config) float32 {
	vacations := float32(config.SharedValues.Vacations)
	t_zm := float32(config.SharedValues.ShiftDuration)
	k_n := float32(0.98)

	return (NominalTimeFund(config) - vacations * t_zm) * k_n
}

func AnnualEquipmentTimeFund(config *config.Config) float32 {
	y_zm := float32(1)
	k_ob := float32(0.75)

	return NominalTimeFund(config)  * y_zm * k_ob
}

func PostsNumber(config *config.Config) map[string]int {
	workmodes:= config.IndividualValues.ATor
	postsAmount := make(map[string]int)
	k_m := float32(1)
	k_y := float32(1)
	y_zm := float32(1)
	nominalTimeFund := NominalTimeFund(config)

	//fmt.Println("Кількість постів")
	for key, element := range workmodes {
		index := key
		result_key := strings.Join([]string{"П", index}, "_")
		//fmt.Println(result_key, ":")
		var a = element["a"]
		var n = element["n"]
		var r = element["r"]
		//fmt.Println("a", a)
		//fmt.Println("n", n)
		//fmt.Println("r", r)

		P := (a * k_m * k_y) / (nominalTimeFund * n * y_zm * r)
		postsAmount[result_key] = int(math.Round(float64(P)))
	}
	return postsAmount
}

func NeededEmployeeNumber(config *config.Config) map[string]int {
	workmodes:= config.IndividualValues.ATor
	employees := make(map[string]int)
	realTimeFund := RealTimeFund(config)

	//fmt.Println("Кількість постів")
	for key, element := range workmodes {
		index := key
		result_key := strings.Join([]string{"Nш", index}, "_")
		//fmt.Println(result_key, ":")
		var a = element["a"]
		//fmt.Println("a", a)
		//fmt.Println("n", n)
		//fmt.Println("r", r)

		P := a / realTimeFund
		employees[result_key] = int(math.Round(float64(P)))
	}
	return employees
}

func TechnologicalEmployeeNumber(config *config.Config) map[string]int {
	workmodes:= config.IndividualValues.ATor
	employees := make(map[string]int)
	nominalTimeFund := NominalTimeFund(config)

	//fmt.Println("Кількість постів")
	for key, element := range workmodes {
		index := key
		result_key := strings.Join([]string{"Nя", index}, "_")
		//fmt.Println(result_key, ":")
		var a = element["a"]
		//fmt.Println("a", a)
		//fmt.Println("n", n)
		//fmt.Println("r", r)

		P := a / nominalTimeFund
		employees[result_key] = int(math.Round(float64(P)))
	}
	return employees
}


func EmployeeNumberByDepartment(config *config.Config) int {
	workmodes:= config.IndividualValues.ATor
	neededEmployeeNumber := NeededEmployeeNumber(config)

	employeesShtSum := float32(0)
	for _, value := range  neededEmployeeNumber {
		n := float32(value)
		employeesShtSum += n
	}

	employeesSum := float32(0) //

	//fmt.Println("Кількість постів")
	for _, element := range workmodes {
		n := element["n"]
		employeesSum += n
	}

	intermidiateSum := employeesSum + employeesShtSum

	employeesDopSum := math.Round(float64(intermidiateSum * 0.2))
	employeesItpSum := math.Round((float64(intermidiateSum) + employeesDopSum) * 0.1)
	employeesMopSum := math.Round((float64(intermidiateSum) + employeesDopSum + employeesItpSum ) * 0.04)

	fmt.Println("Сума шт: ",employeesShtSum )
	fmt.Println("Сума штрих: ",employeesSum )
	fmt.Println("Сума ДОП: ", employeesDopSum)
	fmt.Println("Сума ІТП: ", employeesItpSum)
	fmt.Println("Сума МОП: ", employeesMopSum)

	return 0
}


func Area(config *config.Config) map[string]int {
	posts:= PostsNumber(config)
	areas := make(map[string]int)

	//fmt.Println("Кількість постів")
	for key, element := range posts {
		index, _ := strconv.Atoi(strings.Split(key, "_")[1])

		if index <= 3 {
			result_key := strings.Join([]string{"F","к", strings.Split(key, "_")[1]}, "_")
			areas[result_key] = element * 140
			fmt.Println(result_key, ": ", areas[result_key])
		} else {
			result_key := strings.Join([]string{"F","г", strings.Split(key, "_")[1]}, "_")
			areas[result_key] = element * 120
			fmt.Println(result_key, ": ", areas[result_key])
		}
	}
	return areas
}
