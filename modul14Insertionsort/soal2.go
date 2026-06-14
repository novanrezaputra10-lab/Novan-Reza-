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