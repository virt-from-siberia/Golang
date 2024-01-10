package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "hello my dear friends"

	// fmt.Println(strings.Contains(greeting, "world!"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.ToLower(greeting))

	// fmt.Println(strings.Split(greeting, " "))

	ages := []int{20, 25, 30, 32, 54, 67, 78}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 32)
	fmt.Println(index)

	names := []string{"Alex", "Tom", "Sam", "Mary"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Sam"))

}
