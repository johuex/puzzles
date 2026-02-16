package niceraw

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
	k, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	res := niceRaw(k, line)

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}

var words = "qwertyuiopasdfghjklzxcvbnm"

func niceRaw(k int, raw string) int {
	maxLen := 0
	runes := []rune(raw)

	for _, target := range words {
		left := 0
		replacements := 0

		for right := 0; right < len(runes); right++ {
			if runes[right] != target {
				replacements++
			}

			for replacements > k {
				if runes[left] != target {
					replacements--
				}
				left++
			}

			maxLen = max(maxLen, right-left+1)
		}
	}

	return maxLen
}
