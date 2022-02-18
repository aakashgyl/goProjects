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

func main(){
	fmt.Println("Starting BuyNowPayLater service")
	bpnl := service.GetBPNLServer()
	cliProcessor(bpnl)
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

		//new user u1 u1@email.in 1000 # name, email, credit-limit
		//new merchant m1 2% # name, merchant-fee-percentage
		//new txn u1 m1 400 # user, merchant, txn-amount
		//update merchant m1 1% # merchant, new-mmerchant-fee-perc
		//payback u1 300 # user, payback-amount
		//report fee m1
		//report dues u1
		//report users-at-credit-limit
		//report total-dues

		switch {
		case strings.HasPrefix(line, "new user"):
			value, _ := strconv.ParseFloat(data[4], 32)
			bpnl.AddUser(data[2], data[3], float32(value))

		case strings.HasPrefix(line, "new merchant"):
			value, _ := strconv.ParseFloat(data[3], 32)
			bpnl.AddMerchant(data[2], float32(value))

		case strings.HasPrefix(line, "new txn"):
			value, _ := strconv.ParseFloat(data[4], 32)
			bpnl.NewTransaction(data[2], data[3], float32(value))

		case strings.HasPrefix(line, "update merchant"):
			value, _ := strconv.ParseFloat(data[3], 32)
			bpnl.UpdateMerchantFee(data[2], float32(value))

		case strings.HasPrefix(line, "payback"):
			value, _ := strconv.ParseFloat(data[2], 32)
			bpnl.Payback(data[1], float32(value))

		case strings.HasPrefix(line, "report fee"):
			fmt.Println(bpnl.GetMerchantFeeTotal())

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
