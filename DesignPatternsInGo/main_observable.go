package main

import "github.com/fsilvalucas/DesignPatternsInGo/Observer"

func main() {
	// create our objects
	var wd = new(Observer.WeatherData)
	var d = new(Observer.Display1)
	var d2 = new(Observer.Display2)
	var d3 = new(Observer.Display3)

	// Register our Observer
	wd.RegisterObserver(d)
	wd.RegisterObserver(d2)
	wd.RegisterObserver(d3)

	// Test the SetMeasures
	wd.SetMeasures(24.3, 32, 78.4)
	wd.SetMeasures(25.2, 33, 79.9)
	wd.SetMeasures(26, 34, 77)
}
