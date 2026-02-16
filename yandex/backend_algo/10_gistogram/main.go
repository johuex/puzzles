package gistogram

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
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])

	heights := make([]int, n)
	for i := 0; i < n; i++ {
		heights[i], _ = strconv.Atoi(parts[i+1])
	}

	res := gistogram(heights)

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}

type stackElem struct {
	Val int
	Idx int
}

func gistogram(heights []int) int {
	switch len(heights) {
	case 0:
		return 0
	case 1:
		return heights[0]
	case 2:
		return 2 * (min(heights[0], heights[1]))
	}

	var res int
	stack := make([]stackElem, 0)

	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 || heights[i] > stack[len(stack)-1].Val {
			stack = append(stack, stackElem{heights[i], i})
		} else {
			// we found right border -- elem that less that elem in stack
			for len(stack) > 0 && heights[i] < stack[len(stack)-1].Val {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// left border -  idx of prev elem in stack or -1
				leftIndex := -1
				if len(stack) > 0 {
					leftIndex = stack[len(stack)-1].Idx
				}

				// width = current idx - left border - 1
				width := i - leftIndex - 1
				area := top.Val * width
				if area > res {
					res = area
				}
			}
			stack = append(stack, stackElem{heights[i], i})
		}
	}

	// process left elems in stack
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		leftIndex := -1
		if len(stack) > 0 {
			leftIndex = stack[len(stack)-1].Idx
		}

		// right border == length
		width := len(heights) - leftIndex - 1
		area := top.Val * width
		if area > res {
			res = area
		}
	}

	return res
}
