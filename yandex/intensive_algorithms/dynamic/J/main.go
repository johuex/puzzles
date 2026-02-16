package j

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Store struct {
	RPrice int // P
	OThrs  int // R
	OPrice int // Q
	MTotal int // F
}

func main() {
	// general
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, _ := reader.ReadString('\n')
	lineArr := strings.Split(strings.TrimSpace(line), " ")
	n, _ := strconv.Atoi(lineArr[0])
	l, _ := strconv.Atoi(lineArr[1])

	stores := make([]Store, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		lineArr = strings.Split(strings.TrimSpace(line), " ")
		p, _ := strconv.Atoi(lineArr[0])
		r, _ := strconv.Atoi(lineArr[1])
		q, _ := strconv.Atoi(lineArr[2])
		f, _ := strconv.Atoi(lineArr[3])
		stores[i] = Store{
			RPrice: p,
			OThrs:  r,
			OPrice: q,
			MTotal: f,
		}
	}

	// logic
	ans, ansArr := masquerade(n, l, stores)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
	if ans != -1 {
		ansArrStr := make([]string, n)
		for i := range n {
			ansArrStr[i] = strconv.Itoa(ansArr[i])
		}
		writer.WriteString(strings.Join(ansArrStr, " "))
		writer.WriteByte('\n')
	}
}

func price(m int, store Store) int {
	// simply return price for m meters
	if m > store.MTotal {
		// return max impossible price
		return 2000
	}
	if m >= store.OThrs {
		return store.OPrice * m
	}
	return store.RPrice * m
}

func masquerade(n, l int, stores []Store) (int, []int) {
	costTable := make([][]int, n) // raw -- store, column -- cloth
	getStore := make([]int, n)    // how many items we got in each store
	for i := range n {
		costTable[i] = make([]int, l+1)
	}
	// fill first store prices
	for i := 0; i < l+1; i++ {
		costTable[0][i] = price(i, stores[0])
	}

	// calc all possible costs
	for i := 1; i < n; i++ {
		for j := 0; j < l+1; j++ {
			costTable[i][j] = price(j, stores[i])
			for b := 0; b <= j; b++ {
				pretendCost := price(b, stores[i]) + costTable[i-1][j-b]
				minVal := min(costTable[i][j], pretendCost)
				if minVal < costTable[i][j] {
					costTable[i][j] = minVal
					getStore[i-1] = j - b
					getStore[i] = b
				}
			}
		}
	}

	// found min with L meters
	tmp := make([]int, n)
	for i := range n {
		tmp[i] = costTable[i][l]
	}
	ans := slices.Min(tmp)
	if ans >= 2000 {
		return -1, []int{}
	}

	// we can bought in last store more that required
	if getStore[n-1] < stores[n-1].OThrs {
		alreadyBought := getStore[n-1]
		for i := 1; i <= stores[n-1].OThrs-alreadyBought; i++ {
			val := price(getStore[n-2]-1, stores[n-2]) + price(getStore[n-1]+i, stores[n-1])
			if val < ans {
				ans = val
				getStore[n-1] += i
				getStore[n-2] -= 1
			}
		}
	}

	return ans, getStore
}
