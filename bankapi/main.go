package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"strconv"
	"tamaki8021/bank_system/bankcore"
)

var accounts = map[float64]*bankcore.Account{}

func main()  {
	// 口座情報の初期化
	accounts[1001] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name: "John",
			Address: "Los Angeles, California",
			Phone: "(213) 555 0147",
		},
		Number: 1001,
	}

	http.HandleFunc("/statement", statement)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	// 口座番号の取得
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Invalid account number!")
		return 
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Sprintf(os.Stdout, account.Statement())
		}
	}

}