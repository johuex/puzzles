package e

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
	lineArr := strings.Split(lineStr, " ")
	n, _ := strconv.Atoi(lineArr[0])
	m, _ := strconv.Atoi(lineArr[1])
	k, _ := strconv.Atoi(lineArr[2])

	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	lineArr = strings.Split(lineStr, " ")
	road := make([]int, n)
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		road[i] = converted
	}

	routes := make([][]int, m)
	for i := range m {
		line, _ = reader.ReadString('\n')
		lineStr = strings.TrimSpace(line)
		lineArr = strings.Split(lineStr, " ")
		first, _ := strconv.Atoi(lineArr[0])
		second, _ := strconv.Atoi(lineArr[1])
		routes[i] = []int{first, second}
	}

	// logic
	ans := repair(road, routes, k)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func repair(road []int, routes [][]int, holes int) int {
	negativeSum := make([]int, len(road)+1)
	routeEntries := make([][]int, len(road))
	for i := range len(road) {
		routeEntries[i] = []int{i, 0}
	}

	// refactored, set route entries on road w/o O(n^2)
	diff := make([]int, len(road)+1)
	for _, r := range routes {
		l := r[0] - 1
		rgt := r[1] - 1
		diff[l]++
		diff[rgt+1]--
	}
	cur := 0
	for i := 0; i < len(road); i++ {
		cur += diff[i]
		routeEntries[i][1] = cur
	}

	for idx, elem := range road {
		negativeSum[idx+1] = negativeSum[idx] + elem*routeEntries[idx][1]
	}

	sort.Slice(routeEntries, func(i, j int) bool {
		return routeEntries[i][1] > routeEntries[j][1]
	})

	negativeSumMax := negativeSum[len(negativeSum)-1]
	idx := 0

	for holes > 0 && idx < len(routeEntries) {
		w2Dec := min(road[routeEntries[idx][0]], holes)
		negativeSumMax -= w2Dec * routeEntries[idx][1]
		holes -= w2Dec
		idx += 1
	}

	if negativeSumMax < 0 {
		negativeSumMax = 0
	}
	return negativeSumMax

}
