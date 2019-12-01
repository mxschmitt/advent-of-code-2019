package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}
	defer inFile.Close()
	sum := 0
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("could not parse mass to int: %v", err)
		}
		sum += int(mass/3) - 2
	}
	fmt.Printf("sum of the needed fuel: %d\n", sum)
}
