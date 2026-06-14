# <h1 align="center">Laporan Praktikum Modul 14 Insertion Sort-</h1>
<p align="center">[Novan Reza Putra] - [109082500102]</p>

## Soal 1 

### 1. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

#### soal1.go

```go
package main
import "fmt"

type arrint [100000]int

func insert(T *arrint, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1{
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func jarak(T *arrint, n int, x *int)bool{
	*x = T[1] - T[0]
	for i := 2; i < n; i++ {
		if T[i] - T[i-1] != *x {
			return false
		}
	}
	return true
}


func main() {
	var n, beda int
	var x arrint
	var stop bool

	n = 0
	for !stop {
		fmt.Scan(&x[n])
		if x[n] >= 0 {
			n++
		}
		stop = x[n] < 0
	}

	insert(&x, n)
	
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", x[i])
	}

	if jarak(&x, n, &beda) {
		fmt.Printf("\nData berjarak %d", beda)
	} else {
		fmt.Println("\nData berjarak tidak tetap")
	}
}
```
### Output Soal 1 :

##### Output 
![Screenshot Output](https://https://github.com/novanrezaputra10-lab/Novan-Reza-/blob/main/modul14insertionsort/output/soal1.png)
program ini untuk menerima beberapa angka dari input sampai pengguna memasukkan angka negatif sebagai tanda berhenti habis itu semua data masuk, angka-angka tersebut diurutkan pake Insertion Sort dari kecil ke yang paling besar hasil perutan kemudian ditampilkan ke layar setelah itu, program mengecek apakah selisih antara setiap angka yang berurutan nilainya sama atau tidak jika semua selisihnya sama, program akan menampilkan besar selisih tersebut jika tidak sama, program akan menampilkan bahwa jarak antar data tidak tetap. Jadi, program ini berfungsi untuk mengurutkan data sekaligus mengecek apakah data memiliki pola selisih yang sama


### 2. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
const nMax : integer = 7919
type Buku = <
id, judul, penulis, penerbit : string
eksemplar, tahun, rating : integer >
type DaftarBuku = array [ 1..nMax] of Buku
Pustaka : DaftarBuku
nPustaka: integer
Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.
Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua
adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.
#### soal2.go

```go
package main
import "fmt"

const NMAX = 7919
type Buku struct {
	id, judul, penulis, penerbit string
	stok, tahun, rating int
}

type tabBuku [NMAX]Buku
var data tabBuku
var jumlah int
func isiData(n int) {
	jumlah = n

	for i := 0; i < jumlah; i++ {
		fmt.Scan(&data[i].id,
			&data[i].judul,
			&data[i].penulis,
			&data[i].penerbit,
			&data[i].stok,
			&data[i].tahun,
			&data[i].rating)
	}
}

func favorit(n int) {
	idx := 0
	for i := 1; i < n; i++ {
		if data[i].rating > data[idx].rating {
			idx = i
		}
	}
	fmt.Printf("%s, %s, %s, %d\n",
		data[idx].judul,
		data[idx].penulis,
		data[idx].penerbit,
		data[idx].tahun)
}
func urut(n int) {
	var temp Buku
	for i := 1; i < n; i++ {
		temp = data[i]
		j := i - 1
		for j >= 0 && data[j].rating < temp.rating {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}
func top5(n int) {
	batas := 5
	if n < 5 {
		batas = n
	}
	for i := 0; i < batas; i++ {
		fmt.Println(data[i].judul)
	}
}
func cari(n, rating int) {
	kiri := 0
	kanan := n - 1
	ketemu := false
	for kiri <= kanan && !ketemu {
		tengah := (kiri + kanan) / 2

		if data[tengah].rating == rating {
			fmt.Printf("%s, %s, %s, %d, %d, %d\n",
				data[tengah].judul,
				data[tengah].penulis,
				data[tengah].penerbit,
				data[tengah].tahun,
				data[tengah].stok,
				data[tengah].rating)
			ketemu = true
		} else if data[tengah].rating < rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}
func main() {
	var n, rating int
	fmt.Scan(&n)
	isiData(n)
	fmt.Println("Buku Terfavorit:")
	favorit(n)
	urut(n)
	fmt.Println("5 Buku Terbaru dengan Rating Tertinggi:")
	top5(n)
	fmt.Println("Masukkan rating buku yang ingin dicari:")
	fmt.Scan(&rating)
	cari(n, rating)
}
```
### Output Soal 2 :

##### Output 
![Screenshot Output](https://https://github.com/novanrezaputra10-lab/Novan-Reza-/blob/main/modul14insertionsort/output/soal2.png)
[penjelasan]
Program ini digunakan untuk mengelola data buku di perpustakaan. Pengguna memasukkan jumlah buku yang akan didata, kemudian menginput data setiap buku berupa ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Setelah semua data dimasukkan, program akan menampilkan buku dengan rating tertinggi sebagai buku terfavorit, menampilkan lima judul buku dengan rating tertinggi setelah data diurutkan, serta mencari dan menampilkan data buku berdasarkan rating yang dimasukkan pengguna. Jika tidak ada buku dengan rating tersebut, program akan menampilkan pesan bahwa data tidak ditemukan.
