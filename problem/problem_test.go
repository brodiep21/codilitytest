package problem_test

import (
	"testing"

	"github.com/brodiep21/codilitytest/problem"
)

func TestSolution(t *testing.T) {
	added := "babaa"
	got := problem.Problem(added)
	want := 3

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
func TestSolution2(t *testing.T) {
	added := "bbbaaabbb"
	got := problem.Problem(added)
	want := 0

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
func TestSolution3(t *testing.T) {
	added := "bbbab"
	got := problem.Problem(added)
	want := 4

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
