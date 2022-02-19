package merchant

import "errors"

var NAME_MISSING_ERROR = errors.New("name is missing in request")
var NEGATIVE_FEE_ERROR = errors.New("fee should be non-negative")

type MerchantOps interface {
	GetName() string
	GetCurrentCommsion() float32
	GetCommissionPaid() float32
	UpdateMerchantFeePercent(float32) error
	Purchase(float32)
}

type Merchant struct {
	Name                     string
	Email                    string
	CurrentCommissionPercent float32
	CommissionPaid           float32
}

func GetNewMerchant(name, email string, fee float32) (MerchantOps, error) {
	if name == "" {
		return nil, NAME_MISSING_ERROR
	}

	if fee < 0 {
		return nil, NEGATIVE_FEE_ERROR
	}

	return &Merchant{
		Name:                     name,
		Email:                    email,
		CurrentCommissionPercent: fee,
	}, nil
}

func (m *Merchant) UpdateMerchantFeePercent(fee float32) error {
	if fee < 0 {
		return NEGATIVE_FEE_ERROR
	}
	m.CurrentCommissionPercent = fee
	return nil
}

func (m *Merchant) Purchase(amount float32) {
	m.CommissionPaid = m.CommissionPaid + amount*m.CurrentCommissionPercent/100
}

func (m *Merchant) GetName() string {
	return m.Name
}

func (m *Merchant) GetCurrentCommsion() float32 {
	return m.CurrentCommissionPercent
}

func (m *Merchant) GetCommissionPaid() float32 {
	return m.CommissionPaid
}
