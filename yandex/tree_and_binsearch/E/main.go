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

// TODO: algo works only with positive values
// TODO: need to reimplement for also negative values
func minTaxSum(leaf *TreeLeaf, sum int) int {
	tempSum := sum
	for _, child := range leaf.childs {
		sum += minTaxSum(child, tempSum)
	}
	if sum <= leaf.val {
		sum = leaf.val
	} else {
		sum = sum + leaf.val
	}
	return sum
}

func taxes(relates, values []int) int {
	_, treeRoot := buildTree(relates, values)

	res := minTaxSum(treeRoot, 0)

	return res
}
