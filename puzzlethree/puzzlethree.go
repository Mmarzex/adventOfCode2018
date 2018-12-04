package puzzlethree

import (
	"fmt"
	"os"
)

type Claim struct {
	id, left, right, top, bottom int
	overlaps bool
}

func readInput(path string) []Claim {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	claims := []Claim{}

	for {
		c := Claim{}
		_, err = fmt.Fscanf(file, "#%d @ %d,%d: %dx%d", &c.id, &c.left, &c.top, &c.right, &c.bottom)
		if err != nil {
			return claims
		}
		c.right += c.left
		c.bottom += c.top
		c.overlaps = false
		claims = append(claims, c)
	}
	return claims
}

func RunPuzzle() {
	claims := readInput("inputPuzzleThree.txt")
	board := make([][]int, 1000)
	for i := range board {
		board[i] = make([]int, 1000)
	}

	for _, c := range claims {
		for w := c.left; w < c.right; w++ {
			for h := c.top; h < c.bottom; h++ {
				board[w][h]++
				if board[w][h] > 1 {
					c.overlaps = true
				}
			}
		}
	}

	sum := 0
	for _, l := range board {
		for _, v := range l {
			if v >= 2 {
				sum++
			}
		}
	}

	for _, c := range claims {
		alone := true
		for w := c.left; w < c.right; w++ {
			for h := c.top; h < c.bottom; h++ {
				if board[w][h] != 1 {
					alone = false
					break
				}
			}
			if !alone {
				break
			}
		}
		if alone {
			fmt.Println("Good Claim ", c.id)
			break
		}
	}

	fmt.Println("Square Inches ", sum)

}
