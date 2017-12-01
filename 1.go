package main 

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strconv"
)

func partOne() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	counter := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line) - 1; i++ {
			if line[i] == line[i+1] {
				digit, _ := strconv.Atoi(string(line[i]))
				counter += digit
			}
		}
		if line[len(line)-1] == line[0] {
			digit, _ := strconv.Atoi(string(line[0]))
			counter += digit
		}
	}
	fmt.Println(counter)
}

func partTwo() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	counter := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]
		for i := 0; i < len(firstHalf); i++ {
			if firstHalf[i] == secondHalf[i] {
				digit, _ := strconv.Atoi(string(firstHalf[i]))
				counter += digit
			}
		}
	}
	fmt.Println(counter * 2)
}

func main() {
	partOne()
	partTwo()
}