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
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// logic
	ans := ballJump(n)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func ballJump(n int) int {
	res := make([]int, n)
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	}
	res[0] = 1
	res[1] = 2
	res[2] = 4
	for i := 3; i < n; i++ {
		res[i] = res[i-1] + res[i-2] + res[i-3]
	}
	return res[len(res)-1]
}
