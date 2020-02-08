package cointoss

import (
	"testing"
	"time"
)

func TestHeadsOrTailsIsReturned(t *testing.T) {
	var got int
	want := []int{0, 1}
	for i := 0; i < 20; i++ {
		got = Toss()
		time.Sleep(1 * time.Millisecond)
		if !contains(want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	}
}

func TestWeightingMoreOrLessWorks(t *testing.T) {
	var headsCount int
	var got int
	iterations := 20

	for i := 0; i < iterations; i++ {
		got = Toss()
		time.Sleep(1 * time.Millisecond)
		if got == 0 {
			headsCount++
		}
	}

	if headsCount == 0 || headsCount == iterations {
		t.Errorf("Weighting seems broke. Expected roughly half heads but got %v", headsCount)
	}
}

func contains(want []int, got int) bool {
	for _, n := range want {
		if got == n {
			return true
		}
	}
	return false
}
