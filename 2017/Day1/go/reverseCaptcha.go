package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = process2(scanner.Text())
	}

	fmt.Printf("The solution is: %v", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func process(line string) int {
	result := 0
	for i, c := range line {
		curr, er := strconv.Atoi(string(c))
		if er != nil {
			log.Fatal(er)
		}
		var partner int
		if i+1 == len(line) {
			partner, er = strconv.Atoi(string(line[0]))
			if er != nil {
				log.Fatal(er)
			}
		} else {
			partner, er = strconv.Atoi(string(line[i+1]))
			if er != nil {
				log.Fatal(er)
			}
		}

		if curr == partner {
			result += curr
		}
	}
	return result
}

func process2(line string) int {
	step := len(line) / 2
	result := 0
	for i, c := range line {
		curr, er := strconv.Atoi(string(c))
		if er != nil {
			log.Fatal(er)
		}

		next := step + i
		if next > len(line)-1 {
			next -= len(line)
		}

		partner, er := strconv.Atoi(string(line[next]))
		if er != nil {
			log.Fatal(er)
		}

		//fmt.Printf("%v: %v is %v steps ahead of %v at position %v", i, partner, step, curr, partner)

		if curr == partner {
			result += curr
		}
	}
	return result
}
