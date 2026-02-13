package graphrelations

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	edges := make([][2]int, m)
	for i := 0; i < m; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts = strings.Split(line, " ")
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		edges[i] = [2]int{u, v}
	}

	res := graphSync(n, edges)

	sort.Ints(res)
	writer.WriteString(strconv.Itoa(len(res)) + "\n")
	for i, v := range res {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(v))
	}
	writer.WriteByte('\n')
}

func graphSync(n int, edges [][2]int) []int {
	// build graph
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// walk through graph from 1
	visited := make([]bool, n+1)
	var res []int
	queue := []int{1}
	visited[1] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		res = append(res, v)

		for _, to := range adj[v] {
			if !visited[to] {
				visited[to] = true
				queue = append(queue, to)
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}
