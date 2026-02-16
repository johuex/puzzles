package c

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

	// logic
	ans := passwords(line)
	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func passwords(password string) int {
	// my onw solution
	n := len(password)
	if n <= 1 {
		return 1
	}

	freq := make(map[rune]int)
	for _, ch := range password {
		freq[ch]++
	}

	totalPairs := n * (n - 1) / 2
	samePairs := 0
	for _, f := range freq {
		samePairs += f * (f - 1) / 2
	}
	unique := totalPairs - samePairs
	return unique + 1
}

func passwords1(password string) int {
	// solution from answer
	n := len(password)
	freq := make(map[rune]int)
	for _, ch := range password {
		freq[ch]++
	}
	var ans int
	for _, val := range freq {
		ans += val * (n - val)
	}
	return ans/2 + 1
}
