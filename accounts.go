package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BankAccount struct {
	Name      string
	AccountID int
	Balance   float64
}

func main() {

	content, err := ioutil.ReadFile("accounts.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	var accounts []BankAccount

	err2 := json.Unmarshal(content, &accounts)

	if err2 != nil {
		fmt.Println("Error in unmarshalling json file")
	}

	for i := range accounts {
		accounts[i].Balance = accounts[i].Balance + 100.000
	}

	dat, _ := json.MarshalIndent(accounts, "", "")

	ioutil.WriteFile("accountsUpdated.json", dat, 0644)

}
