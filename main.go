package main

import (
	"time"
)

//场景切换 context switching

//func Hello(n int) (r int) {
//
//	defer func() {
//		fmt.Println(r)
//		r += n
//	}()
//	return n + n // <==> r = n+n
//}

func main() {
	var a = 123

	a = 789

	time.Sleep(time.Second * 2)

}
