# study_golang

노마드 쉽고 빠른 Go 시작하기 Study 진행

Url : https://nomadcoders.co/go-for-beginners

# #0 INTRODUCTION

Software and Installation

- https://repl.it 에 들어가서 진행
- 컴퓨터에 설치 하여 진행
  - VS Code를 사용 Go extension
  - https://golang.org 에서 go 다운로드
    https://go.dev/dl/
  - 고 코드는 고 패스가 설치된 폴더에서만 실행 가능?

# #1 THEORY

- 컴파일 하려면 main.go로 시작

- main.go가 없는 경우 컴파일을 하지 않으려고 만든것

- 

- hello world

- ```go
  // main.go
  package main
  
  import "fmt"
  
  func main() {
  	fmt.Println("Hello world!!")
  }
  ```

- ![image-20220202195713152](https://raw.githubusercontent.com/redplug/shareimages/master/img/image-20220202195713152.png)

## Packages and imports

- 임포트 하는 방법 import  "000"
- export하기 위해서는 대문자로 써야 함.

```go
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
```

Go언어는 타입이 여러가지 있음

string, bool, 숫자타입 여러개

```go
package main

import "fmt"

func multiply(a, b) {
	return a * b
}

func main() {
	fmt.println(multiply(2, 2))
}

```

```
상기와 같이 실행 시 오류가 남
[Running] go run "d:\99_Delete\golang\study_golang\pkg\mod\github.com\learngo\main.go"
# command-line-arguments
.\main.go:5:15: undefined: a
.\main.go:5:18: undefined: b
.\main.go:6:2: too many arguments to return
.\main.go:6:9: undefined: a
.\main.go:6:13: undefined: b
.\main.go:10:2: cannot refer to unexported name fmt.println
.\main.go:10:22: multiply(2, 2) used as value

[Done] exited with code=2 in 1.269 seconds
```

GO의 경우 사전에 변수 타입을 설정해줘야 함, 리턴 값 포함

```go
package main

import "fmt"

func multiply(a int, b int) int {   // a, b int 라고 표기 할 경우 둘다 int로 인식
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
}

```

기입 후 정상적으로 답이 나옴

```
[Running] go run "d:\99_Delete\golang\study_golang\pkg\mod\github.com\learngo\main.go"
4

[Done] exited with code=0 in 4.051 seconds
```

Go의 경우 리턴 값을 여러개 가져갈 수 있음

```go
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLenght, upperName := lenAndUpper("nico")
	fmt.Println(totalLenght, upperName)
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
4 NICO

[Done] exited with code=0 in 2.542 seconds
```



```go
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLenght, _ := lenAndUpper("nico") // _ 를 넣을 경우 값을 무시할 수 있음
	fmt.Println(totalLenght, upperName)
}

```

```go
package main

import (
	"fmt"
)

func repeatMe(words ...string) { // 점을 3개 찍어서 여러가지 변수 받는것으로 설정
	fmt.Println(words)
}

func main() {
	repeatMe("ji", "hwan", "red", "blue")
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
[ji hwan red blue]

[Done] exited with code=0 in 2.134 seconds
```

## 루프

- for만 가지고 돌림

```go
package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
1
2
3
4
5
6

[Done] exited with code=0 in 1.796 seconds

```

range 는 for문 안에서만 사용 가능

```go
package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // _ 자리는 index 자리
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
21

[Done] exited with code=0 in 2.095 seconds
```

defer : 펑션이 끝나고 추가적으로 하고 싶은 일을 명시

```go
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return // naked return
}
func main() {
	totalLenght, up := lenAndUpper("jihwan")
	fmt.Println(totalLenght, up)
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
I'm done
6 JIHWAN

[Done] exited with code=0 in 1.799 seconds
```

naked return : 함수에 리턴 할 값을 미리적어주고 뒤에는 return 만 적어줌

## if ,else

```go
package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // if문을 위한 koreanAge라는 변수를 사용할 수 있음
		return false
	} // 이곳에 else를 사용할 수 있지만 굳이 사용할 필요는 없음
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}

```

```

[Running] go run "d:\99_Delete\golang\study_golang\main.go"
true

[Done] exited with code=0 in 2.843 seconds

```

## switch

```go
package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(18))
}
```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
true

[Done] exited with code=0 in 1.922 seconds
```

```go
package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrink(18))
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
true

