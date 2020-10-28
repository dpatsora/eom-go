package main

import (
	"fmt"
	"github.com/dpatsora/pkg/conf"
)

func main() {
	var c conf
	c.getConf()

	fmt.Println(c)
	fmt.Println(c.IndividualValues)
	//fmt.Println(c.IndividualValues.weather.temperature_under_30)
}