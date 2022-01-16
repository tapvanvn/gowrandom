package gowrandom_test

import (
	"testing"

	"github.com/tapvanvn/gowrandom"
)

func TestLastZeroWeightIn1000(t *testing.T) {

	randomer := gowrandom.MakeWRandom(5)
	for i := 0; i < 5; i++ {
		randomer.SetWeight(i, uint(4-i))
	}

	for i := 0; i < 1000; i++ {
		if randomer.Pick() == 4 {
			t.Fail()
		}
	}
}

func TestFirstZeroWeightIn1000(t *testing.T) {

	randomer := gowrandom.MakeWRandom(5)
	for i := 0; i < 5; i++ {
		randomer.SetWeight(i, uint(i))
	}

	for i := 0; i < 1000; i++ {
		if randomer.Pick() == 0 {
			t.Fail()
		}
	}
}

func TestNearLastZeroWeightIn1000(t *testing.T) {

	randomer := gowrandom.MakeWRandom(5)

	randomer.SetWeight(0, uint(5))
	randomer.SetWeight(1, uint(5))
	randomer.SetWeight(2, uint(5))
	randomer.SetWeight(3, uint(0))
	randomer.SetWeight(4, uint(5))

	for i := 0; i < 1000; i++ {
		if randomer.Pick() == 3 {
			t.Fail()
		}
	}
}
