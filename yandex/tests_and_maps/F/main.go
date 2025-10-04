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
	//k, _ := strconv.Atoi(lineArr[1])

	lines := make([]string, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		lines[i] = strings.TrimSpace(line)
	}
	// logic
	ans := pcq(lines)
	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

// goal: find max row and min column
// max -- sum of + and .
// min -- sum of - and .
func pcq(lines []string) int {
	lineSet := make(map[int]map[rune]int)
	vertSet := make(map[int]map[rune]int)

	// transform input
	// TODO: кажется тут баг
	for i, row := range lines {
		lineSet[i] = make(map[rune]int)
		for j, sym := range row {
			if _, ok := vertSet[j]; !ok {
				vertSet[j] = make(map[rune]int)
			}
			vertSet[j][sym] += 1
			lineSet[i][sym] += 1
		}
	}
	// find max line
	maxSum := -(len(vertSet))
	for _, runeMap := range lineSet {
		tempSum := 0
		for runeSym, runeVal := range runeMap {
			switch runeSym {
			case '-':
				tempSum -= runeVal
			case '+', '?':
				tempSum += runeVal
			}
		}
		if maxSum < tempSum {
			maxSum = tempSum
		}
	}

	// found min column
	minSum := len(lineSet)
	for _, runeMap := range vertSet {
		tempSum := 0
		for runeSym, runeVal := range runeMap {
			switch runeSym {
			case '-', '?':
				tempSum -= runeVal
			case '+':
				tempSum += runeVal
			}
		}
		if minSum > tempSum {
			minSum = tempSum
		}
	}

	return maxSum - minSum
}
