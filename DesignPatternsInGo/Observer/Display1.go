package Observer

import (
	"fmt"
	"time"
)

// Display1 implements Observer, Display
type Display1 struct {
	CurrentTemp Celsius
}


func (d Display1) Display(){
	hour, minutes, seconds := time.Now().Clock()
	currentUTCTimeString := fmt.Sprintf("%d:%02d:%d", hour, minutes, seconds)
	fmt.Println("Valor from Display1: ", d.CurrentTemp, "time: ", currentUTCTimeString )
}

func (d Display1) Update(observable Observable) {
	valor := observable.GetValor("Temperature")
	d.CurrentTemp = Celsius(valor)
	d.Display()
}
