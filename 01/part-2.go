package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateFurtherFuel(mass int) int {
	return int(mass/3) - 2
}

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
		furtherFuel := calculateFurtherFuel(mass)
		for furtherFuel > 0 {
			mass = furtherFuel
			sum += furtherFuel
			furtherFuel = calculateFurtherFuel(mass)
		}

	}
	fmt.Printf("sum for all the fuel requirements: %d\n", sum)
}
