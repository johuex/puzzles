package maxmul

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
	lineArr := strings.Split(line, " ")

	n := len(lineArr)
	nums := make([]int, n)
	for i := range n {
		num, _ := strconv.Atoi(lineArr[i])
		nums[i] = num
	}

	res := MaxMul(nums)

	resArr := make([]string, len(res))
	for idx := range res {
		resArr[idx] = strconv.Itoa(res[idx])
	}
	writer.WriteString(strings.Join(resArr, " "))
	writer.WriteByte('\n')

}

func MaxMul(nums []int) []int {
	// need to found:
	// two max num
	// two min num
	// compare mul and choose the greatest
	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return []int{nums[0], nums[1]}
		}
		return []int{nums[1], nums[0]}
	}

	minVal := make([]int, 2)
	maxVal := make([]int, 2)

	minVal[0] = nums[0]
	minVal[1] = nums[1]
	maxVal[0] = nums[0]
	maxVal[1] = nums[1]

	for i := 2; i < len(nums); i++ {
		if nums[i] > maxVal[0] {
			if nums[i] > maxVal[1] {
				maxVal[0] = maxVal[1]
				maxVal[1] = nums[i]
			} else {
				maxVal[0] = nums[i]
			}
		} else if nums[i] < minVal[1] {
			if nums[i] < minVal[0] {
				minVal[1] = minVal[0]
				minVal[0] = nums[i]
			} else {
				minVal[1] = nums[i]
			}
		}
	}

	if minVal[0] > minVal[1] {
		minVal[1], minVal[0] = minVal[0], minVal[1]
	}

	if maxVal[0] > maxVal[1] {
		maxVal[1], maxVal[0] = maxVal[0], maxVal[1]
	}

	if minVal[0]*minVal[1] > maxVal[0]*maxVal[1] {
		return minVal
	}

	return maxVal
}
