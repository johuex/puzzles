package main

import (
	"sort"
	"strconv"
	"strings"
)

type Task struct {
	name     string
	duration int
	deadline int
}

func MaxProfit(tasks []string) int {
	var n, profit int
	n = len(tasks) // count events
	profit = 0
	var allTasks []Task
	for i := 0; i < n; i++ {
		eventTime := strings.Split(tasks[i], ", ")
		taskName := eventTime[0]
		taskDuration, _ := strconv.Atoi(eventTime[1])
		taskDeadline, _ := strconv.Atoi(eventTime[2])
		task := Task{taskName, taskDuration, taskDeadline}
		allTasks = append(allTasks, task)
	}

	// sort by ending of event, asc
	sort.Slice(
		allTasks,
		func(i, j int) bool {
			return allTasks[i].duration < allTasks[j].duration
		},
	)

	// calculate
	var currentTime int = 0
	for _, task := range allTasks {
		currentTime += task.duration
		profit += task.deadline - currentTime
	}

	return profit

}
