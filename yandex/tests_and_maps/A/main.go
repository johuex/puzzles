package a

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
	n, _ := strconv.Atoi(strings.TrimSpace(line))
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, " ")
	weight := make([]int, n)
	for i := 0; i < len(lineArr); i++ {
		conv, _ := strconv.Atoi(lineArr[i])
		weight[i] = conv
	}

	// logic
	// нужно найти максимум четных и минимум нечетных и свапнуть их
	// если же максимум <= минимуму, то ничего менять не нужно
	// параллельно можно считать сумму для нечетных и сумму четных а затем на основе условия swap пересчитать суммы и вывести индекс радости
	ans := mushroomsJoy(n, weight)
	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func mushroomsJoy(n int, weight []int) int {
	oddSum := 0  // Вася (0,2,..)
	evenSum := 0 // Маша (1,3,...)
	// n <= 2, ок
	oddMin := weight[0]
	evenMax := weight[1]
	for i := range n {
		if i%2 == 0 {
			// считаем Васю
			oddSum += weight[i]
			oddMin = min(oddMin, weight[i])
		} else {
			// считаем Машу
			evenSum += weight[i]
			evenMax = max(evenMax, weight[i])
		}
	}

	var ans int
	if oddMin <= evenMax {
		ans = (oddSum - oddMin + evenMax) - (evenSum - evenMax + oddMin)
	} else {
		ans = oddSum - evenSum
	}
	return ans
}
