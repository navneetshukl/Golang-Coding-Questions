// https://atcoder.jp/contests/abc172/tasks/abc172_d
package main

import "fmt"

func main() {
	var n int64
	fmt.Scanln(&n)
	var ans int64
	ans = n * (n + 1) / 2
	var i int64
	var j int64
	for i = 2; i <= n; i++ {
		for j = i; j <= n; j = j + i {
			ans += j
		}
	}
	fmt.Println(ans)

}



