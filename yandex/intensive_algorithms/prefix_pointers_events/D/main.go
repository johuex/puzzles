package d

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

	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	lineArr := strings.Split(lineStr, " ")
	cheese := make([]int, n)
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		cheese[i] = converted
	}

	// logic
	ans := friends(cheese)

	// output
	ansStr := make([]string, 3)
	for i := range 3 {
		ansStr[i] = strconv.Itoa(ans[i])
	}
	writer.WriteString(strings.Join(ansStr, " "))
	writer.WriteByte('\n')
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func friends(cheese []int) []int {
	// pointers
	left := 0
	right := len(cheese) - 1

	// cnt of eaten cheese
	leftSum := cheese[0]
	rightSum := cheese[right]

	res := []int{Abs(rightSum - leftSum), left, right}

	for right-left > 1 {
		if leftSum > rightSum {
			right -= 1
			rightSum += cheese[right]
		} else {
			left += 1
			leftSum += cheese[left]
		}
		if Abs(leftSum-rightSum) <= res[0] {
			res = []int{Abs(leftSum - rightSum), left, right}
		}
	}

	res[1] += 1
	res[2] += 1
	return res
}
