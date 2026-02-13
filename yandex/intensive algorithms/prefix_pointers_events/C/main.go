package c

import (
	"bufio"
	"fmt"
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
	lineArr := strings.Split(lineStr, " ")
	n, _ := strconv.Atoi(lineArr[0]) // initial cnt of candidates
	profThr, _ := strconv.Atoi(lineArr[1])

	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	lineArr = strings.Split(lineStr, " ")
	candidates := make([]int, n)
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		candidates[i] = converted
	}

	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	m, _ := strconv.Atoi(lineStr)
	events := make([][]int, m)
	for i := range m {
		line, _ := reader.ReadString('\n')
		lineStr := strings.TrimSpace(line)
		lineArr := strings.Split(lineStr, " ")
		switch len(lineArr) {
		case 1:
			first, _ := strconv.Atoi(lineArr[0])
			events[i] = []int{first}

		case 2:
			first, _ := strconv.Atoi(lineArr[0])
			second, _ := strconv.Atoi(lineArr[1])
			events[i] = []int{first, second}

		default:
			fmt.Errorf("unexpected number of fields in line: %v", lineArr)
		}
	}

	// logic
	ans := candidatesCheck(profThr, candidates, events)

	// output
	ansStr := make([]string, len(ans))
	for i := range len(ans) {
		ansStr[i] = strconv.Itoa(ans[i])
	}
	writer.WriteString(strings.Join(ansStr, "\n"))
	writer.WriteByte('\n')
}

func candidatesCheck(profThr int, candidates []int, events [][]int) []int {
	res := make([]int, 0)

	good := make([]int, len(candidates)+1)
	// initail filling prefixSum of good candidates
	for i, v := range candidates {
		good[i+1] = good[i]
		if v >= profThr {
			good[i+1] += 1
		}
	}

	start := 0

	for _, event := range events {
		if event[0] == 1 {
			// add candidate
			v := event[1]
			candidates = append(candidates, v)
			if v >= profThr {
				good = append(good, good[len(good)-1]+1)
			} else {
				good = append(good, good[len(good)-1])
			}
		}

		if event[0] == 2 {
			// fast remove candidate from start
			start++
		}

		if event[0] == 3 {
			// need to Answer Mary about candidates
			k := event[1]
			end := start + k
			cnt := good[end] - good[start]
			res = append(res, cnt)
		}
	}

	return res
}
