// Package weather provides tools to forecast the weather.
package weather

// CurrentCondition and is the current condition.
var CurrentCondition string
// CurrentLocation and is the current location.
var CurrentLocation string

// Forecast and returns the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
