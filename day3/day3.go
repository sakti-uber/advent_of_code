package day3

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Run() {
	fmt.Println(getPower())
	fmt.Println(getHealth())
}

func getHealth() int {
	codes, err := getInputs()
	if err != nil {
		log.Fatalf("error reading inputs, Error: %v", err)
	}

	if len(codes) == 0 {
		return 0
	}

	m := len(codes[0])

	oxyList := codes
	cnt, pos := len(codes), 0
	for cnt > 1 && pos < m {
		zero, one := getFreqMap(oxyList, pos)
		if len(one) >= len(zero) {
			oxyList = one
			cnt = len(one)
		} else {
			oxyList = zero
			cnt = len(zero)
		}
		pos += 1
	}

	oxyNum := binaryStringToInt(oxyList[0])
	fmt.Printf("oxyList: %v, val:%v\n", oxyList, oxyNum)

	co2List := codes
	cnt, pos = len(codes), 0
	for cnt > 1 && pos < m {
		zero, one := getFreqMap(co2List, pos)
		if len(one) < len(zero) {
			co2List = one
			cnt = len(one)
		} else {
			co2List = zero
			cnt = len(zero)
		}
		pos += 1
	}

	co2Num := binaryStringToInt(co2List[0])
	fmt.Printf("\nco2List: %v, val: %v\n", co2List, co2Num)

	return oxyNum * co2Num
}

func binaryStringToInt(binary string) int {
	res := 0
	for i := 0; i < len(binary); i++ {
		res *= 2
		if binary[i] == '1' {
			res += 1
		}
	}
	return res
}

func getFreqMap(list []string, j int) (zero, one []string) {
	for i := range list {
		if list[i][j] == '0' {
			zero = append(zero, list[i])
		} else {
			one = append(one, list[i])
		}
	}
	return
}

func getPower() int {
	codes, err := getInputs()
	if err != nil {
		log.Fatalf("error reading inputs, Error: %v", err)
	}

	if len(codes) == 0 {
		return 0
	}

	n, m := len(codes), len(codes[0])
	cnt := make([]int, m)

	for _, code := range codes {
		code := code
		for i := range code {
			if code[i] == '1' {
				cnt[i] += 1
			}
		}
	}

	gamma, epsilon := 0, 0

	for i := 0; i < m; i++ {
		gamma *= 2
		epsilon *= 2
		if cnt[i] >= n/2 {
			gamma += 1
		} else {
			epsilon += 1
		}
	}
	return gamma * epsilon
}

func getInputs() ([]string, error) {
	// open file
	file, err := os.Open("day3/day3_input.txt")
	if err != nil {
		return nil, err
	}
	// close the file at the end of the program
	defer file.Close()

	var codes []string

	lineNum := 0
	for {
		// Read each integer per line
		var code string
		if _, err := fmt.Fscanf(file, "%s\n", &code); err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			return nil, fmt.Errorf("error reading input test file at line number: %v, Error:%v", lineNum, err)
		}

		codes = append(codes, code)
		lineNum += 1
	}
	return codes, nil
}