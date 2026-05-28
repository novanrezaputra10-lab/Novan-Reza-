# <h1 align="center">Laporan Praktikum Modul 14 Selection sort </h1>
<p align="center">[Novan Reza Putra] - [109082500102]</p>

## Unguided 

### 1.Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. TentunyaHercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000)yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah.

#### soal1.go

```go
package main
import "fmt"

type arrInt [1000000]int

func selectionAsc(T *arrInt, n int) {
	var idxMin int

	i := 1
	for i <= n - 1 {
		idxMin = i - 1
		j := i
		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j++	
		}
		t := T[idxMin]
		T[idxMin] = T[i - 1]
		T[i - 1] = t
		i++
	}
}

func main() {
	var n, m int
	var a arrInt

	fmt.Scan(&n)

	i := 0
	for i < n {
		fmt.Scan(&m)

		j := 0
		for j < m {
			fmt.Scan(&a[j])
			j++
		}

		selectionAsc(&a, m)

			j = 0
				for j < m {
				fmt.Print(a[j], " ")
				j++
		}
		fmt.Println()

		i++
	}

}
```
### Output Unguided :

##### Output 
![Screenshot Output](https://https://github.com/novanrezaputra10-lab/Novan-Reza-/blob/main/modul14selectionsort/output/soal1.png)

##### Output 
[penjelasan]
Program ini dibuat untuk mengurutkan data bilangan dari kecil ke besar menggunakan metode selection sort. Program akan menerima beberapa kelompok data, kemudian setiap kelompok akan diinput jumlah elemennya beserta angka-angkanya. Setelah itu program mencari nilai terkecil pada array lalu menukarnya ke posisi depan secara bertahap sampai seluruh data terurut. Hasil akhir dari setiap kelompok data yang sudah diurutkan kemudian ditampilkan ke layar secara ascending.


### 2.Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selaludiusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. Format Masukan masih persis sama seperti sebelumnya. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output](https://https://github.com/novanrezaputra10-lab/Novan-Reza-/blob/main/modul14selectionsort/output/soal2.png)

##### Output 
[penjelasan]
Program ini digunakan untuk memisahkan bilangan ganjil dan genap dari beberapa kelompok data yang diinputkan pengguna. Setelah data dipisahkan, bilangan ganjil akan diurutkan dari kecil ke besar menggunakan metode selection sort ascending, sedangkan bilangan genap diurutkan dari besar ke kecil menggunakan selection sort descending. Setelah proses pengurutan selesai, program menampilkan seluruh bilangan ganjil terlebih dahulu kemudian dilanjutkan dengan bilangan genap pada setiap kelompok data.

### 3.Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.

#### soal3.go

```go
package main
import "fmt"

type arr [1000000]int

func insertionSort(T *arr, n int) {
	i := 1
	for i <= n-1 {
		j := i
		temp := T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func hitungMedian(T *arr, n int) int {
	var median int
	if n%2 != 0 {
		median = T[n/2]
	} else {
		median = (T[(n/2)-1] + T[n/2]) / 2
	}
	return median
}

func main() {
	var data arr 
	var angka, n int

	fmt.Scan(&angka)

	for angka != -5313 {
		if angka == 0 {
	
			insertionSort(&data, n)
			
			fmt.Println(hitungMedian(&data, n))
		} else {
			
			data[n] = angka
			n = n + 1
		}
		
		fmt.Scan(&angka)
	}
}


```
### Output Unguided :

##### Output 
![Screenshot Output](https://https://github.com/novanrezaputra10-lab/Novan-Reza-/blob/main/modul14selectionsort/output/soal3.png)


##### Output 
[penjelasan]
Kode program di atas merupakan program sederhana yang digunakan untuk mencari nilai median dari sekumpulan data yang diinputkan pengguna. Program menggunakan metode insertion sort untuk mengurutkan data dari kecil ke besar sebelum median dihitung. Array bertipe integer digunakan sebagai tempat penyimpanan data yang dimasukkan. Prosedur `insertionSort` berfungsi untuk melakukan pengurutan data dengan cara membandingkan elemen saat ini dengan elemen sebelumnya, kemudian memindahkan posisi data sampai urut sesuai ascending. Variabel `i` digunakan sebagai iterasi utama, `j` sebagai iterasi pembanding, dan `temp` sebagai penyimpan sementara data yang sedang dipindahkan. Setelah data terurut, fungsi `hitungMedian` digunakan untuk menentukan nilai median. Jika jumlah data ganjil maka median diambil dari nilai tengah array, sedangkan jika jumlah data genap maka median diperoleh dari rata-rata dua nilai tengah. Pada fungsi utama pengguna diminta memasukkan angka secara terus-menerus hingga memasukkan nilai `-5313` sebagai tanda berhenti. Jika pengguna memasukkan angka `0`, maka program akan mengurutkan data yang sudah tersimpan lalu menampilkan nilai median dari data tersebut. Selain itu, setiap angka selain `0` dan `-5313` akan disimpan ke dalam array sebagai data baru.

