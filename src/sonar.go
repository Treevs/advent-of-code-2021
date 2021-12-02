package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var exDepths = fileToArray("../inputs/sonar1.txt")
	var ansDepths = fileToArray("../inputs/sonar2.txt")

	exOneDepth := checkOneDepth(exDepths[:])
	fmt.Printf("exOneDepth: %d\n", exOneDepth)
	ansOneDepth := checkOneDepth(ansDepths[:])
	fmt.Printf("ansOneDepth: %d\n", ansOneDepth)
	exThreeDepth := checkThreeDepth(exDepths[:])
	fmt.Printf("exThreeDepth: %d\n", exThreeDepth)
	ansThreeDepth := checkThreeDepth(ansDepths[:])
	fmt.Printf("ansThreeDepth: %d\n", ansThreeDepth)
}

func checkOneDepth(depths []int) int {
	count := 0
	prev := depths[0]
	for _, depth := range depths {
		if depth > prev {
			count++
		}
		prev = depth
	}
	return count
}
func checkThreeDepth(depths []int) int {
	count := 0
	one := depths[0]
	two := depths[1]
	three := depths[2]
	for _, depth := range depths {
		if depth+three+two > three+two+one {
			count++
		}
		one = two
		two = three
		three = depth
	}
	return count
}

func fileToArray(filePath string) []int {
	file, err := os.Open(filePath)
	array := make([]int, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		array = append(array, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return array
}
