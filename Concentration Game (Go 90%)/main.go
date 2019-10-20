package main

import "fmt"
const nmax = 6
const msg = "MATCH"
//var positions map[int]int
var prs1,prs2 bool
var iguess int
func main() {
	val2position := make(map[int]int,nmax)
	pos2values := make(map[int]int,nmax)
	var n, r1, r2,i2 int
	fmt.Scanln(&n)
	var n2 = n*2
	for i:=1;i<=n/2;i++ {
		i2 = i*2
		r1, r2 = guess(i2-1,i2)
		if r2==0 && r1 == 0 { // matchee 2 cartas
			// que hago?
		}
		_, prs1 = val2position[r1]
		_, prs2 = val2position[r2]
		if prs1 {
			_, _ = guess(i2-1,val2position[r1]) // adivino sabiendo donde está
			val2position[r1] = 0
		} else {val2position[r1] = i2-1}
		if prs2 {
			_, _ = guess(i2,val2position[r2]) // adivino sabiendo donde está
			val2position[r2] = 0
		} else {val2position[r2] = i2}
		pos2values[i2-1] = r1
		pos2values[i2] = r2
		if iguess == n2 {fmt.Println("-1");return}
	}
	fmt.Print("-1")
}

func guess(i int,j int) (int,int)  {
	var r1,r2 int
	fmt.Println(i,j)
	iguess++
	_,err := fmt.Scan(&r1,&r2)
	if err!=nil {return 0,0} else {
		return r1,r2
	}
}