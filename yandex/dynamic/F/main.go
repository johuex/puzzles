package f

import (
	"bufio"
	"os"
	"slices"
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
	n, _ := strconv.Atoi(strings.TrimSpace(line))
	lines := make([]string, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		lines[i] = strings.TrimSpace(line)
	}

	// logic
	ans := coinMax(n, lines)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func coinMax(n int, lines []string) int {
	// 1. создать карту монет и окружить сверху, слева, справа -100000
	coinMap := make([][]int, n+1)
	for i := range n + 1 {
		if i == 0 {
			coinMap[i] = []int{-100000, -100000, -100000}
		}
		coinMap[i] = make([]int, 5)
		coinMap[i][0] = -100000
		coinMap[i][4] = -100000
	}

	// сплошная стена на старте, прохода нет
	if lines[0] == "WWW" {
		return 0
	}
	// идем по карте, берем максимум от сверху, или диагоналей сверху + кол-во монет в текущей клетке
	for i := 1; i < n+1; i++ {
		for j := 1; j < 4; j++ {
			if lines[i-1][j-1] == 'W' {
				coinMap[i][j] = -100000
				continue
			}
			maxVal := max(coinMap[i-1][j-1], coinMap[i-1][j], coinMap[i-1][j+1])
			coinMap[i][j] = maxVal
			if maxVal < 0 {
				// клетка, в которую нельзя придти
				// всегда будет 0, даже есть там есть монета
				//maxVal = 0
				continue
			}

			if lines[i-1][j-1] == 'C' {
				coinMap[i][j] += 1
			}
		}
	}

	res := -100000
	ansIdx := n
	for i := 0; i < n+1; i++ {
		res = slices.Max(coinMap[i])
		if res == -100000 {
			ansIdx = i - 1
			break
		}
	}
	ans := slices.Max(coinMap[ansIdx])
	if ans < 0 {
		ans = 0
	}
	return ans
}
