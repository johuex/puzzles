package sinonims

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

	words := make(map[string]string) // for storing synonims
	for range n {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lineArr := strings.Split(line, " ")
		// store both for faster search
		words[lineArr[0]] = lineArr[1]
		words[lineArr[1]] = lineArr[0]
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	res := words[line]

	writer.WriteString(res)
	writer.WriteByte('\n')

}
