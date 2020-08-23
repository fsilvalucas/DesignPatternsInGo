package strategy

func Teste() {

	var my_new_duck = Duck{
		Fly_behavior:   FlyWithWings{},
		Quack_behavior: Quack{},
	}

	my_new_duck.Perform_fly()
	my_new_duck.Perform_quack()
}
