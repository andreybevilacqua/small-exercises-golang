package main

import (
	"fmt"
	"sort"
)

const incValue = 1

func main() {
	input := []int{1, 3, 3, 3, 2, 2, 2, 2, 3, 0, 7, 7, 7, 7, 8, 9, 9, 9, 9, 9}
	m := createMap(input)

	fmt.Println("Result:", sortMap(m))
}

func createMap(input []int) map[int]int {
	m := make(map[int]int)
	counter := 1

	for i := 0; i < len(input); i++ {
		v, ok := m[input[i]]
		if ok {
			counter = v + incValue
		}
		m[input[i]] = counter
		counter = 1
	}
	return m
}

func sortMap(m map[int]int) []pair {
	var p []pair
	for k, v := range m {
		p = append(p, pair{v, k})
	}
	sort.Sort(pairList(p))

	var result []pair
	for i := 0; i < len(p); i++ {
		result = append(result, pair{p[i].v2, p[i].v1})
	}
	return result
}
