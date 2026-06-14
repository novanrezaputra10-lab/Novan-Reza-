package main
import "fmt"

type arrint [100000]int

func insert(T *arrint, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1{
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func jarak(T *arrint, n int, x *int)bool{
	*x = T[1] - T[0]
	for i := 2; i < n; i++ {
		if T[i] - T[i-1] != *x {
			return false
		}
	}
	return true
}


func main() {
	var n, beda int
	var x arrint
	var stop bool

	n = 0
	for !stop {
		fmt.Scan(&x[n])
		if x[n] >= 0 {
			n++
		}
		stop = x[n] < 0
	}

	insert(&x, n)
	
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", x[i])
	}

	if jarak(&x, n, &beda) {
		fmt.Printf("\nData berjarak %d", beda)
	} else {
		fmt.Println("\nData berjarak tidak tetap")
	}
}