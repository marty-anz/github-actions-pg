package primes

import (
	"fmt"
	"testing"
)

func Test0IsNotPrime(t *testing.T) {
	ps := GetPrimesUpto(1)

	actual := len(ps)
	want := 0

	if actual != want {
		t.Fatalf(`len(ps) = %d want match for %d`, actual, want)
	}
}

func Test1IsNotPrime(t *testing.T) {
	ps := GetPrimesUpto(2)

	fmt.Println(ps)
	actual := ps[0]
	want := uint64(2)

	if actual != want {
		t.Fatalf(`first prime = %d want match for %d`, actual, want)
	}
}

func TestUpto10(t *testing.T) {
	ps := GetPrimesUpto(10)
	actual := len(ps)
	want := 4

	fmt.Println(ps)
	if actual != want {
		t.Fatalf(`len(ps) = %d want match for %d`, actual, want)
	}
}
