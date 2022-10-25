package main

func main() {
	if workers, err := ReadWorkerArray(); err != nil {
		panic(err.Error())
	} else {
		PrintWorkers(workers)
	}
}
