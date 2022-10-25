package main

import (
	"fmt"
	"math"
)

func intRand(elemCount int) []int {
	const m int = 250
	const a int = 8
	const c int = 65
	var randElems []int = make([]int, elemCount+1)
	randElems[0] = elemCount
	for i := 0; i < elemCount; i++ {
		randElems[i+1] = ((a * (randElems[i])) + c) % m
	}
	return randElems[1:]
}

func statistics(array []int) {
	var uniqueArray []int = unique(array)
	var elemCount []int = count(array, uniqueArray)
	var statisticalProbability []float64 = make([]float64, len(elemCount))
	fmt.Println("Статистична імовірність появи випадкових величин:")
	for i := 0; i < len(elemCount); i++ {
		statisticalProbability[i] = float64(elemCount[i]) / float64(len(array))
		fmt.Printf("%f ", statisticalProbability[i])
	}
	var mathHope float64 = 0
	for i := 0; i < len(elemCount); i++ {
		mathHope += float64(uniqueArray[i]) * statisticalProbability[i]
	}
	fmt.Printf("\nМатематичне сподівання дорівнює %f ", mathHope)
	var dispercion float64 = 0
	for i := 0; i < len(elemCount); i++ {
		dispercion += math.Pow(float64(uniqueArray[i])-mathHope, 2) * statisticalProbability[i]
	}
	fmt.Printf("\nДисперсія дорівнює %f ", dispercion)
	fmt.Printf("\nСередньоквадратичне відхилення дорівнює %f ", math.Sqrt(dispercion))
}

func floatRand(elemCount int) []float64 {
	const m int = 4294967296
	const a int = 22695477
	const c int = 1
	var randElems []float64 = make([]float64, elemCount+1)
	var randElemsTemp []int = make([]int, elemCount+1)
	var intRandElems []int = intRand(elemCount)
	randElemsTemp[0] = 2
	for i := 0; i < elemCount; i++ {
		randElemsTemp[i+1] = ((a * (randElemsTemp[i])) + c) % m
		randElems[i] = float64(intRandElems[i]) + float64(randElemsTemp[i+1])/float64(m)
	}
	return randElems
}
