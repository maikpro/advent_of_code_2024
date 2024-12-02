package main

import (
	"testing"

	"github.com/maikpro/advent_of_code_2024/shared"
)

/*
*
parametrized tests for IsReportSafe
*/
func TestIsReportSafe(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		expected bool
	}{
		{"Safe because the levels are all decreasing by 1 or 2.", 0, true},
		{"Unsafe because 2 7 is an increase of 5.", 1, false},
		{"Unsafe because 6 2 is a decrease of 4.", 2, false},
		{"Unsafe because 1 3 is increasing but 3 2 is decreasing.", 3, false},
		{"Unsafe because 4 4 is neither an increase or a decrease.", 4, false},
		{"Safe because the levels are all increasing by 1, 2, or 3.", 5, true},
	}
	reports := CreateReports(shared.ReadTextFile("/example.txt"))

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsReportSafe(reports[tc.index])
			if got != tc.expected {
				t.Errorf("IsReportSafe(%d) = %t; want %t", tc.index, got, tc.expected)
			}
		})
	}
}

func TestCheckReports(t *testing.T) {
	got := CheckReports("/example.txt")
	want := 2

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}
