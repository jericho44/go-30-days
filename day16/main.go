
package main

import "fmt"

type Account struct{ Balance int }
func (a *Account) Deposit(n int) { a.Balance += n }

func main() {
	acc := &Account{100}
	acc.Deposit(50)
	fmt.Println(acc.Balance)
}
