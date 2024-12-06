package main

import (
	"fmt"
	"sync"
)

type humanInfo struct {
	name string
	age  int
	town string
}

func main() {
	// Просто запусти программу
	// 1 task
	echo()

	// 2 task
	vars()

	// 3 task
	isEven(5)

	// 4 task
	cycle()

	// 5 task
	add(3, 5)

	// 6 task
	sum([]int{2, 3, 4, 5, 6})

	// 7 task
	fmt.Println("task 7")
	fmt.Println(humanInfo{"Eddy", 35, "Ney-york"})
	fmt.Println()

	// 8 task
	a := 9
	incr(&a)

	// 9 task
	div(8, 4)

	// 10 task
	wg := sync.WaitGroup{}
	wg.Add(3)
	goChans(&wg)
	wg.Wait()
}

func echo() {
	fmt.Println("task 1")
	fmt.Println("Hello, World!")
	fmt.Println()
}

func vars() {
	fmt.Println("task 2")
	a := 5
	b := "this is string"
	c := true
	fmt.Printf("integer:%d, string:%s, bool:%v", a, b, c)
	fmt.Print("\n\n")
}

func isEven(a int) {
	fmt.Println("task 3")
	if a%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("not even")
	}
	fmt.Println()
}

func cycle() {
	fmt.Println("task 4")
	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}
	fmt.Println()
}

func add(a, b int) {
	fmt.Println("task 5")
	fmt.Println(a + b)
	fmt.Println()
}

func sum(sl []int) {
	fmt.Println("task 6")
	var sum int
	for _, v := range sl {
		sum += v
	}
	fmt.Println(sum)
	fmt.Println()
}

func incr(a *int) {
	fmt.Println("task 8")
	*a++
	fmt.Println(*a)
	fmt.Println()
}

func div(a, b int) {
	fmt.Println("task 9")
	if b == 0 {
		fmt.Println("divide by zero is illegal")
		fmt.Println()
	} else {
		fmt.Println(a / b)
		fmt.Println()
	}
}

func goChans(wg *sync.WaitGroup) {
	fmt.Println("task 10")
	a := make(chan int)
	b := make(chan int)
	go func() {
		defer wg.Done()
		defer close(a)
		for i := 1; i < 6; i++ {
			a <- i
		}
	}()
	go func() {
		defer wg.Done()
		defer close(b)
		for {
			val, ok := <-a
			if !ok {
				break
			}
			b <- val * val
		}
	}()
	go func() {
		defer wg.Done()
		for {
			val, ok := <-b
			if !ok {
				break
			}
			fmt.Println(val)
		}
	}()
}
