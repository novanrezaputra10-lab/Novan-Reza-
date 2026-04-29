# <h1 align="center">Laporan Praktikum Modul 10  </h1>
<p align="center">[Novan Reza Putra] - [109082500102]</p>

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.

Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

#### soal1.go

```go
package main
import "fmt"

func main() {
	var n int
	var berat [1000]float64

	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan berat kelinci ke-", i+1, ": ")
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 0; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println("Kelinci terkecil:", min)
	fmt.Println("Kelinci terbesar:", max)
}
```
### Output Unguided :

##### Output 
![Screenshot Output](https://https://github.com/novanrezaputra10-lab/Novan-Reza-/blob/main/modul10/output/soal1.png)

##### Output 
[penjelasan]
Program ini dipakai untuk mencatat berat anak kelinci, lalu mencari mana yang paling ringan dan paling berat. Pertama, kita diminta memasukkan jumlah kelinci. Setelah itu, kita memasukkan berat tiap kelinci satu per satu, dan semua data tersebut disimpan di dalam array. Lalu program menganggap data pertama sebagai nilai awal untuk yang paling kecil dan paling besar. Setelah itu, semua data dibandingkan satu per satu: kalau ada yang lebih kecil, dijadikan yang terkecil, dan kalau ada yang lebih besar, dijadikan yang terbesar. Di akhir, program menampilkan hasil berat kelinci paling kecil dan paling besar.


### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.

Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2).Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### soal2.go

```go
package main
import "fmt"

func main() {
	var totalIkan int
	var isiWadah int
	var berat [1000]float64

	fmt.Print("Masukan jumlah ikan (x) dan ikan per wadah (y): ")
	fmt.Scan(&totalIkan, &isiWadah)

	for i := 0; i < totalIkan; i++ {
		fmt.Printf("Masukan berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	jumlahWadah := totalIkan / isiWadah
	if totalIkan%isiWadah != 0 {
		jumlahWadah++
	}

	beratWadah := make([]float64, jumlahWadah)

	for i := 0; i < totalIkan; i++ {
		indexWadah := i / isiWadah
		beratWadah[indexWadah] += berat[i]
	}

	fmt.Print("Total berat per wadah: ")
	totalSemua := 0.0
	for i, isi := range beratWadah {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", isi)
		totalSemua += isi
	}
	fmt.Println()

	rataRata := totalSemua / float64(jumlahWadah)
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output](https://https://github.com/novanrezaputra10-lab/Novan-Reza-/blob/main/modul10/output/soal2.png)

##### Output 
[penjelasan]
Program ini dibuat untuk menghitung total berat ikan di setiap wadah dan juga rata-ratanya, kita diminta memasukkan dua angka, yaitu jumlah ikan (totalIkan) dan berapa ikan yang dimasukkan ke tiap wadah (isiWadah). Setelah itu, kita memasukkan berat masing-masing ikan satu per satu, lalu semua disimpan ke dalam array, program menghitung berapa jumlah wadah yang dibutuhkan. Kalau ikan tidak habis dibagi rata, maka akan ditambah satu wadah lagi untuk sisa ikan itu dibuat tempat untuk menyimpan total berat tiap wadah. Program lalu mengelompokkan ikan ke dalam wadah sesuai urutan input. Misalnya tiap wadah isi 2 ikan, maka ikan ke-1 dan ke-2 masuk wadah pertama, ke-3 dan ke-4 ke wadah kedua, dan seterusnya. Berat ikan dalam wadah yang sama dijumlahkan program menampilkan total berat di setiap wadah dalam satu baris. Lalu dihitung juga jumlah semua berat tadi untuk mencari rata-rata berat per wadah, dan hasilnya ditampilkan di baris berikutnya

### 3. SPos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

#### soal3.go

```go
package main
import "fmt"

type arrbalita [100]float64

func hitung(arrberat *arrbalita, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&arrberat[i])
	}

	minBerat := arrberat[0]
	maxBerat := arrberat[0]

	for i := 1; i < jumlah; i++ {
		if arrberat[i] < minBerat {
			minBerat = arrberat[i]
		} else if arrberat[i] > maxBerat {
			maxBerat = arrberat[i]
		}
	}
	fmt.Printf("Berat balita minimum: %.2f kg\n", minBerat)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", maxBerat)
}

func hitungRata(arrberat *arrbalita, jumlah int) float64 {
	totalBerat := 0.0
	for i := 0; i < jumlah; i++ {
		totalBerat += arrberat[i]
	}
	return totalBerat / float64(jumlah)
}

func main() {
	var balita arrbalita
	var jumlah int

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&jumlah)

	hitung(&balita, jumlah)
	rataRata := hitungRata(&balita, jumlah)

	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output](https://https://github.com/novanrezaputra10-lab/Novan-Reza-/blob/main/modul10/output/soal3.png)


##### Output 
[penjelasan]
Program digunakan untuk mencatat berat balita dan menghitung berat minimum, maksimum, serta rata-ratanya. Pertama, pengguna memasukkan jumlah data balita, kemudian program meminta input berat masing-masing balita satu per satu dan menyimpannya ke dalam array. Setelah semua data dimasukkan, program menentukan nilai awal dari data pertama sebagai berat terkecil dan terbesar, lalu membandingkan dengan data lainnya untuk mendapatkan nilai minimum dan maksimum yang sebenarnya. Hasil tersebut kemudian ditampilkan. Selanjutnya, program menghitung rata-rata dengan menjumlahkan seluruh data berat dan membaginya dengan jumlah balita. Terakhir, nilai rata-rata ditampilkan sebagai hasil akhir.
