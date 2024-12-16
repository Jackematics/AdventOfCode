package main

import (
	"aoc_2024/input_reader"
	"aoc_2024/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// func stoneWorker(blinks int, stone int, sum int) int {
// 	if blinks == 0 {
// 		return sum + 1
// 	}

// 	stoneStr := strconv.Itoa(stone)
// 	if stone == 0 {
// 		return stoneWorker(blinks - 1, 1, sum)
// 	} else if len(stoneStr) % 2 == 0 {
// 		leftStr := stoneStr[:(len(stoneStr) / 2)]
// 		rightStr := stoneStr[(len(stoneStr) / 2):]

// 		sum = stoneWorker(blinks - 1, utils.ToInt(leftStr), sum)
// 		return stoneWorker(blinks - 1, utils.ToInt(rightStr), sum)
// 	} else {
// 		return stoneWorker(blinks - 1, stone * 2024, sum)
// 	}
// }

func stoneWorker(blinks int, stone int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	if blinks == 0 {
		ch <- 1
		return
	}
	
	stoneStr := strconv.Itoa(stone)
	if stone == 0 {
		wg.Add(1)
		go stoneWorker(blinks - 1, 1, ch, wg)
	} else if len(stoneStr) % 2 == 0 {
		wg.Add(2)
		leftStr := stoneStr[:(len(stoneStr) / 2)]
		rightStr := stoneStr[(len(stoneStr) / 2):]

		go stoneWorker(blinks - 1, utils.ToInt(leftStr), ch, wg)
		go stoneWorker(blinks - 1, utils.ToInt(rightStr), ch, wg)
	} else {
		wg.Add(1)
		go stoneWorker(blinks - 1, stone * 2024, ch, wg)
	}
}

func partOne(input string, blinks int) int {
	rawStones := strings.Split(input, " ")
	stones := []int{}
	for _, rawStone := range rawStones {
		stones = append(stones, utils.ToInt(rawStone))
	}

	var wg sync.WaitGroup

	ch := make(chan int, 100_000)

	go func() {
		for {
			time.Sleep(100_000)
			sum := 0
		Loop: 
			for {
				select {
				case val := <-ch:
					sum += val
				default:
					break Loop
				}

			}

			if sum > 0 {
				ch <- sum
			}
		}
	}()


	for _, stone := range stones {
		wg.Add(1)
		go stoneWorker(blinks, stone, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for output := range ch {
		sum += output
	}

	return sum
}


func main() {
	exampleOne := "125 17"

	exampleOneBlink := partOne(exampleOne, 1)
	fmt.Println("Example Result Part One One Blink: " + strconv.Itoa(exampleOneBlink))
	
	exampleTwoBlinks := partOne(exampleOne, 2)
	fmt.Println("Example Result Part One Two Blinks: " + strconv.Itoa(exampleTwoBlinks))

	exampleThreeBlinks := partOne(exampleOne, 3)
	fmt.Println("Example Result Part One Three Blinks: " + strconv.Itoa(exampleThreeBlinks))

	exampleFourBlinks := partOne(exampleOne, 4)
	fmt.Println("Example Result Part One Four Blinks: " + strconv.Itoa(exampleFourBlinks))

	exampleFiveBlinks := partOne(exampleOne, 5)
	fmt.Println("Example Result Part One Five Blinks: " + strconv.Itoa(exampleFiveBlinks))

	exampleSixBlinks := partOne(exampleOne, 6)
	fmt.Println("Example Result Part One Six Blinks: " + strconv.Itoa(exampleSixBlinks))

	exampleTwentyFiveBlinks := partOne(exampleOne, 25)
	fmt.Println("Example Result Part One Twenty Five Blinks: " + strconv.Itoa(exampleTwentyFiveBlinks))

	input, _ := input_reader.ReadInput("./input.txt")

	partOneResult := partOne(input[0], 25)
	fmt.Println("Part One Result: " + strconv.Itoa(partOneResult))

	partTwoResult := partOne(input[0], 75)
	fmt.Println("Part Two Result: " + strconv.Itoa(partTwoResult))
}
