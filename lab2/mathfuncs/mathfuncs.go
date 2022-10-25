package mathfuncs

func Min(x ...float64) float64 {
	min := x[0]
	for _, value := range x {
		if value < min {
			min = value
		}
	}
	return min
}
func Average(x ...float64) float64 {
	sum := 0.0
	for _, value := range x {
		sum += value
	}
	return sum / float64(len(x))
}
func Xn(n, x float64) float64 {
	return n * x
}