[Done] exited with code=0 in 1.954 seconds
```



## Pointers

```go
package main

import (
	"fmt"
)

func main() {
	a := 2
	b := a
	fmt.Println(&a, &b) // 앞에 &값을 주면 메모리 주소 값을 볼 수 있음
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
0xc000012098 0xc0000120b0

[Done] exited with code=0 in 2.004 seconds
```

```go
package main

import (
	"fmt"
)

func main() {
	a := 2
	b := &a // 메모리 주소 값을 저장
	fmt.Println(&a, b) // 출력 시 메모리 저장 값이 나옴
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
0xc000012098 0xc000012098

[Done] exited with code=0 in 1.516 seconds
```

```go
package main

import (
	"fmt"
)

func main() {
	a := 2
	b := &a
	fmt.Println(a, *b) // *를 적을 경우 해당 메모리 주소값에 대한 실제 값을 가져옴
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
2 2

[Done] exited with code=0 in 2.045 seconds
```

```go
package main

import (
	"fmt"
)

func main() {
	a := 2
	b := &a
	*b = 20 // a의 메모리 주소값을 가지고 있는 b를 이용해서 a값을 바꿀 수 있음
	fmt.Println(a)
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
20

[Done] exited with code=0 in 2.512 seconds
```

## Arrays and Slices

- 다른 프로그램과 array가 좀 다름
- slice : 기본적으로 array인데 길이가 없음

```go
package main

import (
	"fmt"
)

func main() {
	names := [5]string{"a", "b", "c"} // 배열 생성 방법
	names[3] = "d" // 추가 방법
	fmt.Println(names) // 프린팅 할경우 전체 다 나옴
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
[a b c d ]

[Done] exited with code=0 in 1.685 seconds
```

```go
package main

import (
	"fmt"
)

func main() {
	names := []string{"a", "b", "c"} // slice
	names = append(names, "d")       // append 자체는 변수를 수정해주지 않음.
	fmt.Println(names)

}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
[a b c d]

[Done] exited with code=0 in 2.413 seconds

```

## maps

- 파이썬의 오브젝트랑 비슷함, key, value 형태

```go
package main

import "fmt"

func main() {
	red := map[string]string{"name": "jihwan", "age": "12"}
	fmt.Println(red)
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
map[age:12 name:jihwan]

[Done] exited with code=0 in 2.133 seconds

```

```go
package main

import "fmt"

func main() {
	red := map[string]string{"name": "jihwan", "age": "12"}
	fmt.Println(red)
	for key, value := range red { // for문 작성 가능
		fmt.Println(key, value)
	}
}

```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
map[age:12 name:jihwan]
age 12
name jihwan

[Done] exited with code=0 in 2.231 seconds
```

## Structs

- 맵보다 유연함
- 어떤 struct인지 사전에 타입 정의하고 형태 정의

```go
package main

import "fmt"

type person struct { // 타입 정의, 형태 정의
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"} // 배열
	jihwan := person{name: "jihwan", age: 18, favFood: favFood} // 형태를 명시해주는게 좋음, 방식은 하나로 통일해야 함.
	fmt.Println(jihwan)
}


```

```
[Running] go run "d:\99_Delete\golang\study_golang\main.go"
{jihwan 18 [kimchi ramen]}

[Done] exited with code=0 in 2.502 seconds
```

# \#2 BANK & DICTIONARY PROJECTS



# \#3 URL CHECKER & GO ROUTINES



# \#4 JOB SCRAPPER



# \#5 WEB SERVER WITH ECHO
