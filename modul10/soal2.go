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