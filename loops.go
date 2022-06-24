//need to package our code.
package main

import "fmt"

//format package for printing.

//main function to tell the beginning of the code.
func main() {
	i := 1
	for i <= 5 {
		fmt.Println("webhav", i)
		i++
	}
	//Same thing can be one in a single line-
	for i := 1; i <= 5; i++ {
		fmt.Println("webhav", i)
	}
