package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"time"
)

func readInput() ([]int, []int, error) {
	rawFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer rawFile.Close()

	scanner := bufio.NewScanner(rawFile)

	left, right := []int{}, []int{}
	for scanner.Scan() {
		line := scanner.Text()

		parts := regexp.MustCompile(`\s{2,}`).Split(line, -1)

		lv, err := strconv.Atoi(parts[0])
		if err != nil {
			return []int{}, []int{}, err
		}
		left = append(left, lv)

		rv, err := strconv.Atoi(parts[1])
		if err != nil {
			return []int{}, []int{}, err
		}
		right = append(right, rv)
	}

	return left, right, nil
}

func distance(left, right int) int {
	diff := left - right

	if diff > 0 {
		return diff
	} else {
		return diff * -1
	}
}

func main() {
	start := time.Now()
	left, right, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += distance(left[i], right[i])
	}

	log.Printf("Sum is - %d\n", sum)

	rightCounts := make(map[int]int)
	for _, v := range right {
		rightCounts[v]++
	}

	similarityScore := 0
	for _, v := range left {
		similarityScore += v * rightCounts[v]
	}

	log.Printf("Similarity is - %d\n", similarityScore)

	log.Println(time.Since(start))
}
