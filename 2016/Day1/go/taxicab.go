package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result := distance(line)
		fmt.Printf("Total Distance (blocks): %v", result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func distance(value string) int {
	x := 0
	y := 0
	d := 0 // 0 = N, 1 = E, 2 = S, 3 = W, L = ++, R = --

	entrys := strings.Split(value, ",")
	for _, s := range entrys {
		value := strings.TrimSpace(s)
		dir := value[0]                               // first character
		dis, err := strconv.Atoi(value[1:len(value)]) // remaining string
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: ", s)

		if dir == 'L' {
			if d == 3 {
				d = 0
			} else {
				d++
			}
			fmt.Printf("Turn left | direction = %v || ", d)
		}
		if dir == 'R' {
			if d == 0 {
				d = 3
			} else {
				d--
			}
			fmt.Printf("Turn right | direction = %v || ", d)
		}
		if d == 0 {
			x += dis
			fmt.Printf("Travel north | distance = %v || ", dis)
		}
		if d == 1 {
			y += dis
			fmt.Printf("Travel east | distance = %v || ", dis)
		}
		if d == 2 {
			x -= dis
			fmt.Printf("Travel south | distance = %v || ", dis)
		}
		if d == 3 {
			y -= dis
			fmt.Printf("Travel west | distance = %v || ", dis)
		}
		fmt.Printf("(%v, %v)\n", x, y)
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}
