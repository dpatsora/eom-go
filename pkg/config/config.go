package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	IndividualValues IndividualValues `yaml:"individual_values"`
	SharedValues SharedValues `yaml:"shared_values"`
}

type SharedValues struct {
	CalendarDays int `yaml:"calendar_days"`
	ShiftDuration float32 `yaml:"shift_duration"`
	Holidays int `yaml:"holidays"`
	HardWork int `yaml:"hard_work"`
	KV float32 `yaml:"k_v"`
}

type IndividualValues struct {
	MachineType string `yaml:"machine_type"`
	MachineAmount int `yaml:"machine_amount"`
	KoefZminnosti float32 `yaml:"k_zminnosti"`
	PCh float32 `yaml:"p_ch"`
	KoefNaprac float32 `yaml:"koef_naprac"`
	RelocationAmount float32 `yaml:"relocation_amount"`
	TRozvantazhennya float32 `yaml:"t_rozvantazhennya"`
	VTransportuvannya float32 `yaml:"v_transportuvannya"`
	Weather Weather `yaml:"weather"`
	Periodicity Periodicity `yaml:"periodicity"`
	LaborIntensity LaborIntensity `yaml:"labor_intensity"`
}

type Weather struct {
	TemperatureUnder30 float32 `yaml:"temperature_under_30"`
	WindGreaterThan10 float32 `yaml:"wind_greater_than_10"`
	Rain float32 `yaml:"rain"`
	ColdEarthn float32 `yaml:"cold_earth"`
}

type Periodicity struct {
	To1 int `yaml:"to_1"`
	To2 int `yaml:"to_2"`
	Regular int `yaml:"regular"`
	Capital int `yaml:"capital"`
}

type LaborIntensity struct {
	To1 int `yaml:"to_1"`
	To2 int `yaml:"to_2"`
	Season int `yaml:"season"`
	Regular int `yaml:"regular"`
}

func (c *Config) GetConfig() *Config {

	yamlFile, err := ioutil.ReadFile("task.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}