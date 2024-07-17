package main

import (
	"fmt"
	"testing"
)

type data struct {
	A   int
	B   int
	Sum int
}

func TestAdd(t *testing.T) {
	// check if addition works properly
	var inputs = []data{
		data{
			A:   5,
			B:   5,
			Sum: 10,
		},
		data{
			A:   5,
			B:   5,
			Sum: 11,
		},
	}

	for index, i := range inputs {
		t.Run(fmt.Sprintf("Test sum for %v", index), func(t *testing.T) {
			res := Addition(i.A, i.B)

			if res != i.Sum {
				t.Errorf("sum of %v and %v should be %v and not %v", i.A, i.B, i.Sum, res)
			}
		})
	}
}
