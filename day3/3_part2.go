package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func build_path(wire []string) []point {
	var path []point
	int_point := point{x: 0, y: 0}
	prec_x := int_point.x
	prec_y := int_point.y
	for _, instruction := range wire {
		direction := string(instruction[0])
		value, _ := strconv.Atoi(instruction[1:])
		switch direction {
		case "R":
			for i := 0; i < value; i++ {
				path = append(path, point{x: prec_x + i, y: prec_y})
			}
			prec_x = prec_x + value
		case "L":
			for i := 0; i < value; i++ {
				path = append(path, point{x: prec_x - i, y: prec_y})
			}
			prec_x = prec_x - value
		case "U":
			for i := 0; i < value; i++ {
				path = append(path, point{x: prec_x, y: prec_y + i})
			}
			prec_y = prec_y + value
		case "D":
			for i := 0; i < value; i++ {
				path = append(path, point{x: prec_x, y: prec_y - i})
			}
			prec_y = prec_y - value
		}
	}
	return path

}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("shiit: %s", err)

	}
	text := string(content)
	input := strings.Split(text, "\n")
	wire1 := strings.Split(input[0], ",")
	wire2 := strings.Split(input[1], ",")

	path1 := build_path(wire1)
	path2 := build_path(wire2)
	min_i := 1000000
	min_j := 1000000
	for i, p := range path1 {
		for j, x := range path2 {
			if p.x == 0 && p.y == 0 {
				continue
			}
			if p == x {
				if i+j < min_i+min_j {
					min_i = i
					min_j = j
				}
			}
		}
	}
	fmt.Println(min_i + min_j)
}
