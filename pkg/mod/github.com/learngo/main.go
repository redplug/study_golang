package main

import "fmt"

func main() {
	// const 상수
	const name string = "redplug"
	fmt.Println(name)
	// var 변수
	var name2 string = "redplug2" // name2 := "redplug2" :=은 타입을 찾아줌
	fmt.Println(name2)
	// func 안에서는 안에서만 사용 지역변수, 밖에서는 전역
}
