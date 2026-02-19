package minwindow

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, " ")
	n, _ := strconv.Atoi(lineArr[0])
	k, _ := strconv.Atoi(lineArr[1])

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr = strings.Split(line, " ")
	input := make([]int, n)
	for i := range n {
		num, _ := strconv.Atoi(lineArr[i])
		input[i] = num
	}

	res := minWindow(n, k, input)

	resStrArr := make([]string, len(res))
	for i := range len(res) {
		resStrArr[i] = strconv.Itoa(res[i])
	}
	writer.WriteString(strings.Join(resStrArr, "\n"))
	writer.WriteByte('\n')
}

func minWindow(n, k int, arr []int) []int {
	if n == 0 || k == 0 || k > n {
		return []int{}
	}

	deq := make([]int, 0)
	res := make([]int, n-k+1)
	resIdx := 0

	for i := 0; i < n; i++ {
		// force pop left, out of window
		if len(deq) > 0 && deq[0] < i-k+1 {
			deq = deq[1:]
		}

		// before adding -- pop right idx that greater actual by value (keep monotonously increasing function)
		for len(deq) > 0 && arr[deq[len(deq)-1]] > arr[i] {
			deq = deq[:len(deq)-1]
		}

		// push right
		deq = append(deq, i)

		// store min as deq[0], fill only if window is full
		if i >= k-1 {
			res[resIdx] = arr[deq[0]]
			resIdx++
		}
	}

	return res
}
