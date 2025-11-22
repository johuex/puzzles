package b

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
	lineStr := strings.TrimSpace(line)
	n, _ := strconv.Atoi(lineStr)

	taxTable := make([][]int, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		lineStr := strings.TrimSpace(line)
		lineArr := strings.Split(lineStr, " ")
		first, _ := strconv.Atoi(lineArr[0])
		second, _ := strconv.Atoi(lineArr[1])
		taxTable[i] = []int{first, second}
	}

	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	m, _ := strconv.Atoi(lineStr)

	cars := make([]int, m)
	for i := range m {
		line, _ = reader.ReadString('\n')
		lineStr = strings.TrimSpace(line)
		converted, _ := strconv.Atoi(lineStr)
		cars[i] = converted
	}

	// logic
	ans := taxCalculator(taxTable, cars)

	// output
	ansStr := make([]string, m)
	for i := range m {
		ansStr[i] = strconv.Itoa(ans[i])
	}
	writer.WriteString(strings.Join(ansStr, "\n"))
	writer.WriteByte('\n')
}

func taxCalculator(taxTable [][]int, cars []int) []int {
	res := make([]int, len(cars))
	resWIdx := make([][]int, len(cars))

	carsWIdx := make([][]int, len(cars))

	for idx, car := range cars {
		carsWIdx[idx] = []int{car, idx}
	}

	sort.Slice(carsWIdx, func(i, j int) bool {
		return carsWIdx[i][0] < carsWIdx[j][0]
	})
	taxIdx := 0

	for carIdx := 0; carIdx < len(cars); carIdx++ {
		for taxIdx+1 < len(taxTable) && carsWIdx[carIdx][0] > taxTable[taxIdx+1][0] { // TODO: налоговая ставка может быть длиной 1
			taxIdx += 1
		}
		resWIdx[carIdx] = []int{carsWIdx[carIdx][1], carsWIdx[carIdx][0] * taxTable[taxIdx][1]}
	}

	// reorder answer
	sort.Slice(resWIdx, func(i, j int) bool {
		return resWIdx[i][0] < resWIdx[j][0]
	})

	for i := range len(cars) {
		res[i] = resWIdx[i][1]
	}

	return res
}
