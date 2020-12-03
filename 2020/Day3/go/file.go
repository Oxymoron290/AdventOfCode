package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var mapFile []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mapFile = append(mapFile, line)
	}
	treeMap := buildMap(mapFile)
	collisions := checkSlopeCollisions(treeMap, 1, 3)

	fmt.Printf("Collision with %v trees", collisions)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func buildMap(mapFile []string) [][]bool {
	scale := len(mapFile) * 3
	var treeMap [][]bool
	for _, l := range mapFile {
		var line []bool
		repeat := math.Ceil(float64(scale) / float64(len(l)))
		for i := repeat; i > 0; i-- {
			for _, s := range l {
				line = append(line, s == '#')
			}
		}
		treeMap = append(treeMap, line)
	}

	return treeMap
}

func checkSlopeCollisions(treeMap [][]bool, rise int, run int) int {
	result := 0
	nextLine := rise
	nextCol := run
	for i, l := range treeMap {
		if nextLine > i {
			continue
		}
		nextLine = i + rise

		for j, c := range l {
			if nextCol > j {
				continue
			}
			nextCol = j + run
			fmt.Printf("At position %v row, %v col = %v\n", i, j, c)
			if c {
				result++
			}
			break
		}
	}

	return result
}
