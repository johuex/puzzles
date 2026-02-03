package concpatterns

import "fmt"

func Generator(num int) <-chan int {
	// generator that return all squares from 0 to num*num
	ch := make(chan int)

	go func() {
		for i := 1; i <= num; i++ {
			ch <- i * i
		}
		close(ch)
	}()

	return ch
}

func GeneratorRead() {
	for num := range Generator(100) {
		fmt.Println(num)
	}
}
