// https://atcoder.jp/contests/abc052/tasks

package main

import (
    "fmt"
)


func main() {
   n:=0
    mp:=map[int]int{}
    fmt.Scan(&n)
    for val:=1;val<=n;val++{
        k:=val
        for i:=2;i*i<=k;i++{
            for k%i==0{
                k=k/i
                mp[i]++
            }
        }
        if k>1 {
            mp[k]++
        }
    }
    mod:=1000000007

    ans:=1
    for _,val:=range mp{
        ans=((ans%mod)*(val+1)%mod)%mod
    }
    fmt.Println(ans)
}