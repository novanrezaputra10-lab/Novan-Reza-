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