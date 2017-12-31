package main

import (
	"As/entities"
	"fmt"
	"time"
)


func main()  {
	var n int
	fmt.Println("1 for  Naive Approach \n2 for Optimized Approach\n3 Stop")
	for {
		fmt.Scanln(&n)
		if (n == 3) {
			break
		} else if n == 1{
			t1 := time.Now()
			entities.Naive()
			elapsed := time.Since(t1)
			fmt.Println("Time: ", elapsed)

		} else {
			t1 := time.Now()
			entities.Optimized()
			elapsed := time.Since(t1)
			fmt.Println("Time: ", elapsed)
		}
	}
}
