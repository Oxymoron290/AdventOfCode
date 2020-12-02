package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += measure(scanner.Text())
	}
	fmt.Printf("Total square footage: %v sqft", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func measure(dimensions string) int {
	side := strings.Split(dimensions, "x")
	length, err := strconv.Atoi(side[0])
	if err != nil {
		log.Fatal(err)
	}
	width, err := strconv.Atoi(side[1])
	if err != nil {
		log.Fatal(err)
	}
	height, err := strconv.Atoi(side[2])
	if err != nil {
		log.Fatal(err)
	}

	face1 := length * width
	face2 := width * height
	face3 := height * length

	smallest := 0
	if face1 <= face2 && face1 <= face3 {
		smallest = face1
	} else if face2 <= face1 && face2 <= face3 {
		smallest = face2
	} else {
		smallest = face3
	}

	return ((2 * face1) + (2 * face2) + (2 * face3)) + smallest
}
