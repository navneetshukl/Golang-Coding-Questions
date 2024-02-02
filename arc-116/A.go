// https://atcoder.jp/contests/arc116/tasks/arc116_a

package main

import "fmt"

func A() {

	t := 0
	fmt.Scan(&t)
	for t > 0 {
		var n int64
		fmt.Scan(&n)
		if (n%2) == 0 && (n/2)%2 == 1 {
			fmt.Println("Same")
		} else if n%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
		t--
	}

}
