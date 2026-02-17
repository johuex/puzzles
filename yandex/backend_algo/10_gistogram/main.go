package main

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
	parts := strings.Fields(strings.TrimSpace(line))
	n, _ := strconv.Atoi(parts[0])

	heights := make([]int, n+2) // wrap array by zero from both sides
	for i := 1; i <= n; i++ {
		heights[i], _ = strconv.Atoi(parts[i])
	}
	res := gistogram(heights)

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}

func gistogram(heights []int) int {

	stack := make([]int, 0)
	stack = append(stack, 0) // затычка
	maxArea := 0

	for i := 1; i < len(heights); i++ {
		// Пока текущая высота меньше высоты на вершине стека
		for heights[i] < heights[stack[len(stack)-1]] {
			popedIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Ширина: от текущего ДО элемента стека
			width := i - stack[len(stack)-1] - 1
			area := heights[popedIdx] * width
			maxArea = max(maxArea, area)
		}
		if heights[i] <= heights[stack[len(stack)-1]] {
			stack[len(stack)-1] = i
		} else {
			stack = append(stack, i)
		}
	}

	return maxArea
}
