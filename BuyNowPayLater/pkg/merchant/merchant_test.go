package merchant

import (
	"testing"
)

func TestGetNewMerchant(t *testing.T) {
	type args struct {
		name string
		email string
		fee  float32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid details",
			args: args{
				name: "user1",
				email: "user1@gmail.com",
				fee:  2.5,
			},
			wantErr: false,
		},
		{
			name: "invalid details: name missing",
			args: args{
				name: "",
				email: "user1@gmail.com",
				fee:  2,
			},
			wantErr: true,
		},
		{
			name: "invalid details: negative fee",
			args: args{
				name: "user1",
				fee:  -2,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merchant, err := GetNewMerchant(tt.args.name, tt.args.email, tt.args.fee)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNewMerchant() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if merchant.GetName() != tt.args.name {
					t.Fail()
				}

				if merchant.GetCurrentCommsion() != tt.args.fee {
					t.Fail()
				}
			}
		})
	}
}

func TestMerchant_GetCommissionPaid(t *testing.T) {
	type args struct {
		name string
		email string
		fee  float32
	}
	tests := []struct {
		name    string
		args    args
		commission float32
	}{
		{
			name: "valid details",
			args: args{
				name: "user1",
				email: "user1@gmail.com",
				fee:  2.5,
			},
			commission: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merchant, _ := GetNewMerchant(tt.args.name, tt.args.email, tt.args.fee)
			if commission := merchant.GetCommissionPaid(); commission != tt.commission {
				t.Errorf("GetCommissionPaid() = %v, want %v", commission, tt.commission)
			}
		})
	}
}

func TestMerchant_GetCurrentCommsion(t *testing.T) {
	type args struct {
		name string
		email string
		fee  float32
	}

	tests := []struct {
		name    string
		args    args
		updatedCommission float32
	}{
		{
			name: "valid details",
			args: args{
				name: "user1",
				email: "user1@gmail.com",
				fee:  2.5,
			},
		},
		{
			name: "updated commission",
			args: args{
				name: "user1",
				email: "user1@gmail.com",
				fee:  2.5,
			},
			updatedCommission: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merchant, _ := GetNewMerchant(tt.args.name, tt.args.email, tt.args.fee)
			if got := merchant.GetCurrentCommsion(); got != tt.args.fee {
				t.Errorf("GetCurrentCommsion() = %v, want %v", got, tt.args.fee)
			}

			if tt.updatedCommission != 0 {
				merchant.UpdateMerchantFeePercent(tt.updatedCommission)
				if merchant.GetCurrentCommsion() != tt.updatedCommission {
					t.Fail()
				}
			}
		})
	}
}

func TestMerchant_GetName(t *testing.T) {
	type args struct {
		name string
		email string
		fee  float32
	}
	tests := []struct {
		name    string
		args    args
	}{
		{
			name: "valid details",
			args: args{
				name: "user1",
				email: "user1@gmail.com",
				fee:  2.5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merchant, _ := GetNewMerchant(tt.args.name, tt.args.email, tt.args.fee)
			if got := merchant.GetName(); got != tt.args.name {
				t.Errorf("GetName() = %v, want %v", got, tt.args.name)
			}
		})
	}
}

func TestMerchant_Purchase(t *testing.T) {
	type args struct {
		name string
		email string
		fee  float32
	}
	tests := []struct {
		name    string
		args    args
		amount float32
		commissionPaid float32
	}{
		{
			name: "valid details",
			args: args{
				name: "user1",
				email: "user1@gmail.com",
				fee:  5,
			},
			amount: 2000,
			commissionPaid: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merchant, _ := GetNewMerchant(tt.args.name, tt.args.email, tt.args.fee)
			merchant.Purchase(tt.amount)
			if got := merchant.GetCommissionPaid(); got != tt.commissionPaid {
				t.Errorf("GetCommissionPaid() = %v, want %v", got, tt.args.name)
			}
		})
	}
}

func TestMerchant_UpdateMerchantFeePercent(t *testing.T) {
	type args struct {
		name string
		email string
		fee  float32
	}
	tests := []struct {
		name    string
		args    args
		updateFee float32
		wantErr bool
	}{
		{
			name: "valid details",
			args: args{
				name: "user1",
				email: "user1@gmail.com",
				fee:  2.5,
			},
			updateFee: 5,
			wantErr: false,
		},
		{
			name: "negative fee",
			args: args{
				name: "user1",
				email: "user1@gmail.com",
				fee:  2.5,
			},
			updateFee: -5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merchant, _ := GetNewMerchant(tt.args.name, tt.args.email, tt.args.fee)
			if err := merchant.UpdateMerchantFeePercent(tt.updateFee); (err != nil) != tt.wantErr {
				t.Errorf("UpdateMerchantFeePercent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
