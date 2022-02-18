package user

import "errors"

var NAME_MISSING_ERROR = errors.New("name is missing in request")
var EMAIL_MISSING_ERROR = errors.New("email is missing in request")
var NEGATIVE_CREDIT_LIMIT_ERROR = errors.New("credit limit should be non-negative")
var CREDIT_LIMIT_HIT = errors.New("rejected! (reason: credit limit)")
var EXCESS_PAYBACK_ERROR = errors.New("payback is more than credit usage")

type UserOps interface {
	GetName() string
	UpdateCreditLimit(float32) error
	Purchase(float32) error
	Payback(float32) error
	GetRemainingCredit() float32
	GetCreditLimit() float32
}

type User struct {
	Name        string
	Email       string
	CreditLimit float32
	CreditUsed  float32
}

func GetNewUser(name, email string, creditLimit float32) (UserOps, error) {
	if name == "" {
		return nil, NAME_MISSING_ERROR
	}

	if email == "" {
		return nil, EMAIL_MISSING_ERROR
	}

	if creditLimit < 0 {
		return nil, NEGATIVE_CREDIT_LIMIT_ERROR
	}

	return &User{
		Name:        name,
		Email:       email,
		CreditLimit: creditLimit,
		CreditUsed:  0,
	}, nil
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) UpdateCreditLimit(creditLimit float32) error {
	if creditLimit < 0 {
		return NEGATIVE_CREDIT_LIMIT_ERROR
	}

	u.CreditLimit = creditLimit
	return nil
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
