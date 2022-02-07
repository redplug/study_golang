# study_golang

노마드 쉽고 빠른 Go 시작하기 Study 진행 (무료)

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

## Account + NewAccount

```go
package main

import (
	"STUDY/banking"
	"fmt"
)

func main() {
	account := banking.Account{Owner: "jihwan", Balance: 1000}
	fmt.Println(account)
}

```

```
[Running] go run "c:\Users\blies-pc\go\study\main.go"
{jihwan 1000}

[Done] exited with code=0 in 2.678 seconds

```



## Methods

```go
//main.go
package main

import (
	"fmt"
	"go/STUDY/banking"
	"log"
)

func main() {
	account := banking.NewAccount("jihwan")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil { 
		log.Fatalln(err)
	}
	fmt.Println(account.Balance()) // 에러 출력
}

// banking/banking.go
package banking

import "errors"

type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	return nil
}

```

```
[Running] go run "c:\Users\blies-pc\go\study\main.go"
# go/STUDY/banking
banking\banking.go:33:1: syntax error: non-declaration statement outside function body

[Done] exited with code=2 in 0.779 seconds
```



## Finishing Up

```go
// main.go
package main

import (
	"fmt"
	"go/STUDY/banking"
)

func main() {
	account := banking.NewAccount("jihwan")
	account.Deposit(10)
	account.Withdraw(20)
	fmt.Println(account)
}

// banking/banking.go
package banking

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}

```

```
[Running] go run "c:\Users\blies-pc\go\study\main.go"
jihwan's account.
Has: 10

[Done] exited with code=0 in 1.298 seconds

```



## Dictionary

```go
// main.go
package main

import (
	"fmt"
	"go/STUDY/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	fmt.Println(dictionary["first"])
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}

// mydict/mydict.go
package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

```

```
C:\Users\blies-pc\go\study>go run main.go
First word
Not Found

C:\Users\blies-pc\go\study>
```



## Add Method

```go
// main.go
package main

import (
	"fmt"
	"go/STUDY/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}

// mydict/mydict.go
package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

```

```
C:\Users\blies-pc\go\study>go run main.go
found hello definition: Greeting
That word already exists

C:\Users\blies-pc\go\study>
```



## Update Delete

```
// main.go
package main

import (
	"fmt"
	"go/STUDY/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}

// mydict/mydict.go
package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-existing word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

```

```
C:\Users\blies-pc\go\study>go run main.go
Second

C:\Users\blies-pc\go\study>

baseWord 변경 후
C:\Users\blies-pc\go\study>go run main.go
Cant update non-existing word
First

C:\Users\blies-pc\go\study>
```



```go
// main.go
package main

import (
	"fmt"
	"go/STUDY/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

}

// mydict/mydict.go
package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-existing word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word delete
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

```

```
C:\Users\blies-pc\go\study>go run main.go
Not Found


C:\Users\blies-pc\go\study>go run main.go
```



# \#3 URL CHECKER & GO ROUTINES

## hitURL

```go
package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co",
	}
	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}

```

```
[Running] go run "c:\Users\blies-pc\go\study\main.go"
Checking: https://www.airbnb.com
Checking: https://www.google.com
Checking: https://www.amazon.com
Checking: https://www.reddit.com
Checking: https://soundcloud.com
Checking: https://www.facebook.com
Checking: https://www.instagram.com
Checking: https://academy.nomadcoders.co

[Done] exited with code=0 in 5.407 seconds
```



## Slow URLChecker

```go
package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co",
	}
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

```

```
[Running] go run "c:\Users\blies-pc\go\study\main.go"
Checking: https://www.airbnb.com
Checking: https://www.google.com
Checking: https://www.amazon.com
Checking: https://www.reddit.com
Checking: https://soundcloud.com
Checking: https://www.facebook.com
Checking: https://www.instagram.com
Checking: https://academy.nomadcoders.co
https://www.amazon.com OK
https://www.reddit.com OK
https://soundcloud.com OK
https://www.facebook.com OK
https://www.instagram.com OK
https://academy.nomadcoders.co OK
https://www.airbnb.com OK
https://www.google.com OK

[Done] exited with code=0 in 5.407 seconds

```



## Goroutines

일반적인 프로그램 방식

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	sexyCount("nico")
	sexyCount("jihwan")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}

}

```

```go
[Running] go run "c:\Users\blies-pc\go\study\main.go"
nico is sexy 0
nico is sexy 1
nico is sexy 2
nico is sexy 3
nico is sexy 4
nico is sexy 5
nico is sexy 6
nico is sexy 7
nico is sexy 8
nico is sexy 9
jihwan is sexy 0
jihwan is sexy 1
jihwan is sexy 2
jihwan is sexy 3
jihwan is sexy 4
jihwan is sexy 5
jihwan is sexy 6
jihwan is sexy 7
jihwan is sexy 8
jihwan is sexy 9

[Done] exited with code=0 in 24.184 seconds
```

좀 더 빠른 방식 Goroutines

Goroutines 은 메인 함수가 실행될대만 유지됨

-> 메인함수가 끝나면 종료됨, 메인함수가 종료되면 고루틴도 종료됨

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico")
	go sexyCount("jihwan")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}

}

```

```
[Running] go run "c:\Users\blies-pc\go\study\main.go"
jihwan is sexy 0
nico is sexy 0
nico is sexy 1
jihwan is sexy 1
jihwan is sexy 2
nico is sexy 2
jihwan is sexy 3
nico is sexy 3
nico is sexy 4
jihwan is sexy 4

[Done] exited with code=0 in 7.167 seconds
```

## Channels

goroutine와 메인함수상에 커뮤니케이션

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "jihwan"}
	for _, person := range people {
		go isSexy(person, c)
	}
	result := <-c
	fmt.Println(result)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}

```

```

C:\Users\blies-pc\go\study>go run main.go
jihwan
true

C:\Users\blies-pc\go\study>

```

## URLChecker + Go Routines

코드 미완성

## FAST URLChecker

```go
package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}

}

func hitURL(url string, c chan<- requestResult) {
	// fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}

```

```
[Running] go run "c:\Users\blies-pc\go\study\main.go"
https://www.instagram.com OK
https://www.airbnb.com OK
https://academy.nomadcoders.co OK
https://www.amazon.com OK
https://www.google.com OK
https://www.facebook.com OK
https://soundcloud.com OK
https://www.reddit.com OK

[Done] exited with code=0 in 2.646 seconds
```

순차 할경우 5초 현재 2초대

# \#4 JOB SCRAPPER

## getPages

## extractJob

## Writing Jobs

## Channels Time

## More Channels Baby

## Recap

# \#5 WEB SERVER WITH ECHO

## Setup

## File Download

## Conclusions
