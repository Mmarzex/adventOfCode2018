package puzzleone

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInInput(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat), err
}

func RunPuzzle(inputPath string) {
	input, _ := ReadInInput(inputPath)
	parsedInput := strings.Split(input, "\n")
	currentFrequency := 0
	for _, v := range parsedInput[0 : len(parsedInput)-1] {
		change, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		currentFrequency += change
	}
	fmt.Println(fmt.Sprintf(" Part One Current Frequency %d", currentFrequency))
	freqs := map[int]bool{}
	newFrequency := 0
	found := false
	for !found {
		for _, v := range parsedInput[0 : len(parsedInput)-1] {
			change, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			newFrequency += change
			_, ok := freqs[newFrequency]
			if ok {
				found = true
				break
			}
			freqs[newFrequency] = true
		}
	}
	fmt.Println(fmt.Sprintf("Part Two %d", newFrequency))

}
