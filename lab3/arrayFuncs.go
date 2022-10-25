package main

func unique(array []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range array {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func count(array, uniqueArray []int) []int {
	var list []int = make([]int, len(uniqueArray))
	for i := 0; i < len(uniqueArray); i++ {
		elemCount := 0
		for j := 0; j < len(array); j++ {
			if uniqueArray[i] == array[j] {
				elemCount++
			}
		}
		list[i] = elemCount
	}
	return list
}
