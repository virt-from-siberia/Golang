package main

import (
	"fmt"
	"math"
)

func satHello(n string) {
	fmt.Printf("good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}

}

func circleAria(r float64) float64 {
	return math.Pi * r * r

}

func main() {
	satHello("hello world")
	sayBye("Aleskey")

	cycleNames([]string{"Aleksey", "Ivan", "Tatyana"}, satHello)
	cycleNames([]string{"Aleksey", "Ivan", "Tatyana"}, sayBye)

	a1 := circleAria(10.5)
	a2 := circleAria(15)

	fmt.Println(a1, a2)
}
