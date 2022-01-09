package employee

type Employee struct {
	Id      int
	Counter int
}

func (e *Employee) Get() int {
	return (*e).Counter
}

func (e *Employee) Set(counter int) {
	(*e).Counter = counter
}
