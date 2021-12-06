package day5

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
	moves := parseInputs()
	//fmt.Printf("Moves: %v\n", moves)

	fmt.Printf("Part1 ans: %v\n", part1(moves))
	fmt.Printf("Part2 ans: %v\n", part2(moves))

}

func part1(moves []move) int {
	mp := make(map[int]map[int]int)
	for _, move := range moves {
		if move.horizontal() {
			moveHorizontal(move, mp)
		}

		if move.vertical() {
			moveVertical(move, mp)
		}
	}

	return CalculateScore(mp)
}

func part2(moves []move) int {
	mp := make(map[int]map[int]int)
	for _, move := range moves {
		if move.horizontal() {
			moveHorizontal(move, mp)
		}

		if move.vertical() {
			moveVertical(move, mp)
		}

		if move.posDiagonal() {
			movePosDiagonal(move, mp)
		}

		if move.negDiagonal() {
			moveNegDiagonal(move, mp)
		}
	}

	return CalculateScore(mp)
}

func moveHorizontal(m move, mp map[int]map[int]int) {
	y := m.src.y
	for x := min(m.src.x, m.dst.x); x <= max(m.src.x, m.dst.x); x++ {
		Insert2Map(mp, x, y)
	}
}

func moveVertical(m move, mp map[int]map[int]int) {
	x := m.src.x
	for y := min(m.src.y, m.dst.y); y <= max(m.src.y, m.dst.y); y++ {
		Insert2Map(mp, x, y)
	}
}

func movePosDiagonal(m move, mp map[int]map[int]int) {
	xMin := min(m.src.x, m.dst.x)
	xMax := max(m.src.x, m.dst.x)
	yMin := min(m.src.y, m.dst.y)
	yMax := max(m.src.y, m.dst.y)
	for x, y := xMin, yMin; x <= xMax && y <= yMax; x, y = x+1, y+1 {
		Insert2Map(mp, x, y)
	}
}

func moveNegDiagonal(m move, mp map[int]map[int]int) {
	xMin := min(m.src.x, m.dst.x)
	xMax := max(m.src.x, m.dst.x)
	yMin := min(m.src.y, m.dst.y)
	yMax := max(m.src.y, m.dst.y)
	for x, y := xMin, yMax; x <= xMax && y >= yMin; x, y = x+1, y-1 {
		Insert2Map(mp, x, y)
	}
}

func Insert2Map(mp map[int]map[int]int, x, y int) {
	if _, found := mp[x]; !found {
		mp[x] = map[int]int {y : 0}
	}
	mp[x][y] += 1
}

func CalculateScore(mp map[int]map[int]int) int {
	ans := 0
	for _, val := range mp {
		for _, v := range val {
			if v >= 2 {
				ans += 1
			}
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type point struct {
	x int
	y int
}

type move struct {
	src point
	dst point
}

func (m *move) horizontal() bool {
	return m.src.y == m.dst.y
}

func (m *move) vertical() bool {
	return m.src.x == m.dst.x
}

func (m *move) posDiagonal() bool {
	xDiff := m.src.x - m.dst.x
	yDiff := m.src.y - m.dst.y
	return xDiff == yDiff
}

func (m *move) negDiagonal() bool {
	xDiff := m.src.x - m.dst.x
	yDiff := m.src.y - m.dst.y
	return xDiff == -1 * yDiff
}

func parseInputs() []move {
	file, err := os.Open("day5/day5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var moves []move
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			continue //skip empty lines
		}

		parts := strings.Fields(string(line))
		if len(parts) != 3 {
			log.Fatalf("Invalid input format for : %v", parts)
		}
		var m move

		src := strings.Split(parts[0], ",")
		dst := strings.Split(parts[2], ",")
		m.src.x, _ = strconv.Atoi(src[0])
		m.src.y, _ = strconv.Atoi(src[1])

		m.dst.x, _ = strconv.Atoi(dst[0])
		m.dst.y, _ = strconv.Atoi(dst[1])

		moves = append(moves, m)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return moves
}
