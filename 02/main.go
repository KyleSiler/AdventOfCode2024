package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput(fileName string) [][]int {
	file, err := os.Open(fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	output := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")

		numSplits := []int{}
		for _, v := range splits {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			numSplits = append(numSplits, num)
		}

		output = append(output, numSplits)
	}

	return output
}

func part1(reports [][]int) {
	safeReports := 0
	for _, report := range reports {
		if isSafe(report) == -1 {
			safeReports++
		}
	}
	log.Printf("Safe Reports - %d", safeReports)
}

func isSafe(report []int) int {
	slope := report[0] - report[1]

	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]

		if diff == 0 {
			return i
		}
		if 1 <= diff && diff <= 3 {
			// safe
			if slope == 0 {
				slope = 1
			} else {
				if slope < 0 {
					return i
				}
			}
		} else if -1 >= diff && diff >= -3 {
			// safe
			if slope == 0 {
				slope = -1
			} else {
				if slope > 0 {
					return i
				}
			}
		} else {
			return i
		}
	}
	return -1
}

func part2(reports [][]int) {
	safeReports := 0

	for _, report := range reports {

		index := isSafe(report)

		if index == -1 {
			safeReports++
		} else {
			safe := false
			for i := 0; i < len(report); i++ {
				if isSafe(removeIndex(report, i)) == -1 {
					safe = true
					break
				}
			}
			if safe {
				safeReports++
			}
		}

	}
	log.Printf("Tolerating - %d", safeReports)
}

func removeIndex(report []int, index int) []int {
	if index < 0 {
		return report
	} else {
		left := report[:index]
		right := report[index+1:]
		return slices.Concat(left, right)
	}
}

func main() {
	fileName := os.Args[1]
	reports := readInput(fileName)

	part1(reports)
	part2(reports)
}
