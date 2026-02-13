package turismhigh

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

	points := make([][]int, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lineArr := strings.Split(line, " ")
		x, _ := strconv.Atoi(lineArr[0])
		y, _ := strconv.Atoi(lineArr[1])
		points[i] = []int{x, y}
	}

	line, _ = reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(line))
	routes := make([][]int, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lineArr := strings.Split(line, " ")
		l, _ := strconv.Atoi(lineArr[0])
		r, _ := strconv.Atoi(lineArr[1])
		routes[i] = []int{l, r}
	}

	res := turism(points, routes)

	resArr := make([]string, len(res))
	for idx := range res {
		resArr[idx] = strconv.Itoa(res[idx])
	}
	writer.WriteString(strings.Join(resArr, "\n"))
	writer.WriteByte('\n')

}

func turism(points, routes [][]int) []int {
	n := len(points)
	if n == 0 {
		return make([]int, len(routes))
	}

	prefixLeft2Right := make([]int, n)
	for i := 1; i < n; i++ {
		prefixLeft2Right[i] = prefixLeft2Right[i-1]
		if points[i][1] > points[i-1][1] {
			prefixLeft2Right[i] += points[i][1] - points[i-1][1]
		}
	}

	prefixRight2Left := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		prefixRight2Left[i] = prefixRight2Left[i+1]
		if points[i][1] > points[i+1][1] {
			prefixRight2Left[i] += points[i][1] - points[i+1][1]
		}
	}

	res := make([]int, len(routes))

	for idx, route := range routes {
		s, e := route[0]-1, route[1]-1

		if s == e {
			res[idx] = 0
			continue
		}

		if s < e {
			res[idx] = prefixLeft2Right[e] - prefixLeft2Right[s]
		} else {
			res[idx] = prefixRight2Left[e] - prefixRight2Left[s]
		}

		if res[idx] < 0 {
			res[idx] = 0
		}
	}

	return res
}
