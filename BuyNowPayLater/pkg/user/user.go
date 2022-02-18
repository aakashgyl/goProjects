package user

import "errors"

var CREDIT_LIMIT_HIT = errors.New("rejected! (reason: credit limit)")
var EXCESS_PAYBACK_ERROR = errors.New("payback is more than credit usage")

type UserOps interface {
	GetName() string
	Purchase(float32) error
	Payback(float32) error
	GetRemainingCredit() float32
	GetCreditLimit() float32
}

type User struct {
	Name string
	Email string
	CreditLimit float32
	CreditUsed float32
}

func GetNewUser(name, email string, creditLimit float32) UserOps {
	return &User{
		Name:        name,
		Email:       email,
		CreditLimit: creditLimit,
		CreditUsed:  0,
	}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) Purchase(amount float32) error {
	if u.GetRemainingCredit() < amount {
		return CREDIT_LIMIT_HIT
	}
	u.CreditUsed = u.CreditUsed + amount
	return nil
}

func (u *User) Payback(amount float32) error {
	if amount > u.CreditUsed {
		return EXCESS_PAYBACK_ERROR
	}
	u.CreditUsed = u.CreditUsed - amount
	return nil
}

func (u *User) GetRemainingCredit() float32 {
	return u.CreditLimit - u.CreditUsed
}

func (u *User) GetCreditLimit() float32 {
	return u.CreditLimit
}
