package main

import (
	"fmt"
	"github.com/dpatsora/eom-go/pkg/config"
	"github.com/dpatsora/eom-go/pkg/workmode"
)

func main() {
	config := &config.Config{}
	config.GetConfig()

	fmt.Println("Річний режим роботи:",workmode.AnnualOperationMode(config))
}