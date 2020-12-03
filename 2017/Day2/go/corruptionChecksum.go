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
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cs []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cs = append(cs, rowCheckSum(line))
	}

	fmt.Printf("The spreadsheet's checksum is %v", sum(cs))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sum(cs []int) int {
	result := 0
	for _, v := range cs {
		result += v
	}
	return result
}

func rowCheckSum(line string) int {
	values := strings.Split(line, "\t")
	smallest, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal(err)
	}
	largest, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		if v < smallest {
			smallest = v
		}
		if v > largest {
			largest = v
		}
	}

	return largest - smallest
}
