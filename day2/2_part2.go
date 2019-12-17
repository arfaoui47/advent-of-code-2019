package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const expected = "19690720"

func calc_first_element(file_content []byte, val1 string, val2 string) string {
	text := string(file_content)
	arr := strings.Split(text, ",")
	arr[1] = val1
	arr[2] = val2
	step := 4
	for i := 0; i < len(arr)-4; i += step {
		var arrint = []int{}
		for j := 0; j < step; j++ {
			var v int
			v, _ = strconv.Atoi(arr[i+j])
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

	return arr[0]
}

const n = 100

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading the file! %s", err)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// aux := make([]byte, len(content))
			// copy(aux, content)
			aux := content
			val1 := strconv.Itoa(i)
			val2 := strconv.Itoa(j)
			res := calc_first_element(aux, val1, val2)
			if res == expected {
				fmt.Println(i*100 + j)
			}
		}
	}
}
