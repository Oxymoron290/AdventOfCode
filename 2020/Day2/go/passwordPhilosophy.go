package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type policy struct {
	min   int
	max   int
	value rune
}

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

	total := 0
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if validate(parse(line)) {
			count++
		}
		total++
	}

	fmt.Printf("There are %v out of %v valid passwords\n", count, total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parse(line string) (int, int, string, string) {
	entry := strings.Split(line, ":")
	value := strings.TrimSpace(entry[1])
	entry = strings.Split(entry[0], "-")
	min, err := strconv.Atoi(entry[0])
	if err != nil {
		log.Fatal(err)
	}
	entry = strings.Split(entry[1], " ")
	max, err := strconv.Atoi(entry[0])
	if err != nil {
		log.Fatal(err)
	}
	target := entry[1]

	return min, max, target, value
}

func validate(min int, max int, target string, value string) bool {
	count := strings.Count(value, target)
	result := (count <= max && count >= min)
	//fmt.Printf("min(%v) max(%v) target(%v) %v %v %v\n", min, max, target, value, count, result)
	return result
}
