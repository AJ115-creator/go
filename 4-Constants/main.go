package main

import "fmt"

func main() {
	const name = "hi"

	//Grouping Constants
	const (
		host = "localhost"
		port = 8080
	)
	fmt.Println(port, host)
}
