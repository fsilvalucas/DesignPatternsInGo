package main

import (
	"fmt"

	"github.com/fsilvalucas/DesignPatternsInGo/strategy"
)

func main() {

	var Loud_Duck = strategy.Duck{
		Fly_behavior:   strategy.FlyWithWings{},
		Quack_behavior: strategy.Quack{},
	}

	var Space_Duck = strategy.Duck{
		Fly_behavior:   strategy.FlyWithRocket{},
		Quack_behavior: strategy.Squeak{},
	}

	var boring_Duck = strategy.Duck{
		Fly_behavior:   strategy.NoFlyNoWay{},
		Quack_behavior: strategy.NoQuack{},
	}

	Loud_Duck.Perform_fly()
	Loud_Duck.Perform_quack()

	Space_Duck.Perform_fly()
	Space_Duck.Perform_quack()

	boring_Duck.Perform_fly()
	boring_Duck.Perform_quack()

	fmt.Println("Lets help boring Duck")

	boring_Duck.Set_fly_behavior(strategy.FlyWithRocket{})
	boring_Duck.Set_quack_behavior(strategy.Quack{})

	boring_Duck.Perform_fly()
	boring_Duck.Perform_quack()

	fmt.Println("Now Boring Duck is a Cooooool Duck")

}
