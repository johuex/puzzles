package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type CustomerActionType struct {
	value  int // time of action
	action int // +1 (enter) or -1 (exit)
}

func MaxCustomers(customers []string) int {
	var n, max, count int

	n = len(customers) // count of customers
	var customerActions []CustomerActionType
	for i := 0; i < n; i++ {
		customerTimes := strings.Split(customers[i], ", ")
		enterCustomerTime, _ := strconv.Atoi(customerTimes[0])
		endCustomerTime, _ := strconv.Atoi(customerTimes[1])
		enterAction := CustomerActionType{enterCustomerTime, 1}
		exitAction := CustomerActionType{endCustomerTime, -1}
		customerActions = append(customerActions, enterAction, exitAction)
	}

	// sort through customer action time, asc
	sort.Slice(
		customerActions,
		func(i, j int) bool {
			return customerActions[i].value < customerActions[j].value
		},
	)
	for i := 0; i < n*2; i++ {
		count += customerActions[i].action
		max = int(math.Max(float64(max), float64(count)))
	}
	return max

}
