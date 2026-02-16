package b

import (
	"bufio"
	"container/list"
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

	relates := make([][]int, n-1)
	for i := range n - 1 {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lineArr := strings.Split(line, " ")
		first, _ := strconv.Atoi(lineArr[0])
		second, _ := strconv.Atoi(lineArr[1])
		relates[i] = []int{first, second}
	}

	// logic
	ans := logic(n, relates)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func buildGraph(n int, relates [][]int) []*TreeLeaf {
	leafes := make([]*TreeLeaf, n)
	for i := range leafes {
		leafes[i] = &TreeLeaf{
			val:    i,
			childs: make([]*TreeLeaf, 0),
		}
	}
	for _, rel := range relates {
		a, b := rel[0]-1, rel[1]-1
		leafes[a].childs = append(leafes[a].childs, leafes[b])
		leafes[b].childs = append(leafes[b].childs, leafes[a])
	}
	return leafes
}

func findEdgeLeafes(leafes []*TreeLeaf) []*TreeLeaf {
	res := make([]*TreeLeaf, 0)
	for _, leaf := range leafes {
		if len(leaf.childs) == 1 {
			res = append(res, leaf)
		}
	}
	return res
}

func shortestDist(edgeLeafes, leafes []*TreeLeaf) int {
	// launch BFS waves from edgeLeafes
	// intersection of leafes -- the shortest path
	n := len(leafes)
	distance := make([]int, n)
	source := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = -1
		source[i] = -1
	}

	queue := list.New()
	for _, leaf := range edgeLeafes {
		idx := leaf.val
		queue.PushBack(leaf)
		distance[idx] = 0
		source[idx] = idx
	}

	for queue.Len() > 0 {
		v := queue.Remove(queue.Front()).(*TreeLeaf)
		for _, u := range v.childs {
			if distance[u.val] == -1 {
				distance[u.val] = distance[v.val] + 1
				source[u.val] = source[v.val]
				queue.PushBack(u)
			} else if source[u.val] != source[v.val] {
				// BFS waves have intersection
				return distance[u.val] + distance[v.val] + 1
			}
		}
	}

	return -1 // need for compilation restriction
}

func logic(n int, relates [][]int) int {
	leafes := buildGraph(n, relates)

	edgeLeafes := findEdgeLeafes(leafes)

	return shortestDist(edgeLeafes, leafes)
}
