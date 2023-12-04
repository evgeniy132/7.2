package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(ch chan<- []int) {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 5)
	for i := 0; i < 5; i++ {
		numbers[i] = rand.Intn(100)
	}
	ch <- numbers
}

func findMinMax(numbers []int, result chan<- [2]int) {
	min := numbers[0]
	max := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	result <- [2]int{min, max}
}

func printMinMax(result <-chan [2]int) {
	minMax := <-result
	fmt.Println("Min number", minMax[0])
	fmt.Println("Max number", minMax[1])
}

func main() {
	numbersChannel := make(chan []int)
	minMaxChannel := make(chan [2]int)

	go generateNumbers(numbersChannel)
	numbers := <-numbersChannel

	go findMinMax(numbers, minMaxChannel)
	go printMinMax(minMaxChannel)

	time.Sleep(1 * time.Second)
}
