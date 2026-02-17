package three1

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
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// prev0 -- how many rows end with 0
	// prev1 -- how many rows end with 01
	// prev2 -- how many rows end with 011
	prev0, prev1, prev2 := 1, 1, 0 // base fo DP, N = 1 (00, 01, NaN)
	// they used as counters
	// example for 8:
	// curr0: "000","100","010","110"
	// curr1: "001","101"
	// curr2: "011"

	for i := 2; i <= n; i++ {
		curr0 := prev0 + prev1 + prev2
		curr1 := prev0
		curr2 := prev1

		prev0, prev1, prev2 = curr0, curr1, curr2
	}

	answer := prev0 + prev1 + prev2

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
