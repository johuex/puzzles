package g

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	LEN int = iota
	COST
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

	serial := make([]int, n)
	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	lineArr := strings.Split(lineStr, " ")
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		serial[i] = converted
	}

	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	lineArr = strings.Split(lineStr, " ")
	importance := make([]int, n)
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		importance[i] = converted
	}

	// logic
	ans := compensation(serial, importance)

	ansStr := make([]string, len(ans))
	for i := range len(ans) {
		ansStr[i] = strconv.Itoa(ans[i])
	}
	// output
	writer.WriteString(strings.Join(ansStr, " "))
	writer.WriteByte('\n')
}

func compensation(serial, importance []int) []int {
	n := len(serial)

	seasons := make([][]int, n+1)
	for i := 0; i < n; i++ {
		seasons[i+1] = []int{serial[i], importance[i]}
	}
	seasons[0] = []int{0, 0}

	sort.Slice(seasons, func(i, j int) bool {
		if seasons[i][LEN] != seasons[j][LEN] {
			return seasons[i][LEN] < seasons[j][LEN]
		}
		return seasons[i][COST] < seasons[j][COST]
	})

	seasons[0] = []int{seasons[1][0], seasons[1][1]}
	seasons = append(seasons, []int{seasons[len(seasons)-1][0], seasons[len(seasons)-2][1]})

	preCost := make([]int, n+2)
	preCostEpisode := make([]int, n+2)
	suffCost := make([]int, n+2)
	suffCostEpisode := make([]int, n+2)

	for i := 1; i <= n; i++ {
		preCostEpisode[i] = preCostEpisode[i-1] + seasons[i][COST]
		preCost[i] = preCost[i-1] + preCostEpisode[i-1]*(seasons[i][LEN]-seasons[i-1][LEN])
	}

	for i := n; i > 0; i-- {
		suffCostEpisode[i] = suffCostEpisode[i+1] + seasons[i][COST]
		suffCost[i] = suffCost[i+1] + suffCostEpisode[i+1]*(seasons[i+1][LEN]-seasons[i][LEN])
	}

	best_len := 0
	ans := math.MaxInt64

	for i := 1; i <= n; i++ {
		tmp := preCost[i] + suffCost[i]
		if tmp < ans {
			ans = tmp
			best_len = seasons[i][LEN]
		}
	}

	return []int{best_len, ans}
}
