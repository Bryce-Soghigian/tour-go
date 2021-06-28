package main

import "fmt"


func main(){
	var sum int64
	var i int64
	for i =0; sum < 10;i++{
		sum += i
	}

	fmt.Println(sum)
}