package day4

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	moves, boards, locMap := getInputs()
	fmt.Printf("Moves: %v\n, \nboards: %v\n", len(locMap), len(boards))

	bingoMap := make(map[int]bool)
	for i := 0; i < len(boards); i++ {
		bingoMap[i] = false
	}

	lastScore := 0
	foundLastBingo := false
	for _, move := range moves {
		positions, found := locMap[move]
		if !found {
			continue
		}

		for _, position := range positions {
			if _, found := bingoMap[position.num]; !found {
				continue
			}
			board := &boards[position.num]
			board.mark[position.i][position.j] = true
			//fmt.Printf("\n board: %v\n", board.mark)
			if board.Bingo() {
				score := board.Score(move)
				if score != 0 {
					fmt.Printf("Bingo Reached for board num: %v , Score: %v\n", position.num, score)
				}
				delete(bingoMap, position.num)
				if len(bingoMap) == 0 {
					lastScore = score
					foundLastBingo = true
					break
				}
				//return
			}
		}

		if foundLastBingo {
			break
		}
	}

	fmt.Printf("\n Last Bingo Score: %v\n", lastScore)
}

type pos struct {
	num int
	i int
	j int
}

type Board struct {
	board [5][5]int
	mark  [5][5]bool
}

func (b *Board) Bingo() bool {
	for i := 0; i < 5; i++ {
		bingoR := true
		bingoC := true
		for j := 0; j < 5; j++ {
			if !b.mark[i][j] {
				bingoR = false
			}
			if !b.mark[j][i] {
				bingoC = false
			}
		}
		if bingoR || bingoC {
			return true
		}
	}
	return false
}

func (b *Board) Score(num int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.mark[i][j] {
				sum += b.board[i][j]
			}
		}
	}
	return sum * num
}

func getInputs() ([]int, []Board, map[int][]pos) {
	file, err := os.Open("day4/day4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//Read 1st line
	scanner.Scan()
	line := bytes.TrimSpace(scanner.Bytes())
	codes := strings.Split(string(line), ",")

	moves := make([]int, len(codes))
	for i := range codes {
		moves[i], err = strconv.Atoi(codes[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	mp := make(map[int][]pos)
	var boards []Board
	cnt := 0
	var b Board
	for scanner.Scan() {
		if cnt == 5 {
			cnt = 0
			boards = append(boards, b)
		}
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			continue
		}

		codes := strings.Fields(string(line))
		for j := 0; j < 5; j++ {
			//fmt.Println(codes[j])
			num, err := strconv.Atoi(codes[j])
			if err != nil {
				log.Fatal(err)
			}

			b.board[cnt][j] = num
			mp[num] = append(mp[num], pos{num: len(boards), i: cnt, j: j })
		}
		//fmt.Println(scanner.Text())
		cnt += 1
	}

	boards = append(boards, b)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return moves, boards, mp
}