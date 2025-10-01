package b

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
	b, _ := strconv.Atoi(lineArr[0])
	c, _ := strconv.Atoi(lineArr[0])
	v0, _ := strconv.Atoi(lineArr[0])
	v1, _ := strconv.Atoi(lineArr[0])
	v2, _ := strconv.Atoi(lineArr[0])

	// logic
	ans := route(a, b, c, v0, v1, v2)
	// output
	writer.WriteString(ans)
	writer.WriteByte('\n')
}

func route(a, b, c, v0, v1, v2 int) string {
	r1 := float64(b)/float64(v0) + float64(c)/float64(v1) + float64(a)/float64(v2)
	r2 := float64(b)/float64(v0) + float64(b)/float64(v1) + float64(a)/float64(v0) + float64(a)/float64(v1)
	r3 := float64(b)/float64(v0) + float64(b)/float64(v1) + float64(b)/float64(v0) + float64(c)/float64(v0) + float64(c)/float64(v1) + float64(b)/float64(v1)
	r4 := float64(a)/float64(v0) + float64(a)/float64(v1) + float64(a)/float64(v0) + float64(c)/float64(v0) + float64(c)/float64(v1) + float64(a)/float64(v1)
	r5 := float64(b)/float64(v0) + float64(c)/float64(v1) + float64(c)/float64(v2) + float64(b)/float64(v2)
	r6 := float64(a)/float64(v0) + float64(c)/float64(v1) + float64(c)/float64(v2) + float64(a)/float64(v2)
	r7 := float64(a)/float64(v0) + float64(c)/float64(v1) + float64(b)/float64(v2)
	r8 := float64(b)/float64(v0) + float64(b)/float64(v1) + float64(a)/float64(v0) + float64(c)/float64(v0) + float64(c)/float64(v1) + float64(a)/float64(v1)
	r9 := float64(a)/float64(v0) + float64(a)/float64(v1) + float64(b)/float64(v0) + float64(c)/float64(v0) + float64(c)/float64(v1) + float64(b)/float64(v1)
	minL := min(r1, r2, r3, r4, r5, r6, r7, r8, r9)
	return strconv.FormatFloat(minL, 'f', 15, 64)
}
