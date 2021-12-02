package day1

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Run() {
	fmt.Println(countSingleIncrease())
	fmt.Println(countSlidingIncrease())
}

func countSingleIncrease() int {
	data, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}

	cnt := 0
	for i := range data {
		if i == 0 {
			continue
		}
		if data[i] > data[i-1] {
			cnt += 1
		}
	}

	return cnt
}

func countSlidingIncrease() int {
	data, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}

	cnt, sum1, sum2 := 0, 0, 0
	for i := 0; i < 3 && i < len(data); i++ {
		sum1 += data[i]
		if i > 0 {
			sum2 += data[i]
		}
	}
	for i := 3; i < len(data); i++ {
		sum2 += data[i]
		if sum2 > sum1 {
			cnt += 1
		}
		sum1 = sum2
		sum2 -= data[i-2]
	}

	return cnt
}

func getInputs() ([]int, error) {
	// open file
	file, err := os.Open("day1/day1_input.txt")
	if err != nil {
		return nil, err
	}
	// close the file at the end of the program
	defer file.Close()

	var nums []int
	line, lineNum := 0, 0
	for {
		// Read each integer per line
		if _, err := fmt.Fscanf(file, "%d\n", &line); err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			return nil, fmt.Errorf("error reading input test file at line number: %v, Error:%v", lineNum, err)
		}
		nums = append(nums, line)
		lineNum += 1
	}
	return nums, nil
}
