package h

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
	m, _ := strconv.Atoi(lineArr[1])

	line, _ = reader.ReadString('\n')
	targetLine := strings.TrimSpace(line)

	lines := make([]string, m)
	for i := range m {
		line, _ := reader.ReadString('\n')
		lines[i] = strings.TrimSpace(line)
	}
	// logic
	ans := glue(n, m, targetLine, lines)
	// output
	writer.WriteString(strings.Join(ans, " "))
	writer.WriteByte('\n')
}

func glue(n, m int, targetLine string, inputParts []string) []string {
	// split target line to chunks
	inputPartsMap := orderParts(inputParts)

	var chunks []string
	chunkSize := n / m
	for i := 0; i < n; i += chunkSize {
		end := i + chunkSize
		if end > n {
			end = n
		}
		chunks = append(chunks, targetLine[i:end])
	}

	res := make([]string, len(chunks))
	for idx, chunk := range chunks {
		// always take latest duplicate
		elems := inputPartsMap[chunk]
		res[idx] = strconv.Itoa(elems[len(elems)-1])
		inputPartsMap[chunk] = elems[:len(elems)-1]
	}

	return res
}

// need to store duplicates
func orderParts(parts []string) map[string][]int {
	res := make(map[string][]int)
	for idx, part := range parts {
		if _, ok := res[part]; !ok {
			res[part] = make([]int, 0)
		}
		res[part] = append(res[part], idx+1)
	}
	return res
}
