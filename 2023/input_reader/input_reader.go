package input_reader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var lines []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}
