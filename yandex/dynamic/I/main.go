package i

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// general
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, " ")
	n, _ := strconv.Atoi(lineArr[0])
	m, _ := strconv.Atoi(lineArr[1])

	// read 2D array
	numbers := make([][]int, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lineArr = strings.Split(line, " ")
		tmp := make([]int, m)
		for j := range m {
			tmpN, _ := strconv.Atoi(lineArr[j])
			tmp[j] = tmpN
		}
		numbers[i] = tmp
	}

	// logic
	ans := chain(numbers, n, m)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

type Point struct {
	Val   int
	Coord [][]int
}

func NewPoint() *Point {
	return &Point{
		Coord: make([][]int, 0),
	}
}

// table for fast movements near the point
var nearTable = [][]int{
	{0, -1}, // left
	{0, 1},  // right
	{-1, 0}, // up
	{1, 0},  // down
}

func padWithZeros(table [][]int) [][]int {
	if len(table) == 0 || len(table[0]) == 0 {
		return [][]int{}
	}

	n := len(table)
	m := len(table[0])

	// создаём новую таблицу размером (n+2) x (m+2)
	newTable := make([][]int, n+2)
	for i := range newTable {
		newTable[i] = make([]int, m+2) // автоматически заполнено нулями
	}

	// копируем старую таблицу в центр
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			newTable[i+1][j+1] = table[i][j]
		}
	}

	return newTable
}

func getPoints(numbers [][]int, n, m int) []*Point {
	// create map with value -> [coordinates]
	mc := make(map[int]*Point)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			val := numbers[i][j]
			if val == 0 {
				continue
			}
			p, ok := mc[val]
			if !ok {
				p = NewPoint()
				p.Val = val
			}
			p.Coord = append(p.Coord, []int{i, j})
			mc[val] = p
		}
	}

	// transform to slice
	res := make([]*Point, len(mc))
	var idx int
	for _, val := range mc {
		res[idx] = val
		idx++
	}

	// sort slice in desc order by val
	sort.Slice(res, func(i, j int) bool {
		return res[i].Val > res[j].Val
	})

	return res

}

func chain(numbers [][]int, n, m int) int {
	// dp with borders from all sides
	dp := make([][]int, n+2)
	for i := range n + 2 {
		tmp := make([]int, m+2)
		dp[i] = tmp
	}
	numbers = padWithZeros(numbers)
	points := getPoints(numbers, n+2, m+2)

	// range over point by desc
	for _, point := range points {
		// found max near +-1 points if neighbor == point+1
		pVal := point.Val
		pCoords := point.Coord
		for _, coord := range pCoords {
			// search max val in neighborgs
			maxWays := 0 // max value
			for _, shift := range nearTable {
				newI := coord[0] + shift[0]
				newJ := coord[1] + shift[1]
				neig := numbers[newI][newJ]
				if pVal+1 == neig && dp[newI][newJ] > maxWays {
					maxWays = dp[newI][newJ]
				}
			}
			dp[coord[0]][coord[1]] = maxWays + 1
		}
	}

	var maxRes int
	for _, line := range dp {
		for _, val := range line {
			maxRes = max(maxRes, val)
		}
	}

	return maxRes
}
