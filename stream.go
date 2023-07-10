package stream

import (
	"golang.org/x/exp/constraints"
)

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

func Scale[T Number](x T, s Stream[T]) Stream[T] {
	return Map(
		func(y T) T { return x * y },
		s,
	)
}

func PartialSum[T Number](s Stream[T]) Stream[T] {
	return Cons(
		s.Car(),
		func() Stream[T] {
			return AddStream(s.Cdr(), PartialSum(s))
		},
	)
}

func PISummands(n float64) Stream[float64] {
	negate := func(x float64) float64 { return -x }
	return Cons(
		1/n,
		func() Stream[float64] {
			return Map(negate, PISummands(n+2))
		},
	)
}

var PIStream = Scale(4.0, PartialSum(PISummands(1)))

func square[T Number](x T) T {
	return x * x
}

func EulerTransform[T constraints.Float](s Stream[T]) Stream[T] {
	s0 := s.Car()
	s1 := s.Cdr().Car()
	s2 := s.Cdr().Cdr().Car()
	return Cons(
		s2-square(s2-s1)/(s0-2*s1+s2),
		func() Stream[T] {
			return EulerTransform(s.Cdr())
		},
	)
}

type Transform[T any] func(Stream[T]) Stream[T]

func Tableau[T any](t Transform[T], s Stream[T]) Stream[Stream[T]] {
	return Cons(
		s,
		func() Stream[Stream[T]] {
			return Tableau(t, t(s))
		},
	)
}

func AcceleratedSquence[T any](t Transform[T], s Stream[T]) Stream[T] {
	car := func(x Stream[T]) T {
		return x.Car()
	}
	return Map(car, Tableau(t, s))
}
