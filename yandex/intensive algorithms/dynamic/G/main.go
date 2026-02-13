package g

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
	ans := stairs(n)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func stairs(n int) int {
	res := make([]int, n+1)
	res[0] = 1

	for k := 1; k <= n; k++ {
		for i := n; i >= k; i-- {
			res[i] += res[i-k]
		}
	}
	return res[n]
}
