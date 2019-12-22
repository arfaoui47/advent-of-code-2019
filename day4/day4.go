package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Unique(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func IntToSliceOfInt(n int) []int {
	var n_slice []int
	n_str := strconv.Itoa(n)
	n_str_arr := strings.Split(n_str, "")

	for _, c := range n_str_arr {
		c_int, _ := strconv.Atoi(c)
		n_slice = append(n_slice, c_int)
	}
	return n_slice
}

func IsIncreasing(n_slice []int) bool {
	sorted_n_slice := make([]int, len(n_slice))
	copy(sorted_n_slice, n_slice)
	sort.Ints(sorted_n_slice)
	return Equal(sorted_n_slice, n_slice)
}

func SliceContainElem(n_slice []int, elem int) bool {
	for _, slice_elem := range n_slice {
		if elem == slice_elem {
			return true
		}
	}
	return false
}

func ContainDupes(n_slice []int) bool {
	unique := Unique(n_slice)

	if len(unique) == len(n_slice) {
		return false
	}
	for _, unique_elem := range unique {
		count_elem := 0
		for _, elem := range n_slice {
			if unique_elem == elem {
				count_elem += 1
			}
		}
		if count_elem == 2 {
			return true
		}
	}
	return false
}

func main() {
	puzzle_start, puzzle_end := 248345, 746315
	i_count := 0
	for i := puzzle_start; i <= puzzle_end; i++ {
		i_slice := IntToSliceOfInt(i)
		if IsIncreasing(i_slice) && ContainDupes(i_slice) {
			i_count += 1
		}
	}
	fmt.Println(i_count)
}
