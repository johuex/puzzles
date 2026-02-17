package goodraw

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

	alphabet := make([][]int, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		cnv, _ := strconv.Atoi(line)
		alphabet[i] = []int{i, cnv}
	}

	res := goodRow2(alphabet)

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}

func goodRow(alphabet [][]int) int {
	//BUG: out of memory
	if len(alphabet) == 0 || len(alphabet) == 1 {
		return 0
	}

	// fill chain by first elem
	chunk := alphabet[0]
	chains := make([][]int, chunk[1])
	actualChains := make(map[int]struct{}) // save index of actual chains for better perfomance
	for i := range chunk[1] {
		chains[i] = []int{chunk[0]}
		actualChains[i] = struct{}{}
	}

	// process other "runes"
	for _, chunk := range alphabet[1:] {
		ru := chunk[0]
		cnt := chunk[1]

		// check actual chains
		for idx := range actualChains {
			if chains[idx][len(chains[idx])-1]-ru == -1 {
				chains[idx] = append(chains[idx], ru)
				cnt--
				if cnt == 0 {
					break
				}
			} else {
				// chain become not actual
				delete(actualChains, idx)
			}
		}
		for cnt > 0 {
			chains = append(chains, []int{ru})
			actualChains[len(chains)-1] = struct{}{}
			cnt--
		}
	}

	// count good of raw
	var res int
	for _, chain := range chains {
		res += len(chain) - 1
	}

	return res
}

func goodRow2(alphabet [][]int) int {
	var prev, cur, total int

	for _, batch := range alphabet {
		cur = batch[1]
		total += min(cur, prev)
		prev = cur
	}

	return total
}
