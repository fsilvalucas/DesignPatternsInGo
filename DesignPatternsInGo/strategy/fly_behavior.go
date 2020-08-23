package strategy

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct{}
type FlyWithRocket struct{}
type NoFlyNoWay struct{}

func (f FlyWithWings) Fly() {
	fmt.Println("Look at me!! I can fly")
}

func (f FlyWithRocket) Fly() {
	fmt.Println("Thats insane! I have a Rocket")
}

func (f NoFlyNoWay) Fly() {
	fmt.Println("Sorry i cant fly")
}
