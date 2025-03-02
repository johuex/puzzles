package main

import (
	"sort"
	"strconv"
	"strings"
)

type Event struct {
	start  int
	end    int
	number int
}

func MaxEvent(events []string) []int {
	var n int = len(events) // count events
	var allEvents []Event
	var resultEvents []int //
	for i := 0; i < n; i++ {
		eventTime := strings.Split(events[i], ", ")
		eventStartTime, _ := strconv.Atoi(eventTime[0])
		eventEndTime, _ := strconv.Atoi(eventTime[1])
		event := Event{eventStartTime, eventEndTime, i}
		allEvents = append(allEvents, event)
	}

	// sort by ending of event, asc
	sort.Slice(
		allEvents,
		func(i, j int) bool {
			return allEvents[i].end < allEvents[j].end
		},
	)

	// find
	var lastEndTime int = 0
	for _, event := range allEvents {
		if event.start >= lastEndTime {
			lastEndTime = event.end
			resultEvents = append(resultEvents, event.number)
		}
	}

	return resultEvents
}
