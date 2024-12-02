package utils

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
