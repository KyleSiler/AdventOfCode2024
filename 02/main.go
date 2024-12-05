package main

import (
	"bufio"
	"log"
	"os"
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
		slope := 0

		safe := true

		for i := 0; i < len(report)-1; i++ {
			diff := report[i] - report[i+1]

			if diff == 0 {
				safe = false
				break
			}
			if 1 <= diff && diff <= 3 {
				// safe
				if slope == 0 {
					slope = 1
				} else {
					if slope < 0 {
						safe = false
						break
					}
				}
			} else if -1 >= diff && diff >= -3 {
				// safe
				if slope == 0 {
					slope = -1
				} else {
					if slope > 0 {
						safe = false
						break
					}
				}
			} else {
				safe = false
				break
			}
		}

		if safe {
			safeReports++
		}
	}
	log.Printf("Safe Reports - %d", safeReports)
}

func part2(reports [][]int) {
	safeReports := 0

	for _, report := range reports {
		safe := true
		slope := 0

		//

	}

	log.Printf("Tolerating - %d", safeReports)
}

func main() {
	fileName := os.Args[1]
	reports := readInput(fileName)

	part1(reports)
	part2(reports)
}
