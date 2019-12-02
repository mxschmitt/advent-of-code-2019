package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func scanComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

const OPCODE_ADDITION = 1
const OPCODE_MULTIPLICATION = 2
const OPCODE_EXIT = 99

func main() {
	inFile, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(scanComma)
	originalInput := []int{}
	for scanner.Scan() {
		inputValue, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("could not parse mass to int: %v", err)
		}
		originalInput = append(originalInput, inputValue)
	}
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			input := make([]int, len(originalInput))
			copy(input, originalInput)
			input[1] = noun
			input[2] = verb
			current := 0
			for input[current] != OPCODE_EXIT {
				switch input[current] {
				case OPCODE_ADDITION:
					input[input[current+3]] = input[input[current+1]] + input[input[current+2]]
					current += 4
				case OPCODE_MULTIPLICATION:
					input[input[current+3]] = input[input[current+1]] * input[input[current+2]]
					current += 4
				default:
				}
			}
			if input[0] == 19690720 {
				fmt.Printf("noun: %d verb: %d\n", noun, verb)
			}
		}
	}
}
