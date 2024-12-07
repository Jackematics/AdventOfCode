package utils

import (
	"math"
	"strconv"
)

func Abs(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

func RemoveAt(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func ToInt(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}

func MiddleIndex[T any](slice []T) int {
	return int(math.Floor(float64((len(slice) + 1) / 2))) - 1
}