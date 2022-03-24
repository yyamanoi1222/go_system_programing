package main

import "time"
import "fmt"

func main() {
	c := time.After(time.Second * 5)
	<-c

	fmt.Println("5s")
}
