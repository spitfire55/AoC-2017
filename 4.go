package main 

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	"reflect"
)

func charCountToMap(word string) map[rune]int {
	charMap := make(map[rune]int)
	for _, char := range word {
		charMap[char]++
	}
	return charMap
}

func partOne() {
	passCount := 0

	f, err := os.Open("4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		for i := 0; i < len(fields); i++ {
			for j := i+1; j < len(fields); j++ {
				if fields[i] == fields[j] {
					goto nextOne
				}
			}
		}
		passCount++
		nextOne:
	}
	fmt.Println(passCount)
}

func partTwo() {
	passCount := 0

	f, err := os.Open("4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		// count characters in first word and save count to map
		for i, start := range fields { 
			iMap := charCountToMap(start)
			// count characters in subsequent words and save to map
			for j := i+1; j < len(fields); j++ {
				jMap := charCountToMap(fields[j])
				if reflect.DeepEqual(iMap, jMap) {
					goto nextOne
				}
			}
		}
		passCount++
		nextOne:
	}
	fmt.Println(passCount)
}

func main() {
	partOne()
	partTwo()
}