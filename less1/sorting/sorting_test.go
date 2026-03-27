package main

import "testing"

func TestBubbleSort(t *testing.T) {

	arr := []int{5, 3, 1}
	expected := []int{1, 3, 5}

	BubbleSort(arr)

	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("ошибка: ожидали %v, получили %v", expected, arr)
		}
	}
}
