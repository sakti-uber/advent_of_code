package day2

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Run() {
	fmt.Println(executeCommandsPart1())
	fmt.Println(executeCommandsPart2())
}

func executeCommandsPart1() int {
	commands, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}

	x, y := 0, 0
	for _, command := range commands {
		switch command.direction {
		case "up":
			y -= command.value
		case "down":
			y += command.value
		case "forward":
			x += command.value
		}
	}
	return x * y
}

func executeCommandsPart2() int {
	commands, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}

	x, y, aim := 0, 0, 0
	for _, command := range commands {
		switch command.direction {
		case "up":
			aim -= command.value
		case "down":
			aim += command.value
		case "forward":
			x += command.value
			y += aim * command.value
		}
	}
	return x * y
}

type Command struct {
	direction string
	value     int
}

func getInputs() ([]Command, error) {
	// open file
	file, err := os.Open("day2/day2_input.txt")
	if err != nil {
		return nil, err
	}
	// close the file at the end of the program
	defer file.Close()

	var commands []Command

	lineNum := 0
	for {
		// Read each integer per line
		var command Command
		if _, err := fmt.Fscanf(file, "%s %d\n", &command.direction, &command.value); err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			return nil, fmt.Errorf("error reading input test file at line number: %v, Error:%v", lineNum, err)
		}

		if !(command.direction == "up" || command.direction == "down" || command.direction == "forward") {
			return nil, fmt.Errorf("error reading input test file at line number: %v, invalid direction:%v", lineNum, command.direction)
		}

		commands = append(commands, command)
		lineNum += 1
	}
	return commands, nil
}
