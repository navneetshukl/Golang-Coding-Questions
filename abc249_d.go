// https://atcoder.jp/contests/abc249/tasks/abc249_d

package main

import "fmt"

func Div(n int64)[]int64{
    arr:=[]int64{}
    var i int64
    for i=1;i*i<=n;i++{
        if n%i==0{
            arr=append(arr,i)
            if n/i!=i{
                arr=append(arr,n/i);
            }
        }
    }
    return arr
}
func main() {
	var n int64
	fmt.Scan(&n)
	mp:=map[int64]int64{}
	arr:=[]int64{}
	var i int64
	for i=0;i<n;i++{
	    var val int64
	    fmt.Scan(&val)
	    arr=append(arr,val)
	    mp[val]++
	}
	var ans int64=0
	for i=0;i<n;i++{
	    ar:=Div(arr[i])
	    
	    for _,val:=range ar{
	        ans+=mp[val]*mp[arr[i]/val]
	    }
	}
	fmt.Println(ans)
}
