package main

import (
	"fmt"
	"time"
)

func main() {

	odd := []int{1, 3, 5, 7, 9}
	even := []int{0, 2, 4, 6, 8}

	// Anonymous function
	go func() {
		for i := range odd {
			fmt.Printf("Print Odd: %d\n", odd[i])
		}
	}()

	// Anonymous function
	go func() {
		for i := range even {
			fmt.Printf("Print Even: %d\n", even[i])
		}
	}()

	<-time.After(time.Second * 5)
}

/* OUTPUT
// Output always changes
Print Odd: 1
Print Even: 0
Print Odd: 3
Print Odd: 5
Print Odd: 7
Print Odd: 9
Print Even: 2
Print Even: 4
Print Even: 6
Print Even: 8
*/
