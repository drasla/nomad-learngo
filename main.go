package main

import (
	"fmt"
	accounts "nomad-learngo/banking"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
