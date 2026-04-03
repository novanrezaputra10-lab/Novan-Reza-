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