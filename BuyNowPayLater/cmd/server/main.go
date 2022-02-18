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

	// Add users
	bpnl.AddUser("u1", "u1@users.com", 300)
	bpnl.AddUser("u2", "u2@users.com", 400)
	bpnl.AddUser("u3", "u3@users.com", 500)

	// Add merchants
	bpnl.AddMerchant("m1", 0.5)
	bpnl.AddMerchant("m2", 1.5)
	bpnl.AddMerchant("m3", 1.25)

	// New transactions
	fmt.Println(bpnl.NewTransaction("u2", "m1", 500)) // error
	fmt.Println(bpnl.NewTransaction("u1", "m2", 300)) // success
	fmt.Println(bpnl.NewTransaction("u1", "m3", 10))  // error

	// Users at credit limit
	fmt.Print("Users at credit limit: ")
	fmt.Println(bpnl.GetUsersAtCreditLimit())

	// More transactions
	fmt.Println(bpnl.NewTransaction("u3", "m3", 200)) // success
	fmt.Println(bpnl.NewTransaction("u3", "m3", 300)) // success

	// Users at credit limit
	fmt.Print("Users at credit limit: ")
	fmt.Println(bpnl.GetUsersAtCreditLimit())

	// Report fee
	fmt.Print("Merchant fee of m3: ")
	fmt.Println(bpnl.GetMerchantFeeTotal("m3"))

	// Payback
	bpnl.Payback("u3", 400)

	// Total dues
	fmt.Print("Total dues of all users: ")
	fmt.Println(bpnl.GetTotalDues())

	//cliProcessor(bpnl)
}

func cliProcessor(bpnl service.BNPLServiceOps) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal()
		}

		data := strings.Split(line, " ")

		switch {
		case strings.HasPrefix(line, "new user"):
			value, _ := strconv.ParseFloat(data[4], 32)
			bpnl.AddUser(data[2], data[3], float32(value))

		case strings.HasPrefix(line, "new merchant"):
			value, _ := strconv.ParseFloat(data[3], 32)
			bpnl.AddMerchant(data[2], float32(value))

		case strings.HasPrefix(line, "new txn"):
			value, _ := strconv.ParseFloat(data[4], 32)
			fmt.Println(bpnl.NewTransaction(data[2], data[3], float32(value)))

		case strings.HasPrefix(line, "update user"):
			value, _ := strconv.ParseFloat(data[3], 32)
			bpnl.UpdateUserCreditLimit(data[2], float32(value))

		case strings.HasPrefix(line, "update merchant"):
			value, _ := strconv.ParseFloat(data[3], 32)
			bpnl.UpdateMerchantFee(data[2], float32(value))

		case strings.HasPrefix(line, "payback"):
			value, _ := strconv.ParseFloat(data[2], 32)
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
