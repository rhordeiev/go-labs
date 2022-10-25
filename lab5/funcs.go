package main

import (
	"errors"
	"fmt"
	"strconv"
)

func ReadWorkerArray() ([]*Worker, error) {
	var workersCount int
	var enteredName, enteredWorkPlace string
	var enteredYear, enteredMonth uint
	fmt.Printf("Введіть кількість працівників: ")
	fmt.Scan(&workersCount)
	var workers []*Worker = make([]*Worker, workersCount)
	for i := 0; i < workersCount; i++ {
		fmt.Println("Працівник №" + strconv.Itoa(i+1) + ": ")
		fmt.Printf("Введіть прізвище та ініціали: ")
		fmt.Scan(&enteredName)
		fmt.Printf("Введіть місце роботи: ")
		fmt.Scan(&enteredWorkPlace)
		fmt.Printf("Введіть рік початку роботи: ")
		fmt.Scan(&enteredYear)
		fmt.Printf("Введіть місяць початку роботи: ")
		fmt.Scan(&enteredMonth)
		if enteredName == "" || enteredWorkPlace == "" {
			return workers, errors.New("Поля прізвища та ініціалів і місця роботи повинні бути заповнені")
		}
		workers[i] = NewWorker(enteredName, enteredWorkPlace, enteredYear, enteredMonth)
		fmt.Println("")
	}
	return workers, nil
}

func PrintWorker(worker *Worker) {
	fmt.Println("Прізвище та ініціали: " + worker.GetName())
	fmt.Println("Місце роботи: " + worker.GetWorkPlace())
	fmt.Println("Введіть рік початку роботи: " + strconv.Itoa(int(worker.GetYear())))
	fmt.Println("Введіть місяць початку роботи: " + strconv.Itoa(int(worker.GetMonth())))
}

func PrintWorkers(workers []*Worker) {
	for i := 0; i < len(workers); i++ {
		fmt.Println("Працівник №" + strconv.Itoa(i+1) + ": ")
		fmt.Println("Прізвище та ініціали: " + workers[i].GetName())
		fmt.Println("Місце роботи: " + workers[i].GetWorkPlace())
		fmt.Println("Введіть рік початку роботи: " + strconv.Itoa(int(workers[i].GetYear())))
		fmt.Println("Введіть місяць початку роботи: " + strconv.Itoa(int(workers[i].GetMonth())))
		fmt.Println("")
	}
}

func GetWorkersInfo(workers []*Worker, currentCompany Company) string {
	maxWorkerTotalMoney := workers[0].GetTotalMoney(currentCompany)
	minWorkerTotalMoney := workers[0].GetTotalMoney(currentCompany)
	for i := 1; i < len(workers); i++ {
		if workers[i].GetTotalMoney(currentCompany) > maxWorkerTotalMoney {
			maxWorkerTotalMoney = workers[i].GetTotalMoney(currentCompany)
		} else if workers[i].GetTotalMoney(currentCompany) < minWorkerTotalMoney {
			minWorkerTotalMoney = workers[i].GetTotalMoney(currentCompany)
		}
	}
	return "Найбільша зарплата: " + strconv.FormatFloat(maxWorkerTotalMoney, 'f', 3, 64) + "\nНайменша зарплата: " + strconv.FormatFloat(minWorkerTotalMoney, 'f', 3, 64)
}
