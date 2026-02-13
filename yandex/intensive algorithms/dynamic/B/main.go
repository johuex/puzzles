package b

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

const inf = math.MaxInt32

func main() {
	// general
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, _ := reader.ReadString('\n')
	riverLine := strings.TrimSpace(line)

	// logic
	ans := river(riverLine)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func river(river string) int {
	dpL := 0
	dpR := 1

	for i := 0; i < len(river); i++ {
		var tribL, tribR int
		switch river[i] {
		case 'L':
			tribL = 1
			tribR = 0
		case 'R':
			tribL = 0
			tribR = 1
		case 'B':
			tribL = 1
			tribR = 1
		default:
			tribL = 0
			tribR = 0
		}

		newL := min(dpL+tribL, dpR+1+tribL)
		newR := min(dpR+tribR, dpL+1+tribR)

		if newL > inf {
			newL = inf
		}
		if newR > inf {
			newR = inf
		}

		dpL, dpR = newL, newR
	}
	ans := min(dpR, dpL+1)
	return ans
}
