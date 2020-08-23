package Observer

import (
	"fmt"
	"time"
)

// Display3 implements Observer, Display
type Display3 struct {
	CurrentPressure Pressure
	CurrentHumidity Humidity
}


func (d Display3) Display(){
	hour, minutes, seconds := time.Now().Clock()
	currentUTCTimeString := fmt.Sprintf("%d:%02d:%d", hour, minutes, seconds)
	fmt.Println("Valor from Display3: \tPressure: ", d.CurrentPressure, "Humidity: ",
		d.CurrentHumidity, "time: ", currentUTCTimeString )
}

func (d Display3) Update(observable Observable) {
	pressure := observable.GetValor("Pressure")
	humidity := observable.GetValor("Humidity")

	d.CurrentHumidity = Humidity(humidity)
	d.CurrentPressure = Pressure(pressure)

	d.Display()
}
