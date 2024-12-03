package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reports := strings.Split(string(input), "\n")
	safeReports := 0
	for _, report := range reports[:len(reports)-1] {
		levels := toInt(strings.Split(report, " "))
		isDecreasing := levels[1] < levels[0]
		safe := true
		skippedLevel := false
		for i := 1; i < len(levels); i++ {
			safe = isSafe(levels[i], levels[i-1], isDecreasing); 
			if !safe && i == 1 {
				continue
			} else if !safe && !skippedLevel {
				fmt.Println("Checking skip for levels", levels, "at  index", i)
				safe = isSafe(levels[i], levels[i-2], isDecreasing)
				fmt.Println("Is safe skipping", levels[i-1], safe)
				if !safe {
					break
				}
				skippedLevel = true
			} else if !safe {
				break
			}
		}
		fmt.Println(levels, safe)
		if safe {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}

func isSafe(a, b int, isDecreasing bool) bool {
	diff := intAbsValue(a - b)
	if isDecreasing && a > b {
		return false	
	} else if !isDecreasing && a < b {
		return false
	} else if diff < 1 || diff > 3 {
		return false
	}
	return true
}

func intAbsValue(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func toInt(arr []string) []int {
	intArr := make([]int, len(arr))
	for i, v := range arr {
		intArr[i], _ = strconv.Atoi(v)
	}
	return intArr
}
