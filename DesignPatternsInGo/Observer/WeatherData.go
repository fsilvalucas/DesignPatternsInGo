package Observer

// WeatherData implements Observable
type WeatherData struct {
	temperature Celsius
	humidity    Humidity
	pressure    Pressure
	observers   []*Observer
}

// Register an Observer at our slice
func (wd *WeatherData) RegisterObserver(observer Observer){
	wd.observers = append(wd.observers, &observer)
}

// Remove a Observer of our slice
func (wd *WeatherData) RemoveObserver(observer Observer){
	for i, o := range wd.observers {
		if *o == observer {
			wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)

		}
	}
}

// Notidy all Observer in slice
func (wd *WeatherData) NotifyObserver() {
	for _, observer := range wd.observers {
		var o Observer = *observer
		o.Update(wd)
	}
}

// Return a value specified by str String - For Observer pull the value
func (wd *WeatherData) GetValor(str string) float64{
	switch str {
	case "Temperature":
		return float64(wd.temperature)
	case "Pressure":
		return float64(wd.pressure)
	case "Humidity":
		return float64(wd.humidity)
	default:
		return -1
	}
}

// Just for test, you can replace that to take real Data from somewhere at web
func (wd *WeatherData) SetMeasures(temp float64, hum float64, press float64) {
	wd.temperature = Celsius(temp)
	wd.humidity = Humidity(hum)
	wd.pressure = Pressure(press)
	wd.NotifyObserver()
}
