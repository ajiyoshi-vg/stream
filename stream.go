package stream

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

type Stream[T any] interface {
	Car() T
	Cdr() Stream[T]
}

var (
	_ Stream[int] = (*Cell[int, Stream[int]])(nil)
)

func Take[T any](n int, s Stream[T]) []T {
	ret := make([]T, 0, n)
	for x := s; x != nil && n > 0; x = x.Cdr() {
		n -= 1
		ret = append(ret, x.Car())
	}
	return ret
}

func Map[T, S any](f func(T) S, s Stream[T]) Stream[S] {
	return Cons(
		f(s.Car()),
		func() Stream[S] {
			cdr := s.Cdr()
			if cdr == nil {
				return nil
			}
			return Map(f, cdr)
		},
	)
}

func AddStream[T Number](a, b Stream[T]) Stream[T] {
	return Cons(
		a.Car()+b.Car(),
		func() Stream[T] {
			x, y := a.Cdr(), b.Cdr()
			if x == nil {
				return y
			}
			if y == nil {
				return x
			}
			return AddStream(x, y)
		},
	)
}

var Natural = IntegerStartingFrom(0)

func IntegerStartingFrom(n int) Stream[int] {
	return Cons(
		n,
		func() Stream[int] {
			return IntegerStartingFrom(n + 1)
		},
	)
}

func GenerateFib(a, b int) Stream[int] {
	return Cons(
		a,
		func() Stream[int] {
			return GenerateFib(b, a+b)
		},
	)
}

func GenerateFib2() Stream[int] {
	return Cons(
		0,
		func() Stream[int] {
			return Cons(
				1,
				func() Stream[int] {
					fibs := GenerateFib2()
					return AddStream(fibs.Cdr(), fibs)
				},
			)
		},
	)
}
