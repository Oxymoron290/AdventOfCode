package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lastDigit := 5
	var code []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lastDigit = getDigit(line, lastDigit)
		code = append(code, lastDigit)
	}

	fmt.Printf("Door code is %v", code)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getDigit(line string, startDigit int) int {
	x, y := digitToCoord(startDigit)
	for _, c := range line {
		switch c {
		case 'U':
			if y < 1 {
				y++
			}
		case 'D':
			if y > -1 {
				y--
			}
		case 'L':
			if x > -1 {
				x--
			}
		case 'R':
			if x < 1 {
				x++
			}
		}
	}

	return coordToDigit(x, y)
}

func coordToDigit(x int, y int) int {
	if x == -1 {
		if y == -1 {
			return 7
		}
		if y == 0 {
			return 4
		}
		if y == 1 {
			return 1
		}
	}
	if x == 0 {
		if y == -1 {
			return 8
		}
		if y == 0 {
			return 5
		}
		if y == 1 {
			return 2
		}
	}
	if x == 1 {
		if y == -1 {
			return 9
		}
		if y == 0 {
			return 6
		}
		if y == 1 {
			return 3
		}
	}
	return 0
}

func digitToCoord(digit int) (int, int) {
	switch digit {
	case 1:
		return -1, 1
	case 2:
		return 0, 1
	case 3:
		return 1, 1
	case 4:
		return -1, 0
	case 5:
		return 0, 0
	case 6:
		return 1, 0
	case 7:
		return -1, -1
	case 8:
		return 0, -1
	case 9:
		return 1, -1
	}
	return 0, 0
}
