package main

import "testing"

func TestCalculateDistanceBetweenTwoLists(t *testing.T) {
	got := CalculateDistanceBetweenTwoLists("example.txt")
	want := 11

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func TestCalculateSimilarityScore(t *testing.T) {
	got := CalculateSimilarityScore("example.txt")
	want := 31

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}
