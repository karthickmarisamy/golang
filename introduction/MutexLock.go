package main
import (
	"fmt"
	"time"
	"sync"
)
type Account struct{
	balance int
	lock sync.Mutex
}

func (a *Account) GetBalance()int{
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) WithdrawAmount(v int) int{
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance{
		fmt.Println("Sry! your balance is not enough to withdraw")
	} else{
		a.balance -= v
		fmt.Println("Current Balance\n", a.balance)
	}
	
	return a.balance
}

func main(){
	var acc Account
	acc.balance = 100
	fmt.Println("Current Balance\n", acc.GetBalance())
	for i := 0; i < 12 ; i ++{
		go acc.WithdrawAmount(10)
	}
	time.Sleep( 2 * time.Second)
}