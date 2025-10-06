package f

import (
	"bufio"
	"os"
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

	lines := make([]string, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		lines[i] = strings.TrimSpace(line)
	}
	// logic
	ans := pcq(n, m, lines)
	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

// goal: find max row and min column
// max -- sum of + and .
// min -- sum of - and .
func pcq(n, m int, lines []string) int {
	rowSum := make([]int, n)
	colSum := make([]int, m)
	for i := range n {
		for j := range m {
			if lines[i][j] == '+' {
				rowSum[i] += 1
				colSum[j] += 1
			}
			if lines[i][j] == '-' {
				rowSum[i] -= 1
				colSum[j] -= 1
			}
			if lines[i][j] == '?' {
				rowSum[i] += 1
				colSum[j] -= 1
			}
		}
	}
	ans := -2 * n * m
	for i := range n {
		for j := range m {
			if lines[i][j] == '?' {
				ans = max(ans, rowSum[i]-colSum[j])
			} else {
				ans = max(ans, rowSum[i]-colSum[j]-2)
			}
		}
	}
	return ans
}
