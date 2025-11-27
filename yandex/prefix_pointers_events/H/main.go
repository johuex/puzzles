package h

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
	lineStr := strings.TrimSpace(line)
	n, _ := strconv.Atoi(lineStr)

	prize := make([]int, n)
	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	lineArr := strings.Split(lineStr, " ")
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		prize[i] = converted
	}

	// logic
	ans := boss(prize)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func boss(prize []int) int {
	n := len(prize)
	segments := make([]int, n) // balance of start and end of segments
	for i := range n {
		segments[i] += 1    // start of segment
		if prize[i]+i < n { // end of segment in our field
			idx := prize[i] + i - 1
			if idx < 0 {
				idx *= -1
			}
			segments[idx] -= 1
		}
	}
	var ans int     // total rewward for all employee
	var balance int // cnt of segments

	for i := range n {
		ans += balance * prize[i]
		balance += segments[i]
	}

	return ans
}
