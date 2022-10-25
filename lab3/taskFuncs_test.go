package main

import "testing"

func TestIntRand(t *testing.T) {
	x1 := intRand(2)
	x2 := Min(-1)
	x3 := Min("1")
	res1 := {81, 213}
	res2 := {}
	res3 := {}
	if x1 != res1 {
		t.Errorf("Тест 1 не пройдений! Результат %d, а повинен бути %d", x1, res1)
	}
	if x2 != res2 {
		t.Errorf("Тест 2 не пройдений! Результат %d, а повинен бути %d", x2, res2)
	}
	if x3 != res3 {
		t.Errorf("Тест 3 не пройдений! Результат %d, а повинен бути %d", x3, res3)
	}
}
func TestFloatRand(t *testing.T) {
	x1 := floatRand(2)
	x2 := Min(-1)
	x3 := Min("1")
	res1 := {81.010568, 213.998703}
	res2 := {}
	res3 := {}
	if x1 != res1 {
		t.Errorf("Тест 1 не пройдений! Результат %d, а повинен бути %d", x1, res1)
	}
	if x2 != res2 {
		t.Errorf("Тест 2 не пройдений! Результат %d, а повинен бути %d", x2, res2)
	}
	if x3 != res3 {
		t.Errorf("Тест 3 не пройдений! Результат %d, а повинен бути %d", x3, res3)
	}
}