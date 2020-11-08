package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	defer fmt.Println("hello, world")
	defer fmt.Println("the time is", time.Now())
	fmt.Println("my favorite number is", rand.Intn(10))
	fmt.Println(add(12, 23))
	fmt.Println(swap("あんた", "ばかぁ？"))
	const test = "test"
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	fmt.Println(isOK(true))

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func isOK(x bool) string {
	if x {
		return "OK"
	} else {
		return "NO"
	}
}
