package main

import (
	"fmt"
)

func main() {
	var N, a , b int
	fmt.Scan(&N,&a,&b)
	n1 := N/a
	if !(n1*a <= N && N <= n1*b)  {fmt.Print("NO");return} else{fmt.Println("YES")}

	var ni int

	for ni=n1; (ni*a <= N && N <= ni*b) && ni>0; ni-- {

	}
	ni ++
	if (b-a)==0 {
		for i := 1;i<=ni;i++ {
			fmt.Print(a," ")
		}
		fmt.Println("")
		return
	}
	v := (ni*b - N) / (b-a)
	rem := (ni*b - N) % (b-a)
	for i := 1;i<=v;i++ {
		fmt.Print(a," ")
	}
	if !(rem==0) {
		fmt.Print(a+rem," ")
	}
	for i := 1;i<=ni-v;i++ {
		if !(rem==0) && i==1 {
			continue
		}
		fmt.Print(b," ")
	}
	fmt.Println("")
}

