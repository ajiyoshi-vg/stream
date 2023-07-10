package stream

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func times(n int) func(int) int {
	return func(x int) int {
		return n * x
	}
}

func TestIntStream(t *testing.T) {
	cases := map[string]struct {
		input  Stream[int]
		n      int
		expect []int
	}{
		"take 5": {
			input:  Natural,
			n:      5,
			expect: []int{0, 1, 2, 3, 4},
		},
		"AddStream": {
			input:  AddStream(Natural, Natural),
			n:      5,
			expect: []int{0, 2, 4, 6, 8},
		},
		"Map": {
			input:  Map(times(3), Natural),
			n:      5,
			expect: []int{0, 3, 6, 9, 12},
		},
		"GenerateFib": {
			input:  GenerateFib(0, 1),
			n:      8,
			expect: []int{0, 1, 1, 2, 3, 5, 8, 13},
		},
		"GenerateFib2": {
			input:  GenerateFib2(),
			n:      8,
			expect: []int{0, 1, 1, 2, 3, 5, 8, 13},
		},
		"PartialSum": {
			input:  PartialSum(Natural),
			n:      6,
			expect: []int{0, 1, 3, 6, 10, 15},
		},
		"Scale": {
			input:  Scale(2, Natural),
			n:      5,
			expect: []int{0, 2, 4, 6, 8},
		},
	}

	for title, c := range cases {
		t.Run(title, func(t *testing.T) {
			actual := Take(c.n, c.input)
			assert.Equal(t, c.expect, actual)
		})
	}
}

func TestFloat(t *testing.T) {
	cases := map[string]struct {
		input  Stream[float64]
		n      int
		expect []float64
	}{
		"PIStream": {
			input:  PIStream,
			n:      5,
			expect: []float64{4.0, 2.66, 3.46, 2.89, 3.33},
		},
		"EulerTransformed": {
			input:  EulerTransform(PIStream),
			n:      10,
			expect: []float64{3.16, 3.133, 3.145, 3.139, 3.142, 3.140, 3.142, 3.14254},
		},
		"AcceleratedSquence": {
			input:  AcceleratedSquence(EulerTransform[float64], PIStream),
			n:      8,
			expect: []float64{4, 3.166, 3.142, 3.14159, 3.141592, 3.1415926, 3.14159265, 3.1415926535},
		},
	}

	closeEnough := func(a, b float64) bool {
		return math.Abs(a/b-1) < 0.01
	}

	for title, c := range cases {
		t.Run(title, func(t *testing.T) {
			actuals := Take(c.n, c.input)
			for i, expect := range c.expect {
				actual := actuals[i]
				if !closeEnough(expect, actual) {
					assert.Equal(t, c.expect, actuals)
				}
			}
		})
	}
}
