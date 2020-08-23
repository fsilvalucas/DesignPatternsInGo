package Observer

// Observable interface to work with Observer
type Observable interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObserver()
	GetValor(str string) float64
}
