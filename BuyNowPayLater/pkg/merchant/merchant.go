package merchant

type MerchantOps interface {
	GetName() string
	GetCurrentCommsion() float32
	GetCommissionPaid() float32
	UpdateMerchantFeePercent(float32)
	Purchase(float32)
}

type Merchant struct {
	Name                     string
	CurrentCommissionPercent float32
	CommissionPaid           float32
}

func GetNewMerchant(name string, fee float32) MerchantOps {
	return &Merchant{
		Name:                     name,
		CurrentCommissionPercent: fee,
	}
}

func (m *Merchant) UpdateMerchantFeePercent(fee float32) {
	m.CurrentCommissionPercent = fee
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
