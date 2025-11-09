package j

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type TreeLeaf struct {
	in     int
	out    int
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

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr := strings.Split(line, " ")

	relates := make([]int, n)
	for i := range n {
		converted, _ := strconv.Atoi(lineArr[i])
		relates[i] = converted
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	m, _ := strconv.Atoi(line)

	asks := make([][]int, m)
	for i := range m {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lineArr = strings.Split(line, " ")
		converted, _ := strconv.Atoi(lineArr[0])
		converted1, _ := strconv.Atoi(lineArr[1])
		asks[i] = []int{converted, converted1}
	}

	// logic
	ans := parent(relates, asks)

	// output
	ansStr := make([]string, m)
	for i := range m {
		ansStr[i] = strconv.Itoa(ans[i])
	}

	writer.WriteString(strings.Join(ansStr, "\n"))
	writer.WriteByte('\n')
}

func buildTree(relates []int) ([]*TreeLeaf, *TreeLeaf) {
	// build tree on leafes and relations, return root and list
	var root *TreeLeaf
	leafes := make([]*TreeLeaf, len(relates))
	for i := range leafes {
		leafes[i] = &TreeLeaf{
			val:    i + 1,
			childs: make([]*TreeLeaf, 0),
		}
	}
	for i, rel := range relates {
		if rel == 0 {
			root = leafes[i]
			continue
		}
		parent := leafes[rel-1]
		parent.childs = append(parent.childs, leafes[i])
	}
	return leafes, root
}

func cycleTree(leaf *TreeLeaf, count int) int {
	leaf.in = count
	count += 1
	for _, child := range leaf.childs {
		count = cycleTree(child, count)
	}
	count += 1
	leaf.out = count
	return count

}

func parent(relates []int, asks [][]int) []int {
	leafes, treeRoot := buildTree(relates)

	cycleTree(treeRoot, 0)

	res := make([]int, len(asks))
	for idx, ask := range asks {
		parent := leafes[ask[0]-1]
		child := leafes[ask[1]-1]

		if parent.in <= child.in && parent.out >= child.out {
			res[idx] = 1
		} else {
			res[idx] = 0
		}

	}

	return res
}
