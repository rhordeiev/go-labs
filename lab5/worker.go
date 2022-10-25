package main

import (
	"time"
)

type Worker struct {
	name      string
	year      uint
	month     uint
	workPlace string
}

func NewWorker(enteredName, enteredWorkPlace string, enteredYear, enteredMonth uint) *Worker {
	worker := new(Worker)
	worker.name = enteredName
	worker.workPlace = enteredWorkPlace
	worker.year = enteredYear
	worker.month = enteredMonth
	return worker
}

func (currentWorker *Worker) SetName(enteredName string) {
	currentWorker.name = enteredName
}

func (currentWorker *Worker) SetWorkPlace(enteredWorkPlace string) {
	currentWorker.workPlace = enteredWorkPlace
}

func (currentWorker *Worker) SetYear(enteredYear uint) {
	currentWorker.year = enteredYear
}

func (currentWorker *Worker) SetMonth(enteredMonth uint) {
	currentWorker.month = enteredMonth
}

func (currentWorker Worker) GetName() string {
	return currentWorker.name
}

func (currentWorker Worker) GetWorkPlace() string {
	return currentWorker.workPlace
}

func (currentWorker Worker) GetYear() uint {
	return currentWorker.year
}

func (currentWorker Worker) GetMonth() uint {
	return currentWorker.month
}

func (currentWorker Worker) GetWorkerPosition(currentCompany Company) string {
	return currentCompany.GetPosition()
}

func (currentWorker Worker) GetWorkExpirience() int {
	currentYear := time.Now().Year()
	monthsOnJob := (currentYear - int(currentWorker.year)) * 12
	return monthsOnJob + int(currentWorker.month)
}

func (currentWorker Worker) GetTotalMoney(currentCompany Company) float64 {
	return currentCompany.GetSalary() * float64(currentWorker.GetWorkExpirience())
}
