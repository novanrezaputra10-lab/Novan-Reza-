 package main

import "fmt"

type arr [1000000]int

func insertionSort(T *arr, n int) {
	i := 1
	for i <= n-1 {
		j := i
		temp := T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func hitungMedian(T *arr, n int) int {
	var median int
	if n%2 != 0 {
		median = T[n/2]
	} else {
		median = (T[(n/2)-1] + T[n/2]) / 2
	}
	return median
}

func main() {
	var data arr 
	var angka, n int

	fmt.Scan(&angka)

	for angka != -5313 {
		if angka == 0 {
	
			insertionSort(&data, n)
			
			fmt.Println(hitungMedian(&data, n))
		} else {
			
			data[n] = angka
			n = n + 1
		}
		
		fmt.Scan(&angka)
	}
}

