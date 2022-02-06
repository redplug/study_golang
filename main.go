package main

import (
"STUDY_GOLANG/banking"
"fmt"
)

func main() {
account1 := banking.Account{Owner: "jihwan", Balance: 1000}
fmt.Println(account1)
}