package e

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
	k, _ := strconv.Atoi(lineArr[1])

	// logic
	ans := sumNum(n, k)
	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

// 1 10 -> 1 2 4 8 16 22 24 28 36 42 44
// 2 10 -> 2 4 8 16 22 24 28 36 42 44 48
// 3 10 -> 3 6 12 14 18 26 32 34 38 46 52
// 4 10 -> 4 8 16 22 24 28 36 42 44 48 56
//
// 6 10 -> 6 12 14 18 26 32 34 38 46 52 54
// 7 10 -> 7 14 18 26 32 34 38 46 52 54 58
// 8 10 -> 8 16 22 24 28 36 42 44 48 56 62
// 9 10 -> 9 18 26 32 34 38 46 52 54 58 66
//
// 5 10 -> 5 10 10 10 10 10 10 10 10 10 10
// 10 10 -> 10 10 10 10 10 10 10 10 10 10 10
//
// cycle: 2 4 8 16 ... and etc.
func sumNum(n, k int) int {
	if n%10 == 0 {
		return n
	}
	if n%10 == 5 && k >= 1 {
		return n + 5
	}
	if k == 0 {
		return n
	}

	lastN := n % 10
	f0 := lastN
	f1 := (2 * lastN) % 10
	f2 := (4 * lastN) % 10
	f3 := (8 * lastN) % 10

	if k <= 4 {
		switch k {
		case 1:
			return n + f0
		case 2:
			return n + f0 + f1
		case 3:
			return n + f0 + f1 + f2
		case 4:
			return n + f0 + f1 + f2 + f3
		}
	}

	sumf := f0 + f1 + f2 + f3
	ans := sumf
	cycles := (k - 4) / 4 // first cycle was in sumf
	ans += cycles * 20

	part := (k - 4) % 4 // part of last cycle
	if part > 0 {
		// 6 -- last number in cycle
		offset := (6 * lastN) % 10
		switch part {
		case 1:
			ans += offset
		case 2:
			ans += offset + f1
		case 3:
			ans += offset + f1 + f2
		}
	}

	return n + ans
}

func sumNum1(n, k int) int {
	// solution from answer
	if k > 0 && n%2 != 0 {
		n += n % 10
		k -= 1
	}
	// смысл в том, что есть циклы из четных чисел и мы можем их посчитать
	if n%10 == 0 {
		print(n)
	} else {
		cycles := k / 4
		n += (2 + 4 + 6 + 8) * cycles
		k %= 4
		for range k {
			n += n % 10
		}
	}
	return n
}
