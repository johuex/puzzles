package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, " ")
	n, _ := strconv.Atoi(lineArr[0])
	m, _ := strconv.Atoi(lineArr[1])

	segments := make([][]int, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lineArr := strings.Split(line, " ")
		l, _ := strconv.Atoi(lineArr[0])
		r, _ := strconv.Atoi(lineArr[1])
		segments[i] = []int{l, r}
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr = strings.Split(line, " ")
	points := make([]int, m)
	for i := range m {
		tmp, _ := strconv.Atoi(lineArr[i])
		points[i] = tmp
	}

	res := cntPoints2(segments, points)

	resArr := make([]string, len(res))
	for i := range len(res) {
		resArr[i] = strconv.Itoa(res[i])
	}
	writer.WriteString(strings.Join(resArr, " "))
	writer.WriteByte('\n')
}

func cntPoints(segments [][]int, points []int) []int {
	res := make([]int, len(points))
	for i := range len(res) {
		for j := range len(segments) {
			l := segments[j][0]
			r := segments[j][1]

			if min(l, r) <= points[i] && points[i] <= max(l, r) {
				res[i] += 1
			}
		}
	}
	return res
}

func cntPoints2(segments [][]int, points []int) []int {
	// TODO: переписать на event ???
	mergeArr := make([]int, 0, len(segments)*2+len(points)) // all points                                           // current position
	left := make(map[int]struct{}, len(segments))
	right := make(map[int]struct{}, len(segments))
	po := make(map[int]struct{}, len(points))

	uniqValues := make(map[int]struct{})

	for i := range len(segments) {
		l := segments[i][0]
		r := segments[i][1]
		left[l] = struct{}{}
		right[r] = struct{}{}
		if _, ok := uniqValues[l]; !ok {
			mergeArr = append(mergeArr, l)
			uniqValues[l] = struct{}{}
		}
		if _, ok := uniqValues[r]; !ok {
			mergeArr = append(mergeArr, r)
			uniqValues[r] = struct{}{}
		}
	}

	for _, point := range points {
		if _, ok := uniqValues[point]; !ok {
			mergeArr = append(mergeArr, point)
			uniqValues[point] = struct{}{}
		}
		po[point] = struct{}{}
	}

	slices.Sort(mergeArr)

	var actualCnt int
	res := make([]int, 0, len(points))
	for _, elem := range mergeArr {
		_, lOk := left[elem]
		_, rOk := right[elem]
		_, pOk := po[elem]
		if lOk {
			actualCnt += 1
		}
		if pOk {
			res = append(res, actualCnt)
		}
		if rOk {
			actualCnt -= 1
		}
	}

	return res
}
