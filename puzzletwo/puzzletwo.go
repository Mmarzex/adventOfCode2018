package puzzletwo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readAndParseFile(path string) [][]string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dataString := string(dat)
	combos := strings.Split(dataString, "\n")
	ret := make([][]string, 0)
	for _, v := range combos[0 : len(combos)-1] {
		ret = append(ret, strings.Split(v, ""))
	}
	return ret
}

func RunPuzzle() {
	input := readAndParseFile("inputPuzzleTwo.txt")
	fmt.Println(input)
	twoCodeOccurrences := 0
	threeCodeOccurrences := 0
	for _, code := range input {
		occurrences := map[string]int{}
		for _, c := range code {
			_, ok := occurrences[c]
			if ok {
				occurrences[c] = occurrences[c] + 1
			} else {
				occurrences[c] = 1
			}
		}
		fmt.Println(occurrences)
		twoFound := false
		threeFound := false
		for _, v := range occurrences {
			if !twoFound && v == 2 {
				twoCodeOccurrences += 1
				twoFound = true
			} else if !threeFound && v == 3 {
				threeCodeOccurrences += 1
				threeFound = true
			}
		}
	}
	fmt.Println("Checksum ", twoCodeOccurrences*threeCodeOccurrences)

	for _, codeOne := range input {
		for _, codeTwo := range input {
			diff := 0
			match := make([]string, 0)

			for i, v := range codeOne {
				if v == codeTwo[i] {
					match = append(match, v)
				}
				if v != codeTwo[i] {
					diff += 1
				}
				if diff >= 2 {
					break
				}
			}
			//fmt.Println(fmt.Sprintf("Diff %d", diff))
			if diff == 1 {
				fmt.Println("Found Codes")
				fmt.Println(codeOne)
				fmt.Println(codeTwo)
				fmt.Println("Matched Code ", strings.Join(match, ""))
				break
			}

		}
	}
}
