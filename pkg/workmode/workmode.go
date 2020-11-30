package workmode

import (
	"github.com/dpatsora/eom-go/pkg/config"
)

func AnnualOperationMode(config *config.Config) float32 {
	d_k := float32(config.SharedValues.CalendarDays)
	t_zm := float32(config.SharedValues.ShiftDuration)
	k_zm := float32(config.IndividualValues.KoefZminnosti)
	p_ch := float32(config.IndividualValues.PCh)

	return ((d_k - GeneralWithoutWork(config)) * t_zm * k_zm)/(1 + p_ch * t_zm * k_zm)

}

func GeneralWithoutWork(config *config.Config) float32 {
	return WeekendsAndHolidays(config) + RelocationDays(config) + BadWeather(config) + Unpredictable(config) + RelocationToRepair(config)
}
func RelocationDays(config *config.Config) float32 {
	if config.IndividualValues.MachineType == "avtomobilniy_kran" {
		return 0
	}
	return 2
}

func BadWeather(config *config.Config) float32 {
	weather := config.IndividualValues.Weather
	var raw_d_m float32

	if config.IndividualValues.MachineType == "avtomobilniy_kran" {
		raw_d_m += weather.WindGreaterThan10
		raw_d_m += weather.TemperatureUnder30
		raw_d_m += weather.Rain
	}
	return 0.7 * raw_d_m
}

func WeekendsAndHolidays(config *config.Config) float32 {
	holidays := config.SharedValues.Holidays
	hard_work := config.SharedValues.HardWork

	return float32(holidays - hard_work)
}

func Unpredictable(config *config.Config) float32 {
	calendar_days := float32(config.SharedValues.CalendarDays)

	return 0.03 * (calendar_days - WeekendsAndHolidays(config))
}

func RelocationToRepair(config *config.Config) float32 {
	if config.IndividualValues.KoefNaprac < 0.7 {
		return 0
	}

	l := float32(150)
	v_tr := float32(config.IndividualValues.VTransportuvannya)
	y_zm := float32(1.5)
	t_tr := float32(y_zm * config.SharedValues.ShiftDuration)
	t_nr := float32(config.IndividualValues.TRozvantazhennya)
	t_p := float32(4.5)
	t_0 := float32(6)
	n_sp := float32(config.IndividualValues.MachineAmount)
	n_k := float32(1)

	return ((2 * l)/(t_tr * v_tr) + (2 * t_nr)/t_tr + t_p + t_0) * n_k / n_sp
}

func GeneralOnRepair(config * config.Config) float32  {
	d_k := float32(config.SharedValues.CalendarDays)
	k_zm := float32(config.IndividualValues.KoefZminnosti)
	p_ch := float32(config.IndividualValues.PCh)
	t_zm := float32(config.SharedValues.ShiftDuration)
	d_0 := GeneralWithoutWork(config)

	return ((d_k-d_0) * p_ch * t_zm * k_zm)/(1 + p_ch * t_zm * k_zm)
}

func GeneralWithoutWorkWithRepair(config * config.Config) float32{
	// D_p
	return GeneralWithoutWork(config) + GeneralOnRepair(config)
}