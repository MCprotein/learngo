package accounts

import "errors"

/*
Account struct
struct는 생성자가 없다.

다른 파일에서는
privat: 소문자
public: 대문자
*/
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw you are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

/*
Deposit x amount on your account
function이 아니라 method이다.
(account Account) 는 receiver라고 부른다.
receiver의 type을 그대로 쓰면 Account의 복사본을 만들어서 사용한다.
receiver의 type을 *Account로 쓰면 Account의 복사본을 만들지 않고 이 메소드를 호출한 account를 사용한다.
*/
func (account *Account) Deposit(amount int) {
	account.balance += amount
}

func (account Account) Balance() int {
	return account.balance
}

// Withdraw x amount from your account
func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return errNoMoney
	}

	account.balance -= amount
	return nil
}
