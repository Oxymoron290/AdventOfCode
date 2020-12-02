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

	var value string
	test := false
	var expect int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if value == "" {
			value = scanner.Text()
		} else {
			e, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			expect = e
			test = true
		}
	}

	//floors := countFloors(value)
	//fmt.Printf("Santa traveled to floor #%v\n", floors)
	floors := findBasement(value)
	fmt.Printf("Santa entered the basement at position #%v", floors)
	if test {
		if floors == expect {
			fmt.Println("Test passed!")
		} else {
			fmt.Println("Test Failed!")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func countFloors(value string) int {
	var result = 0
	for _, c := range value {
		if c == '(' {
			result++
		}
		if c == ')' {
			result--
		}
	}
	return result
}

func findBasement(value string) int {
	var result = 0
	for i, c := range value {
		if c == '(' {
			result++
		}
		if c == ')' {
			result--
		}
		if result == -1 {
			return i + 1
		}
	}
	log.Fatal("Santa never entered the basement!")
	return 0
}
