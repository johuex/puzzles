package g

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
	//m, _ := strconv.Atoi(lineArr[1])

	// create table
	table := make([]string, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		table[i] = line
	}

	// logic
	ans := fiveRow(table)
	// output
	writer.WriteString(ans)
	writer.WriteByte('\n')
}

func fiveRow(table []string) string {
	var ans bool

	lineSet := make(map[int]map[rune]struct{})
	vertSet := make(map[int]map[rune]struct{})

	// transform input
	for i, row := range table {
		lineSet[i] = make(map[rune]struct{})
		for j, sym := range row {
			if _, ok := vertSet[j]; !ok {
				vertSet[j] = make(map[rune]struct{})
			}
			vertSet[j][sym] = struct{}{}
			lineSet[i][sym] = struct{}{}
		}
	}
	for _, raw := range lineSet {
		_, xOk := raw['X']
		_, yOk := raw['Y']
		if !(xOk && yOk) {
			ans = true
		}
		break
	}
	if !ans {
		for _, raw := range vertSet {
			_, xOk := raw['X']
			_, yOk := raw['Y']
			if !(xOk && yOk) {
				ans = true
			}
			break
		}
	}
	if !ans {
		// TODO: как проверить диагонали ???
	}

	if ans {
		return "Yes"
	}
	return "No"
}
