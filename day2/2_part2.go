package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("shiit: %s", err)
	}
	text := string(content)
	arr := strings.Split(text, ",")
	step := 4
	for i := 0; i < len(arr)-4; i += step {
		var arrint = []int{}
		for j := 0; j < step; j++ {
			var v int
			v, err = strconv.Atoi(arr[i+j])
			if err != nil {
				panic(err)
			}
			arrint = append(arrint, v)
		}

		var res int
		elem1, _ := strconv.Atoi(arr[arrint[1]])
		elem2, _ := strconv.Atoi(arr[arrint[2]])
		if arr[i] == "1" {
			res = elem1 + elem2
		} else if arr[i] == "2" {
			res = elem1 * elem2
		} else if arr[i] == "99" {
			break
		} else {
			fmt.Println("shouldn't happen!")
		}
		arr[arrint[3]] = strconv.Itoa(res)
	}
	fmt.Println(arr[0])
}
