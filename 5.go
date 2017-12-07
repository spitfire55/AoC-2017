package main 

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strconv"
)

func partOne() {
	f, err := os.Open("5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	jmpTable := make([]int, 1074)

	scanner := bufio.NewScanner(f)
	i := 0 // scanner line counter
	for scanner.Scan() {
		line := scanner.Text()
		jmpVal, _ := strconv.Atoi(line)
		jmpTable[i] = jmpVal
		i++
	}
	pos := 0 // jump table position
	stepCount := 0 // final value of steps taken
	for pos > -1 && pos < len(jmpTable) {
		jmpVal := jmpTable[pos]
		jmpTable[pos]++
		pos += jmpVal
		stepCount++
	}
	fmt.Println(stepCount)
}

func partTwo() {
	f, err := os.Open("5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	jmpTable := make([]int, 1074)

	scanner := bufio.NewScanner(f)
	i := 0 // scanner line counter
	for scanner.Scan() {
		line := scanner.Text()
		jmpVal, _ := strconv.Atoi(line)
		jmpTable[i] = jmpVal
		i++
	}
	pos := 0 // jump table position
	stepCount := 0 // final value of steps taken
	for pos > -1 && pos < len(jmpTable) {
		jmpVal := jmpTable[pos]
		if jmpTable[pos] > 2 {
			jmpTable[pos]--
		} else {
			jmpTable[pos]++
		}
		pos += jmpVal
		stepCount++
	}
	fmt.Println(stepCount)
}

func main() {
	partOne()
	partTwo()
}