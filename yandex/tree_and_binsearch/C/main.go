package c

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
	w, _ := strconv.Atoi(lineArr[1])
	h, _ := strconv.Atoi(lineArr[2])

	words := make([][]float64, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lineArr := strings.Split(line, " ")
		first, _ := strconv.ParseFloat(lineArr[0], 64)
		second, _ := strconv.ParseFloat(lineArr[1], 64)
		words[i] = []float64{first, second}
	}

	// logic
	ans := layout(w, h, words)

	// output
	writer.WriteString(ans)
	writer.WriteByte('\n')
}

func lBinSearch(l, r float64, check func(float64, []any) bool, params []any) float64 {
	for r-l > 1e-6 {
		m := (l + r) / 2
		if check(m, params) {
			r = m
		} else {
			l = m
		}
	}
	return (l + r) / 2
}

func scale(k float64, params []any) bool {
	h := params[0].(float64)
	w := params[1].(float64)
	words := params[2].([][]float64)

	actualLineLength := 0.0
	actualHeight := 0.0
	lastWordHeight := k * words[0][1]

	for _, word := range words {
		wordWidth := k * word[0]
		wordHeight := k * word[1]

		if wordWidth > w || wordHeight > h {
			return true // слово не влезает никуда
		}

		if Abs(wordHeight-lastWordHeight) > 1e-6 || actualLineLength+wordWidth > w {
			// новая строка
			actualHeight += lastWordHeight
			actualLineLength = 0
			lastWordHeight = wordHeight
			if actualHeight+wordHeight > h {
				return true
			}
		}
		actualLineLength += wordWidth
	}

	return false
}

func layout(w, h int, words [][]float64) string {
	left := 0.0
	right := float64(w + h)
	params := []any{float64(h), float64(w), words}
	res := lBinSearch(left, right, scale, params)
	return strconv.FormatFloat(res, 'f', 6, 64)
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
