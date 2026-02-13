package d

import (
	"bufio"
	"math"
	"os"
	"slices"
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
	p, _ := strconv.Atoi(lineArr[1])

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr = strings.Split(line, " ")

	signs := make([]int, len(lineArr))
	for idx, la := range lineArr {
		val, _ := strconv.Atoi(la)
		signs[idx] = val
	}

	// logic
	ans := exchange(p, signs)

	// output
	ansStr := make([]string, n)
	for idx, val := range ans {
		ansStr[idx] = strconv.Itoa(val)
	}
	writer.WriteString(strings.Join(ansStr, " "))
	writer.WriteByte('\n')
}

// left binary search
func lBinSearch(l, r int, check func(int, []int, []int) bool, params []int, signs []int) int {
	for l < r {
		m := (l + r) / 2
		if check(m, params, signs) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func rate(l int, params []int, signs []int) bool {
	target := params[0]
	p := params[1]
	return float64(target) <= float64(p)*float64(signs[l])
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func exchange(p int, signs []int) []int {
	signMap := make(map[int]int)
	for idx, sign := range signs {
		signMap[sign] = idx
	}
	slices.Sort(signs)

	left := 0
	right := len(signs) - 1
	minDiff := math.MaxFloat64
	minIdx := []int{-1, -1}
	for idx, sign := range signs {
		params := []int{sign, p}
		resIdx := lBinSearch(left, right, rate, params, signs)
		if resIdx < 0 || resIdx > len(signs)-1 {
			continue
		}
		candidates := []int{resIdx}
		// maybe last candidate would be better
		if resIdx > 0 {
			candidates = append(candidates, resIdx-1)
		}
		for _, ci := range candidates {
			absRes := Abs(float64(p) - float64(signs[idx])/float64(signs[ci]))
			if absRes < minDiff {
				minIdx[0], minIdx[1] = idx, ci
				minDiff = absRes
			}
		}
	}

	// rewrite by mapped valued
	minIdx[0] = signMap[signs[minIdx[0]]] + 1
	minIdx[1] = signMap[signs[minIdx[1]]] + 1

	return minIdx
}
