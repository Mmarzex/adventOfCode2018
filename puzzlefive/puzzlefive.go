package puzzlefive

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func diffCase(a, b byte) bool {
	return math.Abs(float64(a)-float64(b)) == 32.0
	//return abs(int(a)-int(b)) == 32
}

func readInput(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(dat), "\n")
	return input[0]
}

func partOne(input string) int {
	i := 0
	for i < len(input) - 1 {
		if diffCase(input[i], input[i+1]) {
			if i == 0 {
				input = input[2:]
			} else if i == len(input) - 1 {
				input = input[:len(input) - 2]
				break
			} else {
				input = input[:i] + input[i+2:]
				i--
			}
			i = 0
		} else {
			i++
		}
	}
	fmt.Println(len(input))
	return len(input)
}

func partTwo() {
	input := readInput("inputPuzzleFive.txt")

	smallestLen := -1
	for c := 'a'; c <= 'z'; c++ {
		ss := strings.Replace(input, string(c), "", -1)
		ss = strings.Replace(ss, string(c - 32), "", -1)
		reactionLength := partOne(ss)
		if reactionLength < smallestLen || smallestLen == -1 {
			smallestLen = reactionLength
		}
	}
	fmt.Printf("Smallest Reaction %d\n", smallestLen)
}

func RunPuzzle() {
	input := readInput("inputPuzzleFive.txt")
	partOne(input)
	partTwo()
}
