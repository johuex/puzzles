package j

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

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr = strings.Split(line, " ")
	candidates := make([]int, n)
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		candidates[i] = converted
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr = strings.Split(line, " ")
	signups := make([]int, n)
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		signups[i] = converted
	}

	// logic
	ans := interview(candidates, signups)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func lBinSearch(l, r int, check func(int, []any) bool, params []any) int {
	for l < r {
		m := (l + r) / 2
		if check(m, params) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func slotsAndCandidetes(cIdx, sIdx int, candidates, slots []int) {
	if slots[sIdx] >= candidates[cIdx] {
		// кол-во слотов больше или равно кол-во кандидатов
		slots[sIdx] -= candidates[cIdx]
		candidates[cIdx] = 0
	} else {
		// кол-во слотов меньше, чем кандидатов
		candidates[cIdx] -= slots[sIdx]
		slots[sIdx] = 0
	}
}

func signup(k int, params []any) bool {
	candidates := append([]int(nil), params[0].([]int)...)
	slots := append([]int(nil), params[1].([]int)...)

	slotIdx := 0
	for i := range len(candidates) {
		slotIdx = max(slotIdx, i-k)
		for slotIdx <= min(len(slots)-1, i+k) {
			if candidates[i] == 0 {
				break
			}
			if slots[slotIdx] != 0 {
				slotsAndCandidetes(i, slotIdx, candidates, slots)
			} else {
				slotIdx += 1
			}
		}
		if candidates[i] > 0 {
			return false
		}
	}
	return true
}

func interview(candidates, slots []int) int {
	left := 0
	right := len(candidates)
	params := []any{candidates, slots}

	res := lBinSearch(left, right, signup, params)

	if !signup(res, params) {
		return -1
	}
	return res
}
