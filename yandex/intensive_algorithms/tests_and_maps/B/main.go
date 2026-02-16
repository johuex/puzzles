package b

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
	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, " ")
	a, _ := strconv.Atoi(lineArr[0])
	b, _ := strconv.Atoi(lineArr[1])
	c, _ := strconv.Atoi(lineArr[2])
	v0, _ := strconv.Atoi(lineArr[3])
	v1, _ := strconv.Atoi(lineArr[4])
	v2, _ := strconv.Atoi(lineArr[5])

	// logic
	ans := route(a, b, c, v0, v1, v2)
	// output
	writer.WriteString(ans)
	writer.WriteByte('\n')
}

func route(a, b, c, v0, v1, v2 int) string {
	// not my solution, this is answer

	// релаксация ребер (https://habr.com/ru/articles/201588/)
	// релаксация ребер -- выбор наименьшего из пути
	a = min(a, b+c)
	b = min(b, a+c)
	c = min(c, a+b)

	// home - market - home - post - home
	ans := float64(a)/float64(v0) + float64(a)/float64(v1) + float64(b)/float64(v0) + float64(b)/float64(v1)
	// home - market - post - home
	ans = min(ans, float64(a)/float64(v0)+float64(c)/float64(v1)+float64(b)/float64(v2))
	// home - post - market - home
	ans = min(ans, float64(b)/float64(v0)+float64(c)/float64(v1)+float64(a)/float64(v2))
	return fmt.Sprintf("%.15f", ans)
}
