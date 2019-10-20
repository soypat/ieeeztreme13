package main

import (
	"bytes"
	"fmt"
)
const nmax = 302
//const nmax = 20 //For testing
var b bytes.Buffer
var n int
var isFibo [nmax]bool


func main() {

	fmt.Scanln(&n)
	fibo := 1
	isFibo[1] = true
	for i:=0;fibo<=n;i++ {
		fibo = fibo+i
		isFibo[fibo] = true
	}
	for i:=1;i<=n;i++ {
		if isFibo[i] {
			//b.WriteString("y")
			fmt.Print("y")
		} else {
			//b.WriteString("Y")
			fmt.Print("Y")
		}
	}
	//fmt.Printf("%x",b)
	fmt.Println("")
	fmt.Println(isFibo)
}

