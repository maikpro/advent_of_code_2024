package main

import "testing"

func TestCalculateDistanceBetweenTwoLists(t *testing.T) {
	got := CalculateDistanceBetweenTwoLists("example.txt")
	want := 11

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}
