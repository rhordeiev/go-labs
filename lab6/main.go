package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var watchingMoneyChannel int

func main() {
	var mutex sync.Mutex
	var bank *Bank
	var clients []Client
	var menuOption int
	//var log []string
	//bank = NewBank("ПриватБанк", log, 200000, 5000, 5000, 1.5, 1.7, clients, &mutex)
	for {
		fmt.Println("---------МЕНЮ---------")
		fmt.Println("")
		fmt.Println("1. Cтворити банк")
		fmt.Println("2. Cтворити клієнта для роботи з кредитами")
		fmt.Println("3. Cтворити клієнта для роботи з депозитами")
		fmt.Println("4. Вивести інформацію про клієнта за прізвищем")
		fmt.Println("5. Вивести інформацію про клієнта за номером рахунку")
		fmt.Println("6. Вивести інформацію про виконані операції")
		fmt.Println("7. Завершити роботу програми")
		fmt.Println("")
		fmt.Printf("Оберіть пункт: ")
		fmt.Scan(&menuOption)
		switch menuOption {
		case 1:
			var enteredName string
			var enteredBankMoney, enteredDepositPercent, enteredCreditPercent float64
			var log []string
			fmt.Printf("Введіть назву банку: ")
			fmt.Scan(&enteredName)
			fmt.Printf("Введіть кількість грошей у банку: ")
			fmt.Scan(&enteredBankMoney)
			fmt.Printf("Введіть відсоток депозиту в банку: ")
			fmt.Scan(&enteredDepositPercent)
			fmt.Printf("Введіть відсоток кредиту в банку: ")
			fmt.Scan(&enteredCreditPercent)
			bank = NewBank(enteredName, log, enteredBankMoney, 5000, 5000, enteredDepositPercent, enteredCreditPercent, clients, &mutex)
		case 2:
			var tempClient *Client
			var enteredName, enteredSurname string
			var enteredAccountNumber uint
			fmt.Printf("Введіть назву: ")
			fmt.Scan(&enteredName)
			fmt.Printf("Введіть прізвище: ")
			fmt.Scan(&enteredSurname)
			fmt.Printf("Введіть номер рахунку: ")
			fmt.Scan(&enteredAccountNumber)
			tempClient = NewClient(enteredName, enteredSurname, enteredAccountNumber, 0, 0)
			clients = append(clients, *tempClient)
			bank.SetClients(clients)
			bank.GetClients()[len(bank.GetClients())-1].SetCredit(3000)
			bank.GetClients()[len(bank.GetClients())-1].SetDeposit(2000)
			go goRoutinesCredit(bank)
		case 3:
			var tempClient *Client
			var enteredName, enteredSurname string
			var enteredAccountNumber uint
			fmt.Printf("Введіть назву: ")
			fmt.Scan(&enteredName)
			fmt.Printf("Введіть прізвище: ")
			fmt.Scan(&enteredSurname)
			fmt.Printf("Введіть номер рахунку: ")
			fmt.Scan(&enteredAccountNumber)
			tempClient = NewClient(enteredName, enteredSurname, enteredAccountNumber, 0, 0)
			clients = append(clients, *tempClient)
			bank.SetClients(clients)
			bank.GetClients()[len(bank.GetClients())-1].SetCredit(3000)
			bank.GetClients()[len(bank.GetClients())-1].SetDeposit(2000)
			go goRoutinesDeposit(bank)
		case 4:
			var enteredSurname string
			fmt.Printf("Введіть прізвище для пошуку: ")
			fmt.Scan(&enteredSurname)
			for i := 0; i < len(bank.GetClients()); i++ {
				if bank.GetClients()[i].GetSurname() == enteredSurname {
					fmt.Println("Ім'я: " + bank.GetClients()[i].GetName())
					fmt.Println("Прізвище: " + bank.GetClients()[i].GetSurname())
					fmt.Println("Номер рахунку: " + fmt.Sprintf("%d", bank.GetClients()[i].GetAccountNumber()))
					fmt.Println("Сума депозиту: " + fmt.Sprintf("%f", bank.GetClients()[i].GetDeposit()))
					fmt.Println("Сума кредиту: " + fmt.Sprintf("%f", bank.GetClients()[i].GetCredit()))
					break
				}
			}
		case 5:
			var enteredAccountNumber uint
			fmt.Printf("Введіть номер рахунку для пошуку: ")
			fmt.Scan(&enteredAccountNumber)
			for i := 0; i < len(bank.GetClients()); i++ {
				if bank.GetClients()[i].GetAccountNumber() == enteredAccountNumber {
					fmt.Println("Ім'я: " + bank.GetClients()[i].GetName())
					fmt.Println("Прізвище: " + bank.GetClients()[i].GetSurname())
					fmt.Println("Номер рахунку: " + fmt.Sprintf("%d", bank.GetClients()[i].GetAccountNumber()))
					fmt.Println("Сума депозиту: " + fmt.Sprintf("%f", bank.GetClients()[i].GetDeposit()))
					fmt.Println("Сума кредиту: " + fmt.Sprintf("%f", bank.GetClients()[i].GetCredit()))
					break
				}
			}
		case 6:
			for i := 0; i < len(bank.GetLog()); i++ {
				fmt.Printf(bank.GetLog()[i] + "\n")
			}
		case 7:
			os.Exit(1)
		case 8:
			for i := 0; i < len(bank.GetClients()); i++ {
				fmt.Printf(bank.GetClients()[i].GetSurname() + "\n")
			}
		}
	}
}

func goRoutinesCredit(bank *Bank) {
loopCredit:
	for {
		randNumber := rand.Intn(2) + 1
		switch randNumber {
		case 1:
			bank.GetClients()[len(bank.GetClients())-1].SetCreditToBank(150, 3, bank)
			if watchingMoneyChannel == -1 {
				break loopCredit
			}
			time.Sleep(1 * time.Second)
		case 2:
			bank.GetClients()[len(bank.GetClients())-1].ReturnCreditToBank(150, bank)
			if watchingMoneyChannel == -1 {
				break loopCredit
			}
			time.Sleep(1 * time.Second)
		default:
			bank.GetClients()[len(bank.GetClients())-1].SetCreditToBank(150, 3, bank)
			if watchingMoneyChannel == -1 {
				break loopCredit
			}
			time.Sleep(1 * time.Second)
		}
	}
}
func goRoutinesDeposit(bank *Bank) {
loopDeposit:
	for {
		randNumber := rand.Intn(2) + 1
		switch randNumber {
		case 1:
			bank.GetClients()[len(bank.GetClients())-1].SetDepositToBank(200, 2, bank)
			time.Sleep(1 * time.Second)
		case 2:
			bank.GetClients()[len(bank.GetClients())-1].GetDepositFromBank(200, bank)
			if watchingMoneyChannel == -1 {
				break loopDeposit
			}
			time.Sleep(1 * time.Second)
		default:
			bank.GetClients()[len(bank.GetClients())-1].SetDepositToBank(200, 2, bank)
			time.Sleep(1 * time.Second)
		}
	}
}
