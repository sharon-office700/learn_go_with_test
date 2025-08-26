package integers

import "fmt"

// Add takes two integers and returns the sum of their values
func Add(a, b int) int {
	return a + b
}


func ExampleAdd(a, b int) {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}