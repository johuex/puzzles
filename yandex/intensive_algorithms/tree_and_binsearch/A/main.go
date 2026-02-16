package a

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
	a, _ := strconv.Atoi(lineArr[0])
	b, _ := strconv.Atoi(lineArr[1])
	s, _ := strconv.Atoi(lineArr[2])

	// logic
	ans := shelf(a, b, s)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

// left binary search
func lBinSearch(l, r int, check func(int, []int) bool, params []int) int {
	for l < r {
		m := (l + r) / 2
		if check(m, params) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// l -- square sideб enumeration
func square(l int, params []int) bool {
	a, b, s := params[0], params[1], params[2]
	return l*l >= l*(a+b)-a*b+s
}

func shelf(a, b, s int) int {
	rightBorder := max(a, b) * 1000
	minBorder := min(a, b)
	params := []int{a, b, s}
	res := lBinSearch(minBorder, rightBorder, square, params)

	if res*res != res*(a+b)-a*b+s {
		return -1
	}

	return res
}
