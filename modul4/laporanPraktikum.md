# <h1 align="center">Laporan Praktikum MODUL 4. PROSEDUR</h1>
<p align="center">Novan Reza Putra - 109082500102</p>

## Unguided 

### 1.Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d.Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n! (n−r)!, sedangkan C(n, r) = n! r!(n−r)!

```go
package main
import "fmt"
func factorial(n int, hasil *int64) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= int64(i)
	}
}
func permutation(n, r int, hasil *int64) {
	var fn, fnr int64
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}
func combination(n, r int, hasil *int64) {
	var fn, fr, fnr int64
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}
func main() {
	var a, b, c, d int
	var x, y int64
	fmt.Scan(&a, &b, &c, &d)
	permutation(a, c, &x)
	combination(a, c, &y)
	fmt.Println(x, y)
	permutation(b, d, &x)
	combination(b, d, &y)
	fmt.Println(x, y)
}
```
##### Output 
![Screenshot Output 1](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul4/output/soal1.png)
Saat program dijalankan, kita memasukin tiga data bertipe string secara berurutan yang disimpan divariabel 1,2,3 terus menukar nilainya dengan cara menyimpan nilai var 1 ke var sementara terus nilai var dua dipindahkan di var 1 nilai var 3 dipindahkan di var 2 terakhir nilai sementara dipindahkan divar 3 jadi geser nilai terus program akan menampilkan isi akhir dari variabel 1,2,3


### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer) Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
```go

package main
import "fmt"

func hitungSkor(waktu1, waktu2, waktu3, waktu4, waktu5, waktu6, waktu7, waktu8 int, soal *int, skor *int) {
    *soal = 0
    *skor = 0
    if waktu1 <= 300 {
        *soal++
        *skor += waktu1
    }
    if waktu2 <= 300 {
        *soal++
        *skor += waktu2
    }
    if waktu3 <= 300 {
        *soal++
        *skor += waktu3
    }
    if waktu4 <= 300 {
        *soal++
        *skor += waktu4
    }
    if waktu5 <= 300 {
        *soal++
        *skor += waktu5
    }
    if waktu6 <= 300 {
        *soal++
        *skor += waktu6
    }
    if waktu7 <= 300 {
        *soal++
        *skor += waktu7
    }
    if waktu8 <= 300 {
        *soal++
        *skor += waktu8
    }
}

func main() {
    var nama, pemenang string
    var soal, skor int
    var maxSoal, minSkor int
    var first = true

    for {
        fmt.Scan(&nama)
        if nama == "Selesai" {
            break
        }

        var w1, w2, w3, w4, w5, w6, w7, w8 int
        fmt.Scan(&w1, &w2, &w3, &w4, &w5, &w6, &w7, &w8)

        hitungSkor(w1, w2, w3, w4, w5, w6, w7, w8, &soal, &skor)

        if first || soal > maxSoal || (soal == maxSoal && skor < minSkor) {
            first = false
            maxSoal = soal
            minSkor = skor
            pemenang = nama
        }
    }

    fmt.Println(pemenang, maxSoal, minSkor)
}

```
##### Output 
![Screenshot Output 2](https://github.com/novanrezaputra10-lab/109082500102_Novan-Reza-Putra/blob/main/modul4/output/soal2.png)
Program untuk nentukan pemenang dari jumlah soal yang berhasil diselesaikan dan total waktu tercepat tiap peserta memasukkan nama lalu 8 waktu pengerjaan soal terus menghitung berapa soal yang selesai hanya yang waktunya <300 detik lalu menjumlahkan total waktu setelah di bandingkanpeserta lain pemenang adalah yang jumlah soalnya paling banyak kalau jumlah soal sama maka dipilih yang total waktunya paling kecil proses ini dilakukan berulang sampai input nama “Selesai” terus program menampilkan nama pemenang beserta jumlah soal yang diselesaikan dan total waktunya.