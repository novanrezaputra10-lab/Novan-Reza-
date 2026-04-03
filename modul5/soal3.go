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