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

	treeMap1 := buildMap(mapFile, 1, 1)
	collisions1 := checkSlopeCollisions(treeMap1, 1, 1)
	treeMap2 := buildMap(mapFile, 1, 3)
	collisions2 := checkSlopeCollisions(treeMap2, 1, 3)
	treeMap3 := buildMap(mapFile, 1, 5)
	collisions3 := checkSlopeCollisions(treeMap3, 1, 5)
	treeMap4 := buildMap(mapFile, 1, 7)
	collisions4 := checkSlopeCollisions(treeMap4, 1, 7)
	treeMap5 := buildMap(mapFile, 2, 1)
	collisions5 := checkSlopeCollisions(treeMap5, 2, 1)

	fmt.Printf("Collision with %v trees", collisions1*collisions2*collisions3*collisions4*collisions5)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func buildMap(mapFile []string, rise int, run int) [][]bool {
	scale := (len(mapFile) / rise) * run
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
			//fmt.Printf("At position %v row, %v col = %v\n", i, j, c)
			if c {
				result++
			}
			break
		}
	}

	return result
}
