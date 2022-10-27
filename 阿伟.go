package main

import (
	"fmt"
	"time"
)

func main() {
	B := A(打电动)
	fmt.Println(B)
	C := A(欢迎来我家玩)
	fmt.Println(C)
}

type G func() string

func A(fType G) (B string) {
	B = fType()
	return
}

func 打电动() string {
	return "输了啦，都是你害的"
}
func 欢迎来我家玩() string {
	// 花费 5s 前往杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郎"
}
