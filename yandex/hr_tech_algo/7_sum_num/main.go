package sumnum

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
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	cars := make([]int, n)
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts = strings.Split(line, " ")

	for i := range n {
		tmp, _ := strconv.Atoi(parts[i])
		cars[i] = tmp
	}

	res := sumCars(k, cars)

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}

func sumCars(target int, cars []int) int {
	res := 0
	sum := 0
	left := 0

	for right := 0; right < len(cars); right++ {
		sum += cars[right]

		// move left pointer if sum greater than target
		for sum > target && left <= right {
			sum -= cars[left]
			left++
		}

		if sum == target {
			res++
		}
	}

	return res
}
