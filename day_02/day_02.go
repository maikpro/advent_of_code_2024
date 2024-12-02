package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/maikpro/advent_of_code_2024/shared"
)

/*
*
The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9

This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.
*/

type Report struct {
	levels []Level
}

type Level struct {
	value int
}

func CreateReports(lines []string) []Report {
	var reports []Report

	for _, line := range lines {
		var levels []Level

		// split line into levels
		for _, str := range strings.Fields(line) {
			value, err := strconv.Atoi(str)
			if err != nil {
				log.Println("Error converting string to int:", err)
				return nil
			}

			level := Level{value}
			levels = append(levels, level)
		}

		reports = append(reports, Report{levels})
	}

	return reports
}

func calculateDifferance(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

/*
*
conditions:
- The levels are either all increasing or all decreasing.
- Any two adjacent levels differ by at least one and at most three.
*/
func IsReportSafe(report Report) bool {
	isIncreasing := false
	for i := 0; i < len(report.levels)-1; i++ {
		currentLevel := report.levels[i]
		nextLevel := report.levels[i+1]

		diff := calculateDifferance(currentLevel.value, nextLevel.value)
		//log.Println("diff:", diff)

		if diff == 0 || diff > 3 {
			return false
		}

		// decide if increasing or decreasing by first two
		if report.levels[0].value < report.levels[1].value {
			isIncreasing = true
		}

		// The levels are either all increasing or all decreasing.
		if isIncreasing && currentLevel.value > nextLevel.value {
			return false
		}

		if !isIncreasing && currentLevel.value < nextLevel.value {
			return false
		}

	}
	return true
}

func countSafeReports(reports []Report) int {
	var counter int

	for _, report := range reports {
		if IsReportSafe(report) {
			counter++
		}
	}

	return counter
}

func CheckReports(filename string) int {
	lines := shared.ReadTextFile(filename)
	reports := CreateReports(lines)
	return countSafeReports(reports)
}

func main() {
	log.Println("Advent of Code 2024 - Day 02")

	result := CheckReports("/day_02/input.txt")
	log.Println("result:", result)
}
