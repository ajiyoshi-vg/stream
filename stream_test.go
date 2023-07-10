package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func times(n int) func(int) int {
	return func(x int) int {
		return n * x
	}
}

func TestStream(t *testing.T) {
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
	}

	for title, c := range cases {
		t.Run(title, func(t *testing.T) {
			actual := Take(c.n, c.input)
			assert.Equal(t, c.expect, actual)
		})
	}
}
