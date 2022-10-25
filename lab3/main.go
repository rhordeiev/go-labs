package main

import "fmt"

func main() {
	elemCount := 113
	var randElems []int = make([]int, elemCount)
	randElems = intRand(elemCount)
	fmt.Println("Згенерований масив псевдовипадкових цілих чисел:")
	for i := 0; i < elemCount; i++ {
		fmt.Printf("%d ", randElems[i])
	}
	fmt.Println("\nРозраховані статистичні дані масиву:")
	statistics(randElems)
	fmt.Println("\nЗгенерований масив псевдовипадкових дійсних чисел:")
	var randElemsFloat []float64 = make([]float64, elemCount)
	randElemsFloat = floatRand(elemCount)
	for i := 0; i < elemCount; i++ {
		fmt.Printf("%f ", randElemsFloat[i])
	}
}
