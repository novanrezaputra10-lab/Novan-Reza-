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