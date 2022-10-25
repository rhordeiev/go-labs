package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	http.HandleFunc("/calculateGET", func(w http.ResponseWriter, r *http.Request) {
		a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
		b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
		c, err := strconv.ParseFloat(r.URL.Query().Get("c"), 64)
		if err != nil {
			fmt.Fprintf(w, `<p>%s</p>`, err)
			return
		}
		quadraticEquationCalculation(a, b, c, w)
	})
	http.HandleFunc("/calculatePOST", func(w http.ResponseWriter, r *http.Request) {
		a, err := strconv.ParseFloat(r.FormValue("a"), 64)
		b, err := strconv.ParseFloat(r.FormValue("b"), 64)
		c, err := strconv.ParseFloat(r.FormValue("c"), 64)
		if err != nil {
			fmt.Fprintf(w, `<p>%s</p>`, err)
			return
		}
		quadraticEquationCalculation(a, b, c, w)
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}

func quadraticEquationCalculation(a, b, c float64, w http.ResponseWriter) {
	var res []float64
	if a == 0 {
		fmt.Fprintf(w, `<p>Перший коефіцієнт не може бути нулем</p>`)
		return
	}
	D := b*b - 4*a*c
	if D < 0 {
		fmt.Fprintf(w, `<p>Дискримінант не може бути менше нуля</p>`)
		return
	} else if D == 0 {
		res = append(res, -b+math.Sqrt(D)/(2*a))
		fmt.Fprintf(w, "x1: %v", res[0])
	} else {
		res = append(res, (-b-math.Sqrt(D))/(2*a))
		res = append(res, (-b+math.Sqrt(D))/(2*a))
		fmt.Fprintf(w, "x1: %v; x2: %v", res[0], res[1])
	}
}
