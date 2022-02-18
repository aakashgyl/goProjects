package service

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/goProjects/BuyNowPayLater/pkg/merchant"
	"github.com/goProjects/BuyNowPayLater/pkg/user"
)

var USER_NOT_FOUND = errors.New("user not found")
var MERCHANT_NOT_FOUND = errors.New("merchant not found")

type BNPLServiceOps interface {
	AddUser(name, email string, limit float32)
	AddMerchant(name string, fee float32)
	NewTransaction(userName, merchantName string, txnAmount float32) error
	UpdateUserCreditLimit(name string, limit float32)
	UpdateMerchantFee(merchantName string, newFee float32)
	Payback(user string, amount float32) error
	GetMerchantFeeTotal(name string) float32
	GetDuesOfUser(name string) float32
	GetUsersAtCreditLimit() []string
	GetTotalDues() []string
}

type BNPLServer struct {
	Users     map[string]user.UserOps
	Merchants map[string]merchant.MerchantOps
}

func GetBPNLServer() BNPLServiceOps {
	return &BNPLServer{
		Users:     make(map[string]user.UserOps),
		Merchants: make(map[string]merchant.MerchantOps),
	}
}

func (server *BNPLServer) AddUser(name, email string, limit float32) {
	newUser := user.GetNewUser(name, email, limit)
	server.Users[name] = newUser
	log.Infof("User %q added successfully", name)
}

func (server *BNPLServer) AddMerchant(name string, fee float32) {
	newMerchant := merchant.GetNewMerchant(name, fee)
	server.Merchants[name] = newMerchant
	log.Infof("Merchant %q added successfully", name)
}

func (server *BNPLServer) NewTransaction(userName, merchantName string, txnAmount float32) error {
	var userData user.UserOps
	var merchantData merchant.MerchantOps
	var ok bool

	if userData, ok = server.Users[userName]; !ok {
		return USER_NOT_FOUND
	}

	if merchantData, ok = server.Merchants[merchantName]; !ok {
		return MERCHANT_NOT_FOUND
	}

	err := userData.Purchase(txnAmount)
	if err != nil {
		log.Errorf("Purchase of %v by %q from %q failed due to: %v", txnAmount, userName, merchantName, err.Error())
		return err
	}
	merchantData.Purchase(txnAmount)

	log.Infof("Purchase by %q from %q of %v success", userName, merchantName, txnAmount)
	return nil
}

func (server *BNPLServer) UpdateUserCreditLimit(name string, limit float32) {
	server.Users[name].UpdateCreditLimit(limit)
	log.Info("User credit limit updated")
}

func (server *BNPLServer) UpdateMerchantFee(name string, newFee float32) {
	server.Merchants[name].UpdateMerchantFeePercent(newFee)
	log.Info("Merchant fee updated")
}

func (server *BNPLServer) Payback(userName string, amount float32) error {
	err := server.Users[userName].Payback(amount)
	if err != nil {
		log.Errorf("Payback from %q failed due to: %v", userName, err.Error())
		return err
	}

	log.Infof("Received %v from %s", amount, userName)
	return nil
}

func (server *BNPLServer) GetMerchantFeeTotal(name string) float32 {
	return server.Merchants[name].GetCommissionPaid()
}

func (server *BNPLServer) GetDuesOfUser(userName string) float32 {
	userData := server.Users[userName]
	return userData.GetCreditLimit() - userData.GetRemainingCredit()
}

func (server *BNPLServer) GetUsersAtCreditLimit() []string {
	usersAtCreditLimit := []string{}
	for _, userData := range server.Users {
		if userData.GetRemainingCredit() == 0 {
			usersAtCreditLimit = append(usersAtCreditLimit, userData.GetName())
		}
	}
	return usersAtCreditLimit
}

func (server *BNPLServer) GetTotalDues() []string {
	var total float32
	var totalDuesList []string

	for _, userData := range server.Users {
		dues := userData.GetCreditLimit() - userData.GetRemainingCredit()
		if dues != 0 {
			totalDuesList = append(totalDuesList, userData.GetName()+": "+fmt.Sprint(dues))
			total = total + dues
		}
	}
	totalDuesList = append(totalDuesList, "Total:"+fmt.Sprint(total))
	return totalDuesList
}
