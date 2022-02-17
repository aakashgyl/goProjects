package user

import "errors"

type UserOps interface {
	Purchase(float32) error
	Payback(float32)
	GetRemainingCredit() float32
}

type user struct {
	Name string
	Email string
	CreditLimit float32
	CreditUsed float32
}

func GetNewUser(name, email string, creditLimit float32) UserOps {
	return &user{
		Name:        name,
		Email:       email,
		CreditLimit: creditLimit,
		CreditUsed:  0,
	}
}

func (u *user) Purchase(amount float32) error {
	if u.GetRemainingCredit() < amount {
		return errors.New("rejected! (reason: credit limit)")
	}
	u.CreditUsed = u.CreditUsed + amount
	return nil
}

func (u *user) Payback(amount float32) {
	u.CreditUsed = u.CreditUsed - amount
}

func (u *user) GetRemainingCredit() float32 {
	return u.CreditLimit - u.CreditUsed
}
