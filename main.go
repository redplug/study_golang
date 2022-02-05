package main

import "fmt"

type person struct { // 타입 정의, 형태 정의
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"} // 배열
	jihwan := person{name: "jihwan", age: 18, favFood: favFood}
	fmt.Println(jihwan)
}
