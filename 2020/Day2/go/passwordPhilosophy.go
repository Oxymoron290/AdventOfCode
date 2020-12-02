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
		//if validate(parse(line)) {
		if validateNew(parse(line)) {
			count++
		}
		total++
	}

	fmt.Printf("There are %v out of %v valid passwords\n", count, total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parse(line string) (int, int, rune, string) {
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
	target := []rune(entry[1])[0]

	return min, max, target, value
}

func validate(min int, max int, target rune, value string) bool {
	count := strings.Count(value, string(target))
	result := (count <= max && count >= min)
	//fmt.Printf("min(%v) max(%v) target(%v) %v %v %v\n", min, max, target, value, count, result)
	return result
}

func validateNew(pos1 int, pos2 int, target rune, value string) bool {
	atPos1 := []rune(value)[pos1-1]
	x := atPos1 == target
	atPos2 := []rune(value)[pos2-1]
	y := atPos2 == target
	//fmt.Printf("Value at position %v is %v - Match = %v\nValue at position %v is %v - Match = %v\nOverall Validation: %v\n", pos1, string(atPos1), x, pos2, string(atPos2), y, x != y)
	return x != y
}
