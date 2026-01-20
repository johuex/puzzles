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
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// logic
	ans := winner(n)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func winner(n int) int {
	if n%4 == 0 {
		return 2
	}
	return 1
}

// check that number is prime
func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// solving task with DP
func winnerDP(n int) int {
	dp := make([]bool, n+1)
	dp[0] = false

	for i := 1; i <= n; i++ {
		dp[i] = false
		for k := 1; k <= 3; k++ {
			if i-k >= 0 && !isPrime(i-k) && !dp[i-k] {
				dp[i] = true
				break
			}
		}
	}

	if dp[n] {
		return 1
	}
	return 2
}
