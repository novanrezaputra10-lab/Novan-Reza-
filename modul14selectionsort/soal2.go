package main

import "fmt"

type arrInt [1000000]int

func selectionSortAsc(T *arrInt, n int) {
	var t, i, j, idx_min int

	for i = 1; i <= n-1; {
		idx_min = i - 1

		for j = i; j < n; {
			if T[idx_min] > T[j] {
				idx_min = j
			}

			j++
		}

		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i++
	}
}

func selectionSortDesc(T *arrInt, n int) {
	var t, i, j, idx_max int

	for i = 1; i <= n-1; {
		idx_max = i - 1

		for j = i; j < n; {
			if T[idx_max] < T[j] {
				idx_max = j
			}

			j++
		}

		t = T[idx_max]
		T[idx_max] = T[i-1]
		T[i-1] = t
		i++
	}
}

func main() {
	var n, i, d, jumlah, angka int
	var ganjil, genap arrInt
	var jGanjil, jGenap int

	fmt.Scan(&n)

	for d = 0; d < n; {
		fmt.Scan(&jumlah)

		jGanjil = 0
		jGenap = 0

		for i = 0; i < jumlah; {
			fmt.Scan(&angka)

			if angka%2 != 0 {
				ganjil[jGanjil] = angka
				jGanjil = jGanjil + 1
			} else {
				genap[jGenap] = angka
				jGenap = jGenap + 1
			}

			i++
		}

		selectionSortAsc(&ganjil, jGanjil)
		selectionSortDesc(&genap, jGenap)

		for i = 0; i < jGanjil; {
			if i > 0 {
				fmt.Print(" ")
			}

			fmt.Print(ganjil[i])
			i++
		}

		if jGanjil > 0 && jGenap > 0 {
			fmt.Print(" ")
		}

		for i = 0; i < jGenap; {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(genap[i])

			i++
		}
		fmt.Println()

		d++
	}
}