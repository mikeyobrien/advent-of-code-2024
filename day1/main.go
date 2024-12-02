package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"math"
	"sort"
)


func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	list1 := []float64{}
	list2 := []float64{}
	for _, line := range(lines[:len(lines)-1]) {
		ids := strings.Split(line, "   ")
		id1, _ := strconv.ParseFloat(ids[0], 64)
		id2, _ := strconv.ParseFloat(ids[1], 64)
		list1 = append(list1, id1)
		list2 = append(list2, id2)
	}
	sort.Float64s(list1)
	sort.Float64s(list2)

	res := 0
	for i, _ := range(list1) {
		diff := math.Abs(list1[i] - list2[i])
		fmt.Println(list1[i], list2[i], diff)
		res += int(diff)
	}
	print(res)
}
