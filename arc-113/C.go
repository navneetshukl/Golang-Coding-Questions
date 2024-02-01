//https://atcoder.jp/contests/arc113/tasks/arc113_c

package main

import "fmt"

func C() {
	s := ""
	fmt.Scan(&s)
	var ans int64
	n := len(s)
	ans = 0

	mp := map[string]int{}
	mp[string(s[n-1])]++

	for i := n - 2; i >= 1; i-- {
		if string(s[i]) == string(s[i-1]) {
			x := n - i - 1 - mp[string(s[i])]
			ans += int64(x)

			for i := 'a'; i <= 'z'; i++ {
				mp[string(i)] = 0
			}
			mp[string(s[i])] = n - i

		} else {
			mp[string(s[i])]++
		}
	}
	fmt.Println(ans)
}
