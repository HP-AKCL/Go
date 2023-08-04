package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/akcl-hp/Go/bank"
)

var accounts = map[float64]*bank.Account{}

func main() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "john",
			Address: "los Angeles",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	//http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:2019", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	numberqs := req.URL.Query().Get("number")
	if numberqs == "" {
		fmt.Fprintf(w, "number is null")
		return
	} else if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "number %v is invalid", number)
	} else {
		account, ok := accounts[number]
		if ok {
			fmt.Fprintf(w, account.Statement())
		} else {
			fmt.Fprintf(w, "%v", number)
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")
	if numberqs == "" {
		fmt.Fprintf(w, "account number is misssing")
	} else if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "invalid number")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "invalid amount")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "account canot be found")
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")
	destqs := req.URL.Query().Get("dest")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else if dest, err := strconv.ParseFloat(destqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account destination number!")
	} else {
		if accountA, ok := accounts[number]; !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else if accountB, ok := accounts[dest]; !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", dest)
		} else {
			err := accountA.Transfer(amount, accountB)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, accountA.Statement())
			}
		}
	}
}

// CustomAccount ...
type CustomAccount struct {
	*bank.Account
}

// Statement ...
func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(json)
}
