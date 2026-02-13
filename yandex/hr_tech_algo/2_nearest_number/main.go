package nearestnumber

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
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, " ")

	nums := make([]int, n)
	for i := range n {
		num, _ := strconv.Atoi(lineArr[i])
		nums[i] = num
	}

	line, _ = reader.ReadString('\n')
	target, _ := strconv.Atoi(strings.TrimSpace(line))

	res := nearest(nums, target)

	ans := strconv.Itoa(res)
	writer.WriteString(ans)
	writer.WriteByte('\n')

}

func nearest(nums []int, target int) int {
	// arr can be unsorted
	// so we search numbre with lowest diff
	diff := Abs(target - nums[0])
	neig := nums[0]

	for _, num := range nums {
		tmpDiff := Abs(target - num)
		if tmpDiff < diff {
			neig = num
			diff = tmpDiff
		}
	}

	return neig
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
