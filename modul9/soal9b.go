package main

import (
	"fmt"
	"math"
)

func showOddIndex(arr []int) {
	var i int

	for i = 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
}

func showEvenIndex(arr []int) {
	var i int

	for i = 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
}

func showMultipleIndex(arr []int) {
	var x, i int

	fmt.Print("input x: ")
	fmt.Scan(&x)

	for i = 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
}

func deleteIndex(arr []int) []int {
	var index int

	fmt.Print("input index array: ")
	fmt.Scan(&index)

	return append(arr[:index], arr[index+1:]...)
}

func showMean(arr []int) float64 {
	var r, i int

	for i = 0; i < len(arr); i++ {
		r += arr[i]
	}
	return float64(r) / float64(len(arr))
}

func showStdDev(arr []int) float64 {
	var i int
	var m, s, jk float64

	m = showMean(arr)

	for i = 0; i < len(arr); i++ {
		s = float64(arr[i]) - m
		jk += s * s
	}

	return math.Sqrt(jk / float64(len(arr)))
}

func showFrequency(arr []int) int {
	var c, i, fq int

	fmt.Print("input c: ")
	fmt.Scan(&c)

	for i = 0; i < len(arr); i++ {
		if arr[i] == c {
			fq++
		}
	}

	return fq
}

func main() {
	var n, i int
	var arr []int

	fmt.Print("input jumlah array: ")
	fmt.Scan(&n)

	arr = make([]int, n)

	for i = 0; i < n; i++ {
		fmt.Printf("input elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println()
	fmt.Println("a. Menampilkan keseluruhan isi dari array.")
	fmt.Println(arr)

	fmt.Println()
	fmt.Println("b. Menampilkan elemen-elemen array dengan indeks ganjil saja.")
	showOddIndex(arr)

	fmt.Println()
	fmt.Println("c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).")
	showEvenIndex(arr)

	fmt.Println()
	fmt.Println("d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.")
	showMultipleIndex(arr)

	fmt.Println()
	fmt.Println("e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil")
	arr = deleteIndex(arr)
	fmt.Println(arr)

	fmt.Println()
	fmt.Println("f. Menampilkan rata-rata dari bilangan yang ada di dalam array.")
	fmt.Println(showMean(arr))

	fmt.Println()
	fmt.Println("g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.")
	fmt.Println(showStdDev(arr))

	fmt.Println()
	fmt.Println("h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.")
	fmt.Println(showFrequency(arr))
}