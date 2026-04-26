package main

import (
	"fmt"
)

func main() {
	var ka, kb, p string
	var hasil []string
	var i int
	var pk int

	pk = 1

	fmt.Print("Klub A: ")
	fmt.Scan(&ka)
	fmt.Print("Klub B: ")
	fmt.Scan(&kb)

	for {
		var sa, sb int

		fmt.Printf("Pertandingan %d : ", pk)
		fmt.Scan(&sa, &sb)

		if sa < 0 || sb < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		if sa > sb {
			fmt.Printf("Hasil %d : %s\n", pk, ka)
			hasil = append(hasil, ka)
		} else if sb > sa {
			fmt.Printf("Hasil %d : %s\n", pk, kb)
			hasil = append(hasil, kb)
		} else {
			fmt.Printf("Hasil %d : Draw\n", pk)
		}

		pk++
	}

	fmt.Println("\nDaftar klub yg menang:")
	for i, p = range hasil {
		if p != "" {
			fmt.Printf("%d. %s\n", i+1, p)
		}
	}
}
