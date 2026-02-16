package e

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type TreeLeaf struct {
	val    int
	childs []*TreeLeaf
}

func main() {
	// general
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n, _ := strconv.Atoi(line)

	relates := make([]int, n)
	relates[0] = -1 // it's root
	for i := 1; i < n; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		converted, _ := strconv.Atoi(line)
		relates[i] = converted
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, " ")

	values := make([]int, n)
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		values[i] = converted
	}

	// logic
	ans := taxes(relates, values)

	// output

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func buildTree(relates, values []int) ([]*TreeLeaf, *TreeLeaf) {
	// build tree on leafes and relations, return root and list
	var root *TreeLeaf
	leafes := make([]*TreeLeaf, len(relates))
	for i := range leafes {
		leafes[i] = &TreeLeaf{
			val:    values[i],
			childs: make([]*TreeLeaf, 0),
		}
	}
	for i, rel := range relates {
		if rel == -1 {
			root = leafes[i]
			continue
		}
		parent := leafes[rel]
		parent.childs = append(parent.childs, leafes[i])
	}
	return leafes, root
}

func minTaxSum(leaf *TreeLeaf, inc, dec int) (int, int) {
	tempInc := inc
	tempDec := dec
	for _, child := range leaf.childs {
		rInc, rDec := minTaxSum(child, tempInc, tempDec)
		inc += rInc
		dec += rDec

	}
	if inc == 0 && dec == 0 {
		// it's leaf
		if leaf.val < 0 {
			inc = Abs(leaf.val)
		} else {
			dec = -Abs(leaf.val)
		}
		return inc, dec
	}
	if leaf.val != 0 {
		diff := leaf.val + inc + dec
		if diff > 0 {
			dec -= diff
		} else if diff < 0 {
			inc -= diff
		}
	}
	return inc, dec
}

func taxes(relates, values []int) int {
	_, treeRoot := buildTree(relates, values)

	resInc, resDec := minTaxSum(treeRoot, 0, 0)

	return resInc - resDec
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
