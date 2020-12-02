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

	var entries []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry, er := strconv.Atoi(scanner.Text())
		if er != nil {
			log.Fatal(er)
		}
		entries = append(entries, entry)
	}

	//v1, v2 := findSummers(2020, entries)
	//fmt.Printf("The answer is %v * %v = %v", v1, v2, v1*v2)
	v1, v2, v3 := findTriSummers(2020, entries)
	fmt.Printf("The answer is %v * %v * %v = %v", v1, v2, v3, v1*v2*v3)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findSummers(value int, entries []int) (int, int) {
	for index, element := range entries {
		for j := index + 1; j < len(entries); j++ {
			sum := element + entries[j]
			fmt.Printf("%v + %v = %v\n", element, entries[j], sum)
			if sum == value {
				return element, entries[j]
			}
		}
	}
	log.Fatal("no answer found")
	return 0, 0
}

func findTriSummers(value int, entries []int) (int, int, int) {
	for index, element := range entries {
		for j := index + 1; j < len(entries)-1; j++ {
			for k := j + 1; k < len(entries); k++ {
				sum := element + entries[j] + entries[k]
				fmt.Printf("%v + %v + %v = %v\n", element, entries[j], entries[k], sum)
				if sum == value {
					return element, entries[j], entries[k]
				}
			}
		}
	}
	log.Fatal("no answer found")
	return 0, 0, 0
}
