package main

import "fmt"

func strings() {
	age := 38
	name := "Alex"

	fmt.Print("Hello!")
	fmt.Print("World! \n")

	fmt.Println("Hello Aleksey!")

	fmt.Println(name, "is", age, "years old")

	// PrintF
	fmt.Printf("%v is %v years old \n", name, age)
	fmt.Printf("%q is %q years old \n", name, age)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points. \n", 225.55)
	fmt.Printf("you scored %0.2f points. \n", 225.55)

	// SprintF
	var str string = fmt.Sprintf("%v is %v years old \n", name, age)
	fmt.Println("the saved string is:", str)
}

func arraysAndSlice() {
	// var ages [3]int =

	var ages = [3]int{20, 25, 30}

	names := [4]string{"Alex", "Tom", "Sam", "Mary"}
	names[1] = "Bill"

	fmt.Println(ages, len(ages), cap(ages))
	fmt.Println(names, len(names), cap(names))

	var scores = []int{100, 90, 95}
	newScores := append(scores, 85)

	fmt.Println(newScores)

	// slice ranges

	rangeOne := names[1:3]
	fmt.Println("slice ranges")
	fmt.Println(rangeOne)

}

func main() {
	// strings()
	arraysAndSlice()
}
