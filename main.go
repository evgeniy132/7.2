package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(ch chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		ch <- rand.Intn(100)
	}
	close(ch)
}

func findMinMax(numbers <-chan int, result chan<- [2]int) {
	min := <-numbers
	max := min
	for num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	result <- [2]int{min, max}
	close(result)
}

func printMinMax(result <-chan [2]int, done chan<- bool) {
	defer close(done)
	minMax, isOpen := <-result
	if !isOpen {
		return
	}
	fmt.Println("Min number", minMax[0])
	fmt.Println("Max number", minMax[1])
}

func main() {
	numbersChannel := make(chan int)
	minMaxChannel := make(chan [2]int)
	done := make(chan bool)

	go generateNumbers(numbersChannel)
	go findMinMax(numbersChannel, minMaxChannel)
	go printMinMax(minMaxChannel, done)

	<-done
}
