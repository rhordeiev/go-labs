package main

import (
	"fmt"
)

type Client struct {
	name          string
	surname       string
	accountNumber uint
	cDeposit      float64
	cCredit       float64
}

func NewClient(enteredName, enteredSurname string, enteredAccountNumber uint, enteredDeposit, enteredCredit float64) *Client {
	client := new(Client)
	client.name = enteredName
	client.surname = enteredSurname
	client.accountNumber = enteredAccountNumber
	client.cDeposit = enteredDeposit
	client.cCredit = enteredCredit
	return client
}

func (currentClient *Client) SetName(enteredName string) {
	currentClient.name = enteredName
}

func (currentClient *Client) SetSurname(enteredSurname string) {
	currentClient.surname = enteredSurname
}

func (currentClient *Client) SetAccountNumber(enteredAccountNumber uint) {
	currentClient.accountNumber = enteredAccountNumber
}

func (currentClient *Client) SetDeposit(enteredDeposit float64) {
	currentClient.cDeposit = enteredDeposit
}

func (currentClient *Client) SetCredit(enteredCredit float64) {
	currentClient.cCredit = enteredCredit
}

func (currentClient Client) GetName() string {
	return currentClient.name
}

func (currentClient Client) GetSurname() string {
	return currentClient.surname
}

func (currentClient Client) GetAccountNumber() uint {
	return currentClient.accountNumber
}

func (currentClient Client) GetDeposit() float64 {
	return currentClient.cDeposit
}

func (currentClient Client) GetCredit() float64 {
	return currentClient.cCredit
}

func (currentClient *Client) SetDepositToBank(enteredDepositMoney float64, enteredYears int, bank *Bank) {
	currentClient.SetDeposit(currentClient.GetDeposit() + bank.CalculateDeposit(enteredDepositMoney, enteredYears))
	bank.AddStringToLog(currentClient.name + " " + currentClient.surname + " поклав на депозит " + fmt.Sprintf("%f", bank.CalculateDeposit(enteredDepositMoney, enteredYears)))
}

func (currentClient *Client) SetCreditToBank(enteredCreditMoney float64, enteredYears int, bank *Bank) {
	if bank.CalculateCredit(enteredCreditMoney, enteredYears) != -1 {
		currentClient.SetCredit(currentClient.GetCredit() + bank.CalculateCredit(enteredCreditMoney, enteredYears))
		bank.AddStringToLog(currentClient.name + " " + currentClient.surname + " взяв у кредит " + fmt.Sprintf("%f", bank.CalculateCredit(enteredCreditMoney, enteredYears)))
	} else {
		bank.AddStringToLog(currentClient.name + " " + currentClient.surname + " недостатньо грошей, щоб взяти кредит!")
		currentClient.SetCredit(-1)
		watchingMoneyChannel = -1
	}
}

func (currentClient *Client) GetDepositFromBank(enteredDepositMoney float64, bank *Bank) {
	if currentClient.GetDeposit()-enteredDepositMoney > 0 && bank.GetDeposit()-enteredDepositMoney > 0 {
		currentClient.SetDeposit(currentClient.GetDeposit() - enteredDepositMoney)
		bank.SetDeposit(bank.GetDeposit() - enteredDepositMoney)
		bank.AddStringToLog(currentClient.name + " " + currentClient.surname + " забрав депозит " + fmt.Sprintf("%f", enteredDepositMoney))
	} else {
		bank.AddStringToLog(currentClient.name + " " + currentClient.surname + " неможливо забрати депозит!")
		currentClient.SetDeposit(-1)
		watchingMoneyChannel = -1
	}
}

func (currentClient *Client) ReturnCreditToBank(enteredCreditMoney float64, bank *Bank) {
	if currentClient.GetCredit()-enteredCreditMoney > 0 && bank.GetCredit()-enteredCreditMoney > 0 {
		currentClient.SetCredit(currentClient.GetCredit() - enteredCreditMoney)
		bank.SetCredit(bank.GetCredit() - enteredCreditMoney)
		bank.SetBankMoney(bank.GetBankMoney() + enteredCreditMoney)
		bank.AddStringToLog(currentClient.name + " " + currentClient.surname + " відплатив кредит на суму " + fmt.Sprintf("%f", enteredCreditMoney))
	} else {
		bank.AddStringToLog(currentClient.name + " " + currentClient.surname + " неможливо відплатити кредит!")
		currentClient.SetCredit(-1)
		watchingMoneyChannel = -1
	}
}
