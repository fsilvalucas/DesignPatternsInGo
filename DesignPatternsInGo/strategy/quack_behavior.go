package strategy

import (
	"fmt"
)

type QuackBehavior interface {
	Quack()
}

type Quack struct{}
type NoQuack struct{}
type Squeak struct{}

func (q Quack) Quack() {
	fmt.Println("Quack Quack")
}

func (q NoQuack) Quack() {
	fmt.Println("Dude, i cant quack")
}

func (q Squeak) Quack() {
	fmt.Println("Squeak Squack!!")
}
