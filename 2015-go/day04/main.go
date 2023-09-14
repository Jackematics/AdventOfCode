package main

import (
	"aoc_2015/day04/advent_coin_miner"
	"fmt"
)

func main() {
	result := advent_coin_miner.FindFirstValidMd5Hash("bgvyzdsv")

	fmt.Println("Solution: ", result)
}
