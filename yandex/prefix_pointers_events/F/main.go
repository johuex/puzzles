package f

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Numeric interface {
	int | float64
}

type Event struct {
	time  float64
	eType int //1- -> occupied, 2 -> free, 3 -> car arrived
	idx   int // user only for cars
}

func main() {
	// general
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, _ := reader.ReadString('\n')
	lineStr := strings.TrimSpace(line)
	lineArr := strings.Split(lineStr, " ")
	n, _ := strconv.Atoi(lineArr[0])
	m, _ := strconv.Atoi(lineArr[1])
	x, _ := strconv.Atoi(lineArr[2])

	trains := make([][]int, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		lineStr = strings.TrimSpace(line)
		lineArr = strings.Split(lineStr, " ")
		first, _ := strconv.Atoi(lineArr[0])
		second, _ := strconv.Atoi(lineArr[1])
		third, _ := strconv.Atoi(lineArr[2])
		trains[i] = []int{first, second, third}
	}

	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	lineArr = strings.Split(lineStr, " ")
	cars := make([]int, m)
	for i := range m {
		converted, _ := strconv.Atoi(lineArr[i])
		cars[i] = converted
	}

	// logic
	ans := railwayCrossing(trains, cars, x)

	ansStr := make([]string, len(ans))
	for i := range len(ans) {
		ansStr[i] = strconv.FormatFloat(ans[i], 'f', 6, 64)
	}
	// output
	writer.WriteString(strings.Join(ansStr, "\n"))
	writer.WriteByte('\n')
}

func Abs[T Numeric](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func createEvents(trains [][]int, cars []int, x int) []Event {
	events := make([]Event, len(trains)*2+len(cars))
	eIdx := 0

	// type: 1- -> occupied, 2 -> free, 3 -> car arrived
	// event: (time, type, num), num only for cars
	for idx, car := range cars {
		events[eIdx] = Event{float64(car), 3, idx}
		eIdx += 1
	}

	for idx, train := range trains {
		var timeOcc, timeFree float64
		if train[0]-train[1] < 0 {
			// positive train direction
			timeOcc = (float64(x - train[1])) / float64(train[2])
			timeFree = (float64(x - train[0])) / float64(train[2])

		} else {
			// negative train direction
			timeOcc = (float64(train[1] - x)) / float64(train[2])
			timeFree = (float64(train[0] - x)) / float64(train[2])
		}
		events[eIdx] = Event{timeOcc, 1, idx}
		eIdx += 1
		events[eIdx] = Event{timeFree, 2, idx}
		eIdx += 1
	}

	sort.Slice(events, func(i, j int) bool {
		// sort by time and then by event time (arrival earlier than departure)
		if events[i].time != events[j].time {
			return events[i].time < events[j].time
		}
		return events[i].eType < events[j].eType
	})

	return events
}

func railwayCrossing(trains [][]int, cars []int, cross int) []float64 {
	res := make([]float64, len(cars))
	carQueue := make(chan int, len(cars)) // car queue on cross
	events := createEvents(trains, cars, cross)
	var trainCrossOccupied int // how many trains occupied cross at this moment, index on result array

	for _, event := range events {
		switch event.eType {
		case 1:
			// train occupied cross
			trainCrossOccupied += 1
		case 2:
			// train free cross
			trainCrossOccupied -= 1
		case 3:
			// pass car index in queue
			carQueue <- event.idx
		}
		// check car queue and empty it when cross in free
		var emptyQueue bool
		for trainCrossOccupied == 0 && !emptyQueue {
			select {
			case carIdx := <-carQueue:
				// calculate car waiting time on cross
				res[carIdx] = event.time
			default:
				// no cars in queue
				emptyQueue = true
			}
		}
	}
	return res
}
