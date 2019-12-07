package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("shiit: %s", err)
	}

	scanner := bufio.NewScanner(file)
	var final_sum int32

	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		final_sum += int32(num/3) - 2
	}

	fmt.Println("fuel: ", final_sum)
}
