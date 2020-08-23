package strategy

type Duck struct {
	Fly_behavior   FlyBehavior
	Quack_behavior QuackBehavior
}

func (duck *Duck) Perform_fly() {
	duck.Fly_behavior.Fly()
}

func (duck *Duck) Perform_quack() {
	duck.Quack_behavior.Quack()
}

func (duck *Duck) Set_fly_behavior(fb FlyBehavior) {
	duck.Fly_behavior = fb
}

func (duck *Duck) Set_quack_behavior(qb QuackBehavior) {
	duck.Quack_behavior = qb
}

