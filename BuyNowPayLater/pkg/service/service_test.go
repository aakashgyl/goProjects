package service

import (
	"reflect"
	"testing"
)

func TestBNPLServer_AddMerchant(t *testing.T) {
	type args struct {
		name  string
		email string
		fee   float32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid details",
			args: args{
				name:  "user1",
				email: "user1@gmail.com",
				fee:   2.5,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()
			if err := server.AddMerchant(tt.args.name, tt.args.email, tt.args.fee); (err != nil) != tt.wantErr {
				t.Errorf("AddMerchant() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBNPLServer_AddUser(t *testing.T) {
	type args struct {
		name  string
		email string
		limit float32
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid details",
			args: args{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()
			if err := server.AddUser(tt.args.name, tt.args.email, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBNPLServer_GetDuesOfUser(t *testing.T) {
	type userArgs struct {
		name  string
		email string
		limit float32
	}
	type merchantArgs struct {
		name  string
		email string
		fee   float32
	}

	tests := []struct {
		name      string
		uArgs     userArgs
		mArgs     merchantArgs
		txnAmount float32
		wantErr   bool
	}{
		{
			name: "valid details",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			txnAmount: 500,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()
			server.AddUser(tt.uArgs.name, tt.uArgs.email, tt.uArgs.limit)
			server.AddMerchant(tt.mArgs.name, tt.mArgs.email, tt.mArgs.fee)
			server.NewTransaction(tt.uArgs.name, tt.mArgs.name, tt.txnAmount)

			if got := server.GetDuesOfUser(tt.uArgs.name); got != tt.txnAmount {
				t.Errorf("GetDuesOfUser() = %v, want %v", got, tt.txnAmount)
			}
		})
	}
}

func TestBNPLServer_GetMerchantFeeTotal(t *testing.T) {
	type userArgs struct {
		name  string
		email string
		limit float32
	}
	type merchantArgs struct {
		name  string
		email string
		fee   float32
	}

	tests := []struct {
		name        string
		uArgs       userArgs
		mArgs       merchantArgs
		txnAmount   float32
		merchantFee float32
		wantErr     bool
	}{
		{
			name: "valid details",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			txnAmount:   500,
			merchantFee: 25,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()
			server.AddUser(tt.uArgs.name, tt.uArgs.email, tt.uArgs.limit)
			server.AddMerchant(tt.mArgs.name, tt.mArgs.email, tt.mArgs.fee)
			server.NewTransaction(tt.uArgs.name, tt.mArgs.name, tt.txnAmount)

			if got := server.GetMerchantFeeTotal(tt.mArgs.name); got != tt.merchantFee {
				t.Errorf("GetMerchantFeeTotal() = %v, want %v", got, tt.merchantFee)
			}
		})
	}
}

func TestBNPLServer_GetUsersAtCreditLimit(t *testing.T) {
	type userArgs struct {
		name  string
		email string
		limit float32
	}
	type merchantArgs struct {
		name  string
		email string
		fee   float32
	}

	tests := []struct {
		name         string
		uArgs        userArgs
		mArgs        merchantArgs
		txnAmount    float32
		usersAtLimit []string
		wantErr      bool
	}{
		{
			name: "one user at limit",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			txnAmount:    500,
			usersAtLimit: []string{"user1"},
			wantErr:      false,
		},
		{
			name: "no user at limit",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			txnAmount:    300,
			usersAtLimit: []string{},
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()
			server.AddUser(tt.uArgs.name, tt.uArgs.email, tt.uArgs.limit)
			server.AddMerchant(tt.mArgs.name, tt.mArgs.email, tt.mArgs.fee)
			server.NewTransaction(tt.uArgs.name, tt.mArgs.name, tt.txnAmount)
			if got := server.GetUsersAtCreditLimit(); !reflect.DeepEqual(got, tt.usersAtLimit) {
				t.Errorf("GetUsersAtCreditLimit() = %v, want %v", got, tt.usersAtLimit)
			}
		})
	}
}

func TestBNPLServer_NewTransaction(t *testing.T) {
	type userArgs struct {
		name  string
		email string
		limit float32
	}
	type merchantArgs struct {
		name  string
		email string
		fee   float32
	}

	tests := []struct {
		name      string
		uArgs     userArgs
		mArgs     merchantArgs
		txnAmount float32
		wantErr   bool
	}{
		{
			name: "purchase same as limit",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			txnAmount: 500,
			wantErr:   false,
		},
		{
			name: "purchase beyond limit",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			txnAmount: 600,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()
			server.AddUser(tt.uArgs.name, tt.uArgs.email, tt.uArgs.limit)
			server.AddMerchant(tt.mArgs.name, tt.mArgs.email, tt.mArgs.fee)

			if err := server.NewTransaction(tt.uArgs.name, tt.mArgs.name, tt.txnAmount); (err != nil) != tt.wantErr {
				t.Errorf("NewTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBNPLServer_Payback(t *testing.T) {
	type userArgs struct {
		name  string
		email string
		limit float32
	}
	type merchantArgs struct {
		name  string
		email string
		fee   float32
	}

	tests := []struct {
		name      string
		uArgs     userArgs
		mArgs     merchantArgs
		txnAmount float32
		payback   float32
		wantErr   bool
	}{
		{
			name: "payment less than purchase",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			txnAmount: 500,
			payback:   25,
			wantErr:   false,
		},
		{
			name: "payment more than purchase",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			txnAmount: 500,
			payback:   25000,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()
			server.AddUser(tt.uArgs.name, tt.uArgs.email, tt.uArgs.limit)
			server.AddMerchant(tt.mArgs.name, tt.mArgs.email, tt.mArgs.fee)
			server.NewTransaction(tt.uArgs.name, tt.mArgs.name, tt.txnAmount)

			if err := server.Payback(tt.uArgs.name, tt.payback); (err != nil) != tt.wantErr {
				t.Errorf("Payback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBNPLServer_UpdateMerchantFee(t *testing.T) {
	type merchantArgs struct {
		name  string
		email string
		fee   float32
	}

	tests := []struct {
		name                  string
		mArgs                 merchantArgs
		newMerchantFeePercent float32
		wantErr               bool
	}{
		{
			name: "valid details",
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			newMerchantFeePercent: 2,
			wantErr:               false,
		},
		{
			name: "negative fee",
			mArgs: merchantArgs{
				name:  "m1",
				email: "m1@gmail.com",
				fee:   5,
			},
			newMerchantFeePercent: -2,
			wantErr:               true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()
			server.AddMerchant(tt.mArgs.name, tt.mArgs.email, tt.mArgs.fee)

			if err := server.UpdateMerchantFee(tt.mArgs.name, tt.newMerchantFeePercent); (err != nil) != tt.wantErr {
				t.Errorf("UpdateMerchantFee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBNPLServer_UpdateUserCreditLimit(t *testing.T) {
	type userArgs struct {
		name  string
		email string
		limit float32
	}

	tests := []struct {
		name           string
		uArgs          userArgs
		newCreditLimit float32
		wantErr        bool
	}{
		{
			name: "valid details",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			newCreditLimit: 700,
			wantErr:        false,
		},
		{
			name: "negative credit limit",
			uArgs: userArgs{
				name:  "user1",
				email: "user1@gmail.com",
				limit: 500,
			},
			newCreditLimit: -700,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()
			server.AddUser(tt.uArgs.name, tt.uArgs.email, tt.uArgs.limit)

			if err := server.UpdateUserCreditLimit(tt.uArgs.name, tt.newCreditLimit); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserCreditLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetBPNLServer(t *testing.T) {
	tests := []struct {
		name    string
		wantNil bool
	}{
		{
			name:    "valid details",
			wantNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := GetBPNLServer()

			if tt.wantNil && server != nil || !tt.wantNil && server == nil {
				t.Fail()
			}
		})
	}
}
