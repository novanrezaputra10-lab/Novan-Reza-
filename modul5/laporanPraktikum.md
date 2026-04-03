# <h1 align="center"> Laporan Praktikum MODUL 5. REKURSIF </h1>
<p align="center">Novan Reza Putra - 109082500102</p>

## Unguided 

### 1.Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum  dapatdiformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

```go
package main
import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)
	fmt.Println("Deret Fibonacci:")


	for i := 0; i <= n; i++ {
		fmt.Print(fib(i), " ")
	}
}
```
##### Output 
![Screenshot Output 1](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul5/output/soal1.png)
Program ini menampilkan deret Fibonacci sampai n dengan menggunakan fungsi rekursif di mana setiap angka adalah hasil penjumlahan dua angka sebelumnya if n < 2
kalau n = 0 atau 1 langsung balikin nilainya return fib(n-1) + fib(n-2) kalau n lebih dari 1 ambil dari dua angka sebelumnya lalu dijumlahin fungsi fib dipakai buat ngitung fib secara rekursif main program minta input n lalu pakai perulangan untuk manggil fungsi fib dari 0 sampai n dan menampilkan hasilnya


### 2.Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.
```go
package main
import "fmt"

func bintang(b int) {
	if b == 0 {
		return
	}
	bintang(b - 1)
	for i := 0; i < b; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
func main() {
	var b int
	fmt.Print("Masukkan junlah pola bintang: ")
	fmt.Scan(&b)
	bintang(b)
}
```
##### Output 
![Screenshot Output 2](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul5/output/soal2.png)
    Program ini digunakan untuk menampilkan pola bintang berbentuk segitiga menggunakan rekursif user memasukkan jumlah baris (b), lalu program akan menampilkan bintang mulai dari 1 sampai b baris func bintang(b int) fungsi ini buat cetak pola bintang if b == 0 { return } kalau sudah 0 → berhenti (biar ga loop terus) bintang(b - 1) fungsi dipanggil dulu dengan nilai lebih kecil jadi dia turun sampai 0 dulu cetak * sebanyak yang dimasukan misal 3 jadi hasilnya naik   

    
### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

```go

package main
import "fmt"

func faktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(n, i+1)
}
func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}
	
```

##### Output 
![Screenshot Output 3](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul5/output/soal3.png)
Program untuk menampilkan faktor dari suatu bilangan N menggunakan rekursif user memasukkan angka n program akan mengecek semua angka dari 1 sampai n jika suatu angka bisa membagi n tanpa sisa maka angka tersebut adalah faktor dan akan ditampilkan if i > n return kalau sudah lewat n  berhenti if n%i == 0 fmt.Print(i, " ") kalau n habis dibagi i → berarti i adalah faktor faktor(n, i+1) ini akan cek angka berikutnya fmt.Scan(&n) faktor(n, 1) input n lalu mulai cek dari 1 misal masukin 5 keluarnya 1 5 begitulah selesai