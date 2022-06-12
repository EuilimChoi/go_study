package accounts

import "errors"

// Account struct

type Account struct { // 첫 글자가 대문자여야 export된다!!
	owner   string // 얘네 들도 대문자여야 public 하다
	balance int
}

// 만약 위처럼 모두 퍼블릭하게 내보낸다면 불러오는 곳에서 마음대로 바꿀수 있게된다. 하지만 안에 내용을 private하고 import하는 곳에서 사용만하게 하려면 아래와 같이 한다.

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount of your account
func (a *Account) Deposit(amount int) { //자바스크립트로 따지면 메서드 같은건데 고에서는 receiver라고 부른다 a는 Account 타입을 같는 리시버라고 보면된다. *리시버는 첫글자가 소문자여야한다.
	a.balance += amount
} // 리시버는 기본적으로 메서드를 실행할 사본을 만든다. 따라서 처음에는 a Account로 작성하면 원래 메서드의 사본에다가 메서드를 실행하고 아래의 Balance를 조회하면 사본이 아닌 원본에 조회를 해버린다.
// 따라서 a *Account라고 선언해 이 메서드가 원래 Account에 사용되도록 해줘야 한다.

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

var errNomoney = errors.New("Can't withdraw") //에러 정의

// withdraw x from your account
// 에러 잡는 연습 만약 출금액이 잔액보다 크면 애러를 뱉도록 해보자
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		//return errors.New("Can't withdraw") => 이렇게 써도된다 하지만 코드가 지져분해지는 느낌.. 아예 에러를 위에 정의하자
		return errNomoney
	}

	a.balance -= amount
	return nil
}

// chansgeOwner of the Account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// ReturnOwner
func (a *Account) ReturnOwner() string {
	return a.owner
}
