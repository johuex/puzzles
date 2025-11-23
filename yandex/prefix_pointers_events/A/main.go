package a

import (
	"bufio"
	"os"
	"sort"
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
	lineStr := strings.TrimSpace(line)
	n, _ := strconv.Atoi(lineStr)

	firstDirection := make([]string, n)
	for i := range n {
		line, _ = reader.ReadString('\n')
		lineStr = strings.TrimSpace(line)
		firstDirection[i] = lineStr
	}

	line, _ = reader.ReadString('\n')
	lineStr = strings.TrimSpace(line)
	m, _ := strconv.Atoi(lineStr)

	secondDirection := make([]string, m)
	for i := range m {
		line, _ = reader.ReadString('\n')
		lineStr = strings.TrimSpace(line)
		secondDirection[i] = lineStr
	}

	// logic
	ans := schedule(firstDirection, secondDirection)

	// output
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func createEvents(firstDirection, secondDirection []string) [][]int {
	events := make([][]int, (len(firstDirection)+len(secondDirection))*2)
	eIdx := 0

	// type: 1- -> arrival, 2 -> departure
	// point:  0 -> 1'st office or 1 -> 2'nd office
	// event: (minutes, type, point)
	for _, event := range firstDirection {
		parts := strings.Split(event, "-")
		startParts := strings.Split(parts[0], ":")
		startHours, _ := strconv.Atoi(startParts[0])
		startMinutes, _ := strconv.Atoi(startParts[1])
		events[eIdx] = []int{startHours*60 + startMinutes, 2, 0}
		eIdx += 1

		endParts := strings.Split(parts[1], ":")
		endHours, _ := strconv.Atoi(endParts[0])
		endMinutes, _ := strconv.Atoi(endParts[1])
		events[eIdx] = []int{endHours*60 + endMinutes, 1, 1}
		eIdx += 1
	}

	for _, event := range secondDirection {
		parts := strings.Split(event, "-")
		startParts := strings.Split(parts[0], ":")
		startHours, _ := strconv.Atoi(startParts[0])
		startMinutes, _ := strconv.Atoi(startParts[1])
		events[eIdx] = []int{startHours*60 + startMinutes, 2, 1}
		eIdx += 1

		endParts := strings.Split(parts[1], ":")
		endHours, _ := strconv.Atoi(endParts[0])
		endMinutes, _ := strconv.Atoi(endParts[1])
		events[eIdx] = []int{endHours*60 + endMinutes, 1, 0}
		eIdx += 1
	}

	sort.Slice(events, func(i, j int) bool {
		// sort by time and then by event time (arrival earlier than departure)
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	return events
}

func schedule(firstDirection, secondDirection []string) int {
	// type: 1- -> arrival, 2 -> departure
	// point: 0 or 1
	// event: (minutes, type, point)
	events := createEvents(firstDirection, secondDirection)

	busesCount := make([]int, 2)
	minBusesCount := make([]int, 2)

	for i := 0; i < len(events); i++ {
		event := events[i]
		switch event[1] {
		case 1:
			// bus arrived to city
			busesCount[event[2]] += 1
		case 2:
			// bus departed from city
			busesCount[event[2]] -= 1
			minBusesCount[event[2]] = min(minBusesCount[event[2]], busesCount[event[2]])
		}
	}

	// return sum of negative numbers because this buses missed on station but started their trip
	ans := 0
	for i := 0; i < len(busesCount); i++ {
		if minBusesCount[i] < 0 {
			ans -= minBusesCount[i]
		}
	}
	return ans
}
