package main

import "fmt"

func main() {
	//While loop using For
	i := 1
	for i <= 10 {
		fmt.Print(i)
		i++
	}
	//Classic for Loop
	// for j := 1; j <= 10; j++ {
	// 	if j == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(j)
	// }

	//Range Keyword
	for i := range 3 {
		fmt.Println(i)
	}
}
