package user

import (
	"testing"
)

func TestGetNewUser(t *testing.T) {
	type args struct {
		name        string
		email       string
		creditLimit float32
	}

	tests := []struct {
		name        string
		args        args
		errExpected bool
	}{
		{
			name: "valid details",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			errExpected: false,
		},
		{
			name: "invalid details: name missing",
			args: args{
				name:        "",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			errExpected: true,
		},
		{
			name: "invalid details: email missing",
			args: args{
				name:        "user1",
				email:       "",
				creditLimit: 500,
			},
			errExpected: true,
		},
		{
			name: "invalid details: negative credit limit",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: -500,
			},
			errExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := GetNewUser(tt.args.name, tt.args.email, tt.args.creditLimit)
			if err == nil && tt.errExpected || err != nil && !tt.errExpected {
				t.Fail()
			}

			if user.GetName() != tt.args.name {
				t.Error("Name mismatch error")
			}

			if user.GetCreditLimit() != tt.args.creditLimit {
				t.Errorf("Credit limit mismatch error")
			}
		})
	}
}

func TestUser_GetRemainingCredit(t *testing.T) {
	type args struct {
		name        string
		email       string
		creditLimit float32
	}

	tests := []struct {
		name            string
		args            args
		purchase        float32
		remainingCredit float32
	}{
		{
			name: "valid data",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			purchase:        300,
			remainingCredit: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, _ := GetNewUser(tt.args.name, tt.args.email, tt.args.creditLimit)
			user.Purchase(tt.purchase)
			if got := user.GetRemainingCredit(); got != tt.remainingCredit {
				t.Errorf("GetRemainingCredit() = %v, want %v", got, tt.remainingCredit)
			}
		})
	}
}

func TestUser_Payback(t *testing.T) {
	type args struct {
		name        string
		email       string
		creditLimit float32
	}

	tests := []struct {
		name     string
		args     args
		purchase float32
		payback  float32
		wantErr  bool
	}{
		{
			name: "payback less than purchase",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			purchase: 300,
			payback:  200,
			wantErr:  false,
		},
		{
			name: "payback equal to purchase",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			purchase: 300,
			payback:  300,
			wantErr:  false,
		},
		{
			name: "payback more than purchase",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			purchase: 300,
			payback:  400,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, _ := GetNewUser(tt.args.name, tt.args.email, tt.args.creditLimit)
			user.Purchase(tt.purchase)

			if err := user.Payback(tt.payback); (err != nil) != tt.wantErr {
				t.Errorf("Payback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Purchase(t *testing.T) {
	type args struct {
		name        string
		email       string
		creditLimit float32
	}

	tests := []struct {
		name     string
		args     args
		purchase float32
		wantErr  bool
	}{
		{
			name: "purchase less than credit limit",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			purchase: 300,
			wantErr:  false,
		},
		{
			name: "purchase more than credit limit",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			purchase: 700,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, _ := GetNewUser(tt.args.name, tt.args.email, tt.args.creditLimit)

			if err := user.Purchase(tt.purchase); (err != nil) != tt.wantErr {
				t.Errorf("Purchase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_UpdateCreditLimit(t *testing.T) {
	type args struct {
		name        string
		email       string
		creditLimit float32
	}

	tests := []struct {
		name           string
		args           args
		newCreditLimit float32
		wantErr        bool
	}{
		{
			name: "update credit limit",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			newCreditLimit: 700,
			wantErr:        false,
		},
		{
			name: "negative credit limit",
			args: args{
				name:        "user1",
				email:       "user1@gmail.com",
				creditLimit: 500,
			},
			newCreditLimit: -100,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, _ := GetNewUser(tt.args.name, tt.args.email, tt.args.creditLimit)

			if err := user.UpdateCreditLimit(tt.newCreditLimit); (err != nil) != tt.wantErr {
				t.Errorf("Purchase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
