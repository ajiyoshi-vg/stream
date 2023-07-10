package stream

type Thunk[T any] func() T

type Cell[T, S any] struct {
	car T
	cdr Thunk[S]
}

func Cons[T, S any](car T, cdr Thunk[S]) *Cell[T, S] {
	return &Cell[T, S]{
		car: car,
		cdr: cdr,
	}
}
func (x Cell[T, S]) Car() T {
	return x.car
}
func (x Cell[T, S]) Cdr() S {
	return x.cdr.Force()
}

func Car[T, S any](x Cell[T, S]) T {
	return x.Car()
}
func Cdr[T, S any](x Cell[T, S]) S {
	return x.Cdr()
}

func (x Thunk[T]) Force() T {
	return x()
}
