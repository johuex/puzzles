package pointssegments

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
	lineArr := strings.Split(line, " ")
	n, _ := strconv.Atoi(lineArr[0])
	m, _ := strconv.Atoi(lineArr[1])

	segments := make([][]int, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lineArr := strings.Split(line, " ")
		l, _ := strconv.Atoi(lineArr[0])
		r, _ := strconv.Atoi(lineArr[1])
		segments[i] = []int{l, r}
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	lineArr = strings.Split(line, " ")
	points := make([]int, m)
	for i := range m {
		tmp, _ := strconv.Atoi(lineArr[i])
		points[i] = tmp
	}

	res := cntPoints(segments, points)

	resArr := make([]string, len(res))
	for i := range len(res) {
		resArr[i] = strconv.Itoa(res[i])
	}
	writer.WriteString(strings.Join(resArr, " "))
	writer.WriteByte('\n')
}

type Event struct {
	Coord int
	CType int
	idx   int // для точек
}

func cntPoints(segments [][]int, points []int) []int {
	// -1 -- конец отрезка, 0 -- точка запроса, 1 -- начало отрезка
	res := make([]int, len(points))

	events := make([]Event, 0, len(points)+len(segments)*2)

	for _, segment := range segments {
		start, end := segment[0], segment[1]
		if start <= end {
			events = append(events, Event{Coord: start, CType: 1})
			events = append(events, Event{Coord: end, CType: -1})
		} else {
			// swap point if order is invalid
			events = append(events, Event{Coord: end, CType: 1})
			events = append(events, Event{Coord: start, CType: -1})
		}
	}

	for idx, point := range points {
		events = append(events, Event{Coord: point, CType: 0, idx: idx})
	}

	// first: by coord, than by event
	sort.Slice(events, func(i, j int) bool {
		if events[i].Coord != events[j].Coord {
			return events[i].Coord < events[j].Coord
		}
		return events[i].CType > events[j].CType
	})

	var cnt int
	for _, event := range events {
		switch event.CType {
		case 1:
			cnt++
		case 0:
			res[event.idx] = cnt
		case -1:
			cnt--
		}
	}

	return res
}
