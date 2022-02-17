package merchant

type MerchantOps interface {
	UpdateMerchantFeePercent(float32)
	Purchase(float32)
}

type merchant struct {
	Name string
	CurrentCommissionPercent float32
	CommissionPaid float32
}

func GetNewMerchant(name string, fee float32) MerchantOps {
	return &merchant{
		Name:       name,
		CurrentCommissionPercent: fee,
	}
}

func (m *merchant) UpdateMerchantFeePercent(fee float32) {
	m.CurrentCommissionPercent = fee
}

func (m *merchant) Purchase(amount float32) {
	m.CommissionPaid = m.CommissionPaid + amount * m.CurrentCommissionPercent/100
}
