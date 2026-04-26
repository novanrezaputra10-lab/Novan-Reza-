package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var k rune

	*n = 0

	for *n < NMAX {
		fmt.Scanf("%c", &k)

		if k == '.' {
			break
		}

		if k != '\n' && k != ' ' {
			t[*n] = k
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}

	fmt.Println()
}

func balikkanArray(t *tabel, n int) {
	var i int

	for i = 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrome(t tabel, n int) bool {
	var temp tabel
	var i int

	for i = 0; i < n; i++ {
		temp[i] = t[i]
	}

	balikkanArray(&temp, n)

	for i = 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}

	return true
}

func main() {
	var tab, asli tabel
	var m, i int

	fmt.Print("teks: ")
	isiArray(&tab, &m)

	for i = 0; i < m; i++ {
		asli[i] = tab[i]
	}

	balikkanArray(&tab, m)

	fmt.Print("Reverse Teks: ")
	cetakArray(tab, m)

	fmt.Print("Palindrome? ")
	if palindrome(asli, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
