// https://atcoder.jp/contests/agc031/tasks/agc031_a

package main

import "fmt"

func A() {
	var n int64
	fmt.Scan(&n)
	mp := map[string]int64{}

	str := ""
	fmt.Scan(&str)
	var i int64
	for i = 0; i < n; i++ {
		mp[string(str[i])]++
	}

	var ans int64
	ans = 1

	var mod int64 = 1000000007

	for _, val := range mp {
		ans = ((ans % mod) * ((val + 1) % mod)) % mod
	}
	fmt.Println(ans - 1)

}
