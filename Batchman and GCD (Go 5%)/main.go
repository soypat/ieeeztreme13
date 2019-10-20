package main

import "fmt"
const nmax = 100100 // Constraint
//const nmax = 7



func main() {
	 mygcds := make(map[int]bool,nmax)
	var A [nmax]int
	var n, k, r int
	fmt.Scanln(&n, &k)
	for i:=1;i<=n;i++ {
		fmt.Scan(&A[i])
	}
	//fmt.Println(A)
	for i:=1;i<=n;i++ {
		for j:=2;j<=n;j++ {
			_,presente := mygcds[gcd(A[i],A[j])]
			if presente {continue}
			mygcds[gcd(A[i],A[j])] = true
			r++
		}
	}
	fmt.Println(r)
	return
	fmt.Println(mygcds)
}

func gcd(x int,y int) int {
	var aux = 0
	for y!=0 {
		aux =y
		y = x%y
		x = aux
	}
	return x
}