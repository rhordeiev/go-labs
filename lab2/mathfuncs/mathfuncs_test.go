package mathfuncs

import "testing"

func TestMin(t *testing.T) {
	x1 := Min(1, 2, 3)
	x2 := Min(-1, -10, 13)
	x3 := Min(0.1, 0.2, -0.3)
	res1 := 1.0
	res2 := -10.0
	res3 := -0.3
	if x1 != res1 {
		t.Errorf("Тест 1 не пройдений! Результат %f, а повинен бути %f", x1, res1)
	}
	if x2 != res2 {
		t.Errorf("Тест 2 не пройдений! Результат %f, а повинен бути %f", x2, res2)
	}
	if x3 != res3 {
		t.Errorf("Тест 3 не пройдений! Результат %f, а повинен бути %f", x3, res3)
	}
}
func TestAverage(t *testing.T) {
	x1 := Average(3, 4, 5)
	x2 := Average(-7, -11, 13)
	x3 := Average(7.5, 1.3, -12.7)
	res1 := 4.0
	res2 := -1.666667
	res3 := -1.3
	if x1 != res1 {
		t.Errorf("Тест 1 не пройдений! Результат %f, а повинен бути %f", x1, res1)
	}
	if x2 != res2 {
		t.Errorf("Тест 2 не пройдений! Результат %f, а повинен бути %f", x2, res2)
	}
	if x3 != res3 {
		t.Errorf("Тест 3 не пройдений! Результат %f, а повинен бути %f", x3, res3)
	}
}
func TestXn(t *testing.T) {
	x1 := Xn(3, 8)
	x2 := Xn(-9, -14)
	x3 := Xn(-7.3, 8.9)
	res1 := 21.0
	res2 := 126.0
	res3 := -64.970000
	if x1 != res1 {
		t.Errorf("Тест 1 не пройдений! Результат %f, а повинен бути %f", x1, res1)
	}
	if x2 != res2 {
		t.Errorf("Тест 2 не пройдений! Результат %f, а повинен бути %f", x2, res2)
	}
	if x3 != res3 {
		t.Errorf("Тест 3 не пройдений! Результат %f, а повинен бути %f", x3, res3)
	}
}
