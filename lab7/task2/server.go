package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	http.HandleFunc("/calculateGET", func(w http.ResponseWriter, r *http.Request) {
		enteredString := strings.Join(r.URL.Query()["array[]"], " ")
		sumOddCalculation(enteredString, w)
		multiplicationBetweenZerosCalculation(enteredString, w)
	})
	http.HandleFunc("/calculatePOST", func(w http.ResponseWriter, r *http.Request) {
		enteredString := r.FormValue("array")
		sumOddCalculation(enteredString, w)
		multiplicationBetweenZerosCalculation(enteredString, w)
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}

func sumOddCalculation(enteredString string, w http.ResponseWriter) {
	var values []float64
	words := strings.Fields(enteredString)
	if enteredString == "" {
		fmt.Fprintf(w, `<p>Введіть дані у поле</p>`)
		return
	}
	for i := 0; i < len(words); i++ {
		tempValue, err := strconv.ParseFloat(words[i], 64)
		if err != nil {
			fmt.Fprintf(w, `<p>%s</p>`, err)
			return
		}
		values = append(values, tempValue)
	}
	var sum float64 = 0
	for i := 0; i < len(values); i++ {
		if (i+1)%2 != 0 {
			sum += values[i]
		}
	}
	fmt.Fprintf(w, "Сума непарних елементів: %v", sum)
}

func multiplicationBetweenZerosCalculation(enteredString string, w http.ResponseWriter) {
	var values []float64
	words := strings.Fields(enteredString)
	if enteredString == "" {
		fmt.Fprintf(w, `<p>Введіть дані у поле</p>`)
		return
	}
	for i := 0; i < len(words); i++ {
		tempValue, err := strconv.ParseFloat(words[i], 64)
		if err != nil {
			fmt.Fprintf(w, `<p>%s</p>`, err)
			return
		}
		values = append(values, tempValue)
	}
	var multiplication float64 = 1
	var foundFirstZero bool = false
	var foundLastZero bool = false
	var zeroElementsBetweenZeros bool = false
upperLoop:
	for i := 0; i < len(values); i++ {
		if values[i] == 0 {
			foundFirstZero = true
			for j := i + 1; j < len(values); j++ {
				if values[j] == 0 {
					if j == i+1 {
						zeroElementsBetweenZeros = true
					}
					foundLastZero = true
					break upperLoop
				}
				multiplication *= values[j]
			}
		}
	}
	if foundFirstZero && foundLastZero && !zeroElementsBetweenZeros {
		fmt.Fprintf(w, "\nДобуток елементів між першим та останнім нулем: %v", multiplication)
	} else if !foundFirstZero {
		fmt.Fprintf(w, "\nНульових елементів не знайдено")
	} else if !foundLastZero {
		fmt.Fprintf(w, "\nОстаннього нульового елемента не знайдено")
	} else {
		fmt.Fprintf(w, "\nПоміж нулів немає елементів")
	}
}
