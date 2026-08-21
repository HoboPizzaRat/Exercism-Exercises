// Package weather does some cool stuff.
package weather



var (
	// CurrentCondition stores the current weather state.
	CurrentCondition string 
	// CurrentLocation Stores the city name.
	CurrentLocation  string	
)
// Forecast of the something.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
