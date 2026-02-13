package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	res := countWords(reader)

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}

func countWords(reader *bufio.Reader) int {
	wordMap := make(map[string]struct{})

	delimiter := map[rune]struct{}{
		' ':  {},
		'\n': {},
		'\t': {},
		'\r': {},
	}
	tmpWord := []rune{}

	for {
		word, _, err := reader.ReadRune()
		if errors.Is(err, io.EOF) {
			break
		}
		if _, ok := delimiter[word]; !ok {
			tmpWord = append(tmpWord, word)
		} else {
			if len(tmpWord) > 0 {
				wordMap[string(tmpWord)] = struct{}{}
			}
			tmpWord = []rune{}
		}
	}
	if len(tmpWord) > 0 {
		wordMap[string(tmpWord)] = struct{}{}
	}
	return len(wordMap)
}
