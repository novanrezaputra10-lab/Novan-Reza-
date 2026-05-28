package main

import "fmt"

type arrInt [1000000]int

func selectionAsc(T *arrInt, n int) {
	var idxMin int

	i := 1
	for i <= n - 1 {
		idxMin = i - 1
		j := i
		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j++	
		}
		t := T[idxMin]
		T[idxMin] = T[i - 1]
		T[i - 1] = t
		i++
	}
}

func main() {
	var n, m int
	var a arrInt

	fmt.Scan(&n)

	i := 0
	for i < n {
		fmt.Scan(&m)

		j := 0
		for j < m {
			fmt.Scan(&a[j])
			j++
		}

		selectionAsc(&a, m)

			j = 0
				for j < m {
				fmt.Print(a[j], " ")
				j++
		}
		fmt.Println()

		i++
	}

}