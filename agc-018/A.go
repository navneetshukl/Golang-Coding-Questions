// https://atcoder.jp/contests/agc018/tasks/agc018_a

package main

import (
	"fmt"
	"sort"
)

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a == b {
		return a
	}

	if a > b {
		return gcd(a-b, b)
	}
	return gcd(a, b-a)
}

func A() {
	n, k := 0, 0

	fmt.Scan(&n, &k)

	arr := []int{}
	for i := 0; i < n; i++ {
		val := 0
		fmt.Scan(&val)

		arr = append(arr, val)
	}
	sort.Ints(arr)

	if arr[n-1] < k {
		fmt.Println("IMPOSSIBLE")
		return
	} else {
		for i := 0; i < n; i++ {
			if arr[i] == k {
				fmt.Println("POSSIBLE")
				return
			}
		}
		g := 0
		for i := 1; i < n; i++ {
			g = gcd(g, arr[i]-arr[i-1])
		}
		if k%g == 0 {
			fmt.Println("POSSIBLE")
		} else {
			fmt.Println("IMPOSSIBLE")
		}

	}
}
