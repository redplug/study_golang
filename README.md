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





# \#2 BANK & DICTIONARY PROJECTS



# \#3 URL CHECKER & GO ROUTINES



# \#4 JOB SCRAPPER



# \#5 WEB SERVER WITH ECHO
