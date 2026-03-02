package main

func countEven(num int) int {
	var cnt int
	for i := 1; i <= num; i++ {
		if evenNum(i) {
			cnt += 1
		}
	}

	return cnt
}

func evenNum(num int) bool {
	var sum int
	for num > 0 {
		tmp := num % 10
		sum += tmp
		num = num / 10
	}
	return sum%2 == 0
}
