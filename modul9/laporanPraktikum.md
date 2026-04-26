# <h1 align="center">Laporan Praktikum Modul 9 - Array</h1>
<p align="center">Novan Reza Putra - 109082500102</p>

## Unguided 

### 1. [Soal 9a]

```go
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) < float64(c.r)
}

func main() {
	var l1, l2 Lingkaran
	var t Titik
	var dl1, dl2 bool

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	dl1 = didalam(l1, t)
	dl2 = didalam(l2, t)

	if dl1 && dl2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dl1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dl2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

##### Output 
![Screenshot Output Unguided 9a](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul9/output/soal9a.png)

##### Penjelasan
Jadi program ini intinya ngecek posisi sebuah titik, apakah dia masuk ke dalam lingkaran atau nggak. Di sini ada dua lingkaran yang masing-masing punya titik pusat dan jari-jari. Pas dijalankan, kita diminta masukin 3 baris angka, baris pertama buat lingkaran 1 (koordinat pusat dan radius), baris kedua buat lingkaran 2, dan baris ketiga itu titik yang mau dicek. Setelah itu program bakal hitung jarak dari titik tersebut ke pusat masing-masing lingkaran pakai rumus jarak (yang ada akar-akar itu). Kalau jaraknya lebih kecil dari jari-jari berarti titiknya ada di dalam lingkaran, kalau lebih besar berarti di luar. Terakhir, program bakal nampilin hasilnya: kalau masuk ke dua lingkaran tampil “di dalam lingkaran 1 dan 2”, kalau cuma di lingkaran 1 tampil “di dalam lingkaran 1”, kalau cuma di lingkaran 2 tampil “di dalam lingkaran 2”, dan kalau nggak masuk dua-duanya tampil “di luar lingkaran 1 dan 2”. Intinya simpel, cuma hitung jarak, bandingin sama radius, terus tentuin posisinya.

### 2. [Soal 9b]

```go
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
```

##### Output 
![Screenshot Output Unguided 9b](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul9/output/soal9b.png)

##### Penjelasan
Jadi program ini pertama-tama minta kita masukin jumlah elemen array, terus kita isi satu per satu nilai array-nya sesuai jumlah itu. Setelah itu, program langsung nampilin semua isi array. Lanjut, program bakal nampilin elemen yang ada di indeks ganjil aja, terus indeks genap aja (dengan asumsi indeks 0 itu genap). Abis itu, kita diminta masukin nilai x, dan program bakal nampilin elemen array yang indeksnya kelipatan x. Selanjutnya, program minta kita masukin indeks yang mau dihapus, lalu elemen di posisi itu bakal dihapus dan array hasilnya ditampilin lagi. Setelah itu, program ngitung rata-rata dari semua elemen array dan langsung ditampilin. Nggak cuma itu, program juga ngitung standar deviasi (simpangan baku) dari array tersebut. Terakhir, kita diminta masukin satu angka lagi, dan program bakal ngecek berapa kali angka itu muncul di dalam array. Intinya sih program ini kayak ngolah array dari berbagai sisi, mulai dari nampilin, ngapus, sampai ngitung statistiknya.

### 3. [Soal 9c]

```go
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
```

##### Output 
![Screenshot Output Unguided 9c](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul9/output/soal9c.png)

##### Penjelasan
Jadi program ini awalnya kita diminta masukin nama dua klub, yaitu Klub A sama Klub B. Setelah itu program masuk ke perulangan buat nyimulasikan beberapa pertandingan. Di setiap pertandingan, kita diminta masukin skor dari kedua klub. Kalau kita masukin angka negatif di salah satu skor, itu jadi tanda kalau pertandingan udah selesai dan program bakal berhenti. Selama skornya masih valid, program bakal bandingin hasilnya: kalau skor Klub A lebih besar, berarti Klub A menang dan namanya disimpan ke daftar pemenang; kalau Klub B yang lebih besar, maka Klub B yang disimpan; kalau skornya sama, dianggap draw dan nggak disimpan. Setelah semua pertandingan selesai, program bakal nampilin daftar klub yang menang dari setiap pertandingan yang tadi dimasukin. Jadi intinya program ini buat nyatet hasil pertandingan dan nampilin siapa aja yang pernah menang.

### 4. [Soal 9d]

```go
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
```

##### Output 
![Screenshot Output Unguided 9d](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul9/output/soal9b.png)

##### Penjelasan
Program ini buat ngecek kata atau kalimat yang kita masukin itu palindrome atau bukan, sekalian nampilin versi kebaliknya juga. Jadi pas dijalankan, kita disuruh masukin teks dan diakhiri pakai titik. Program bakal baca satu per satu hurufnya, tapi spasi sama enter nggak dihitung, jadi yang masuk ke array cuma hurufnya aja. Setelah itu, teks yang tadi disalin dulu buat disimpen versi aslinya. Terus array yang utama dibalik, jadi huruf depan jadi belakang, dan sebaliknya, lalu hasilnya ditampilin sebagai teks yang sudah dibalik. Nah habis itu baru dicek, apakah teks yang asli sama dengan yang sudah dibalik tadi. Kalau sama berarti palindrome dan hasilnya true, kalau beda berarti false. Jadi intinya program ini cuma baca teks, dibalik, terus dibandingin sama yang awal.