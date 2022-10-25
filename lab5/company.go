package main

type Company struct {
	name     string
	position string
	salary   float64
}

func NewCompany(enteredName, enteredPosition string, enteredSalary float64) *Company {
	company := new(Company)
	company.name = enteredName
	company.position = enteredPosition
	company.salary = enteredSalary
	return company
}

func (currentCompany *Company) SetName(enteredName string) {
	currentCompany.name = enteredName
}

func (currentCompany *Company) SetPosition(enteredPosition string) {
	currentCompany.position = enteredPosition
}

func (currentCompany *Company) SetSalary(enteredSalary float64) {
	currentCompany.salary = enteredSalary
}

func (currentCompany Company) GetName() string {
	return currentCompany.name
}

func (currentCompany Company) GetPosition() string {
	return currentCompany.position
}

func (currentCompany Company) GetSalary() float64 {
	return currentCompany.salary
}

func (currentCompany Company) GetWorkersInCompanyCount(workers []Worker) uint {
	workersCount := 0
	for i := 0; i < len(workers); i++ {
		if workers[i].workPlace == currentCompany.name {
			workersCount++
		}
	}
	return uint(workersCount)
}

func (currentCompany Company) GetWorkerNameWhoEarnedTheMost(workers []Worker) string {
	maxWorkerTotalMoney := workers[0].GetTotalMoney(currentCompany)
	maxWorkerName := workers[0].GetName()
	for i := 1; i < len(workers); i++ {
		if workers[i].GetTotalMoney(currentCompany) > maxWorkerTotalMoney {
			maxWorkerTotalMoney = workers[i].GetTotalMoney(currentCompany)
			maxWorkerName = workers[i].GetName()
		}
	}
	return maxWorkerName
}

func (currentCompany Company) GetAverageWorkingPeriodOfTime(workers []Worker) float64 {
	workingPeriodOfTimeSum := 0
	for i := 0; i < len(workers); i++ {
		workingPeriodOfTimeSum += workers[i].GetWorkExpirience()
	}
	return float64(workingPeriodOfTimeSum) / float64(len(workers))
}
