# <h1 align="center">Laporan Praktikum Modul 3 - FUNGSI </h1>
<p align="center">Novan Reza Putra - 109082500102</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n! (n−r)! , sedangkan C(n, r) = n! r!(n−r)!
```go
package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output 1](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul3/satu.png)
[penjelasan] : Program ini buat menghitung permutasi (urutan penting) dan kombinasi (urutan tidak penting) dari dua pasangan angka yang dimasukkan, dengan bantuan perhitungan faktorial.


### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x 2 , g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!

```go

package main

import "fmt"

func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	hasil1 := f(g(h(a)))
	hasil2 := g(h(f(b)))
	hasil3 := h(f(g(c)))

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}
```
### Output Unguided :

##### Output 
![Screenshot Output 2 ](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul3/output/dua.png)
[penjelasan] : Program ini menghitung hasil dari gabungan beberapa fungsi (f, g, h) yang dimasukkan ke dalam satu sama lain (komposisi fungsi) untuk tiga angka input, lalu menampilkan hasilnya.

### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lijngkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)

	in1 := didalam(cx1, cy1, r1, x, y)
	in2 := didalam(cx2, cy2, r2, x, y)

	if in1 && in2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}
	
```
### Output Unguided :

##### Output 
![Screenshot Output 3](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul3/output/tiga.png)
[penjelasan] Program ini dibuat untuk menentukan lokasi titik sembarang.Di program ini terdapat 3 func, Pertama untuk menghitung jarak titik sembarang dari titik pusat,kedua untuk membandingkan jarak titik sembarang terhadap jari" lingkaran, dan yang terakhir main untuk membaca inputan ser dan menentukan output yang akan dikeluarkan program