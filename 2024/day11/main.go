package main

import (
	"aoc_2024/input_reader"
	"aoc_2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func stoneWorker(blinks int, stones []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for blink := 1; blink <= blinks; blink++ {
		for i := len(stones) - 1; i >= 0; i-- {
			stoneStr := strconv.Itoa(stones[i])
			if stones[i] == 0 {
				stones[i] = 1
			} else if len(stoneStr) % 2 == 0 {
				leftStr := stoneStr[:(len(stoneStr) / 2)]
				rightStr := stoneStr[(len(stoneStr) / 2):]

				stones[i] = utils.ToInt(leftStr)
				stones = slices.Insert(stones, i + 1, utils.ToInt(rightStr))
			} else {
				stones[i] *= 2024
			}
		}
	}

	ch <- len(stones)
}

func partOne(input string, blinks int) int {
	rawStones := strings.Split(input, " ")
	stones := []int{}
	for _, rawStone := range rawStones {
		stones = append(stones, utils.ToInt(rawStone))
	}

	var wg sync.WaitGroup

	ch := make(chan int, len(stones))

	for _, stone := range stones {
		wg.Add(1)
		go stoneWorker(blinks, []int{stone}, ch, &wg)
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
