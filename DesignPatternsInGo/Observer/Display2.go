package Observer

import (
	"fmt"
	"time"
)

// Display2 implements Observer, Display
// Equal to Display1 but we gonna parse Celsius to Fahrenheit
type Display2 struct {
	CurrentTemp Fahrenheit
}

func (d Display2) Display() {
	hour, minutes, seconds := time.Now().Clock()
	currentUTCTimeString := fmt.Sprintf("%d:%02d:%d", hour, minutes, seconds)
	fmt.Println("Valor from Display2: ", d.CurrentTemp, "time: ", currentUTCTimeString )
}


func (d Display2) Update(observable Observable) {
	valorInCelsius := Celsius(observable.GetValor("Temperature"))
	d.CurrentTemp = CToF(valorInCelsius)
	d.Display()
}