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
