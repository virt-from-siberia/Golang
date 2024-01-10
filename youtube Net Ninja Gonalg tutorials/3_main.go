package main

func main() {
	x := 0

	for x < 5 {
		println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		println(i)
	}

	names := []string{"Alex", "Tom", "Sam", "Mary"}

	for i := 0; i < len(names); i++ {
		println(names[i])
	}
}
