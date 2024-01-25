// https://atcoder.jp/contests/abc170/tasks/abc170_d

package main

import(
    "fmt")

func main() {
    var n int
    fmt.Scan(&n)
    mp:=map[int]int{}
    m:=[]int{}
    for i:=0;i<n;i++{
         val:=0
         fmt.Scan(&val)
         m=append(m,val)

    }
    for _,val:=range m{
        mp[val]++
    }
    ans:=0
    for i:=0;i<n;i++{
        if mp[m[i]]>1{
            continue;
        }
        k:=m[i]
        c:=0
        for j:=1;j*j<=k;j++{
            if k%j==0 && j!=m[i]{
                if mp[j]!=0{
                    c++
                    break
                }
                if (k/j)!=j && (k/j)!=m[i]{
                    va:=k/j
                    if mp[va]!=0{
                        c++
                        break;
                    }
                }
                
            }
        }
        if c==0{
           // fmt.Println("Value is ",m[i])
            ans++
        }
    }
    fmt.Println(ans)
    
}
