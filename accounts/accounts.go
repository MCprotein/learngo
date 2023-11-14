package accounts

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

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
