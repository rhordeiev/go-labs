package main

import (
	"sync"
)

type Bank struct {
	name           string
	log            []string
	bankMoney      float64
	deposit        float64
	credit         float64
	depositPercent float64
	creditPercent  float64
	clients        []Client
	mutex          *sync.Mutex
}

func NewBank(enteredName string, enteredLog []string, enteredBankMoney, enteredDeposit, enteredCredit, enteredDepositPercent, enteredCreditPercent float64,
	enteredClients []Client, mutex *sync.Mutex) *Bank {
	bank := new(Bank)
	bank.name = enteredName
	bank.log = enteredLog
	bank.bankMoney = enteredBankMoney
	bank.deposit = enteredDeposit
	bank.credit = enteredCredit
	bank.depositPercent = enteredDepositPercent
	bank.creditPercent = enteredCreditPercent
	bank.clients = enteredClients
	bank.mutex = mutex
	return bank
}

func (currentBank *Bank) SetName(enteredName string) {
	currentBank.name = enteredName
}

func (currentBank *Bank) SetLog(enteredLog []string) {
	currentBank.log = enteredLog
}

func (currentBank *Bank) SetBankMoney(enteredBankMoney float64) {
	currentBank.bankMoney = enteredBankMoney
}

func (currentBank *Bank) SetDeposit(enteredDeposit float64) {
	currentBank.deposit = enteredDeposit
}

func (currentBank *Bank) SetCredit(enteredCredit float64) {
	currentBank.credit = enteredCredit
}

func (currentBank *Bank) SetClients(enteredClients []Client) {
	currentBank.clients = enteredClients
}

func (currentBank Bank) GetName() string {
	return currentBank.name
}

func (currentBank Bank) GetLog() []string {
	return currentBank.log
}

func (currentBank Bank) GetBankMoney() float64 {
	return currentBank.bankMoney
}

func (currentBank Bank) GetDeposit() float64 {
	return currentBank.deposit
}

func (currentBank Bank) GetCredit() float64 {
	return currentBank.credit
}

func (currentBank Bank) GetClients() []Client {
	return currentBank.clients
}

func (currentBank *Bank) AddStringToLog(enteredString string) {
	currentBank.mutex.Lock()
	currentBank.log = append(currentBank.log, enteredString)
	currentBank.mutex.Unlock()
}

func (currentBank *Bank) CalculateDeposit(enteredMoney float64, years int) float64 {
	currentBank.mutex.Lock()
	var resultMoney float64
	resultMoney = enteredMoney
	for i := 0; i < years; i++ {
		resultMoney += resultMoney * (currentBank.depositPercent / 100)
	}
	currentBank.SetDeposit(currentBank.GetDeposit() + resultMoney)
	currentBank.mutex.Unlock()
	return resultMoney
}

func (currentBank *Bank) CalculateCredit(enteredMoney float64, years int) float64 {
	currentBank.mutex.Lock()
	var resultMoney float64
	resultMoney = enteredMoney
	for i := 0; i < years; i++ {
		resultMoney += resultMoney * (currentBank.depositPercent / 100)
	}
	if currentBank.bankMoney > currentBank.GetCredit()+resultMoney {
		currentBank.SetCredit(currentBank.GetCredit() + resultMoney)
		currentBank.SetBankMoney(currentBank.GetBankMoney() - resultMoney)
	} else {
		resultMoney = -1
	}
	currentBank.mutex.Unlock()
	return resultMoney
}
