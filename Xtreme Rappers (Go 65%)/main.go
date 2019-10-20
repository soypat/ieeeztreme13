package main

import "fmt"

func main() {
	var k, j int
	fmt.Scanln(&k, &j)
	M := max(k,j)
	m := min(k,j)
	if m*2 <= M {fmt.Print(m);return}
	if k<j && j>2*(k) {fmt.Print(k/3);return}
    if k>j && k>2*(j) {fmt.Print(j/3);return}
	if k==j {
		fmt.Print(k/3);return
	} else {
		diff := M - m
		if diff <= m-diff {fmt.Print(diff+m/3);return}
	}
	fmt.Println(2*(m/3))
}

func max(a int,b int) int {
	if a>b{return a}else{return b}
}

func min(a int,b int) int {
	if a<b{return a}else{return b}
}