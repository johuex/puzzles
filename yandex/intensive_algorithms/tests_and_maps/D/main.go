package d

import (
	"bufio"
	"os"
	"sort"
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
	k, _ := strconv.Atoi(lineArr[1])

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr = strings.Split(line, " ")
	taskThemes := make([]int, n)
	for i := 0; i < len(lineArr); i++ {
		conv, _ := strconv.Atoi(lineArr[i])
		taskThemes[i] = conv
	}

	// logic
	ans := tasks(k, taskThemes)
	// output
	ansStr := make([]string, len(ans))
	for i := 0; i < len(ans); i++ {
		ansStr[i] = strconv.Itoa(ans[i])
	}
	writer.WriteString(strings.Join(ansStr, " "))
	writer.WriteByte('\n')
}

func tasks(k int, taskThemes []int) []int {

	ans := make([]int, k)
	ansIdx := 0

	taskThemesCount := make(map[int]int)
	// count tasks per theme
	for _, i := range taskThemes {
		if _, ok := taskThemesCount[i]; !ok {
			taskThemesCount[i] = 0

		}
		taskThemesCount[i] += 1
	}

	// sort them by count
	tSort := make([][]int, 0)
	for key, value := range taskThemesCount {
		tSort = append(tSort, []int{key, value})
	}

	sort.Slice(tSort, func(i, j int) bool {
		return tSort[i][1] < tSort[j][1]
	})

	for ansIdx <= k-1 {
		for j := range len(tSort) {
			themeCount := tSort[j]
			theme := themeCount[0]
			count := themeCount[1]
			if count == 0 {
				continue
			}
			ans[ansIdx] = theme
			ansIdx += 1
			if ansIdx == k {
				break
			}
			tSort[j][1] -= 1 // decr count
		}
		if ansIdx == k {
			break
		}
	}
	return ans
}
