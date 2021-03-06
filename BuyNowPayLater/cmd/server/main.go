package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/goProjects/BuyNowPayLater/pkg/service"
)

func main() {
	fmt.Println("Starting BuyNowPayLater service...")
	bpnl := service.GetBPNLServer()

	//// Add users
	//bpnl.AddUser("u1", "u1@users.com", 300)
	//bpnl.AddUser("u2", "u2@users.com", 400)
	//bpnl.AddUser("u3", "u3@users.com", 500)
	//
	//// Add merchants
	//bpnl.AddMerchant("m1", "m1@gmail.com", 0.5)
	//bpnl.AddMerchant("m2", "m2@gmail.com", 1.5)
	//bpnl.AddMerchant("m3", "m3@gmail.com", 1.25)
	//
	//// New transactions
	//fmt.Println(bpnl.NewTransaction("u2", "m1", 500)) // error
	//fmt.Println(bpnl.NewTransaction("u1", "m2", 300)) // success
	//fmt.Println(bpnl.NewTransaction("u1", "m3", 10))  // error
	//
	//// Users at credit limit
	//fmt.Print("Users at credit limit: ")
	//fmt.Println(bpnl.GetUsersAtCreditLimit())
	//
	//// More transactions
	//fmt.Println(bpnl.NewTransaction("u3", "m3", 200)) // success
	//fmt.Println(bpnl.NewTransaction("u3", "m3", 300)) // success
	//
	//// Users at credit limit
	//fmt.Print("Users at credit limit: ")
	//fmt.Println(bpnl.GetUsersAtCreditLimit())
	//
	//// Report fee
	//fmt.Print("Merchant fee of m3: ")
	//fmt.Println(bpnl.GetMerchantFeeTotal("m3"))
	//
	//// Payback
	//fmt.Println(bpnl.Payback("u3", 400))
	//
	//// Total dues
	//fmt.Print("Total dues of all users: ")
	//fmt.Println(bpnl.GetTotalDues())

	cliProcessor(bpnl)
}

func cliProcessor(bpnl service.BNPLServiceOps) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal()
		}

		data := strings.Split(strings.TrimSuffix(line, "\n"), " ")

		switch {
		case strings.HasPrefix(line, "new user"):
			value, err := strconv.ParseFloat(data[4], 32)
			if err != nil {
				fmt.Println("wrong input, try again. Error: ", err)
			}
			fmt.Println(bpnl.AddUser(data[2], data[3], float32(value)))

		case strings.HasPrefix(line, "new merchant"):
			data[4] = TrimSuffix(data[4], "%")
			value, err := strconv.ParseFloat(data[4], 32)
			if err != nil {
				fmt.Println("wrong input, try again. Error: ", err)
			}
			fmt.Println(bpnl.AddMerchant(data[2], data[3], float32(value)))

		case strings.HasPrefix(line, "new txn"):
			value, err := strconv.ParseFloat(data[4], 32)
			if err != nil {
				fmt.Println("wrong input, try again. Error: ", err)
			}
			fmt.Println(bpnl.NewTransaction(data[2], data[3], float32(value)))

		case strings.HasPrefix(line, "update user"):
			value, err := strconv.ParseFloat(data[3], 32)
			if err != nil {
				fmt.Println("wrong input, try again. Error: ", err)
			}
			fmt.Println(bpnl.UpdateUserCreditLimit(data[2], float32(value)))

		case strings.HasPrefix(line, "update merchant"):
			data[3] = TrimSuffix(data[3], "%")
			value, err := strconv.ParseFloat(data[3], 32)
			if err != nil {
				fmt.Println("wrong input, try again. Error: ", err)
			}
			fmt.Println(bpnl.UpdateMerchantFee(data[2], float32(value)))

		case strings.HasPrefix(line, "payback"):
			value, err := strconv.ParseFloat(data[2], 32)
			if err != nil {
				fmt.Println("wrong input, try again. Error: ", err)
			}
			fmt.Println(bpnl.Payback(data[1], float32(value)))

		case strings.HasPrefix(line, "report fee"):
			fmt.Println(bpnl.GetMerchantFeeTotal(data[2]))

		case strings.HasPrefix(line, "report dues"):
			fmt.Println(bpnl.GetDuesOfUser(data[2]))

		case strings.HasPrefix(line, "report users-at-credit-limit"):
			fmt.Println(bpnl.GetUsersAtCreditLimit())

		case strings.HasPrefix(line, "report total-dues"):
			fmt.Println(bpnl.GetTotalDues())

		default:
			fmt.Println("Invalid input")
		}
	}
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
