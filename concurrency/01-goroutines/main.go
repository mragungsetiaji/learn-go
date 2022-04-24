package main

import (
	"fmt"
	"time"
)

func main() {
	odd := []int{1, 3, 5, 7, 9}
	even := []int{0, 2, 4, 6, 8}
	go listOdd(odd)
	go listEven(even)

	<-time.After(time.Second * 10)
}

func listOdd(items []int) {
	for i := range items {
		fmt.Printf("Print Odd: %d\n", items[i])
	}
}

func listEven(items []int) {
	for i := range items {
		fmt.Printf("Print Even: %d\n", items[i])
	}
}

/* OUTPUT
// Output always changes
Print Even: 0
Print Odd: 1
Print Even: 2
Print Even: 4
Print Even: 6
Print Even: 8
Print Odd: 3
Print Odd: 5
Print Odd: 7
Print Odd: 9
*/
