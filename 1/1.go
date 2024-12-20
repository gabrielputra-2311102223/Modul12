package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Print("Masukkan banyaknya daerah: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan banyaknya rumah kerabat di daerah %d: ", i+1)
		fmt.Scanln(&m)

		var rumahKerabat []int
		fmt.Printf("Masukkan nomor rumah kerabat di daerah %d (dipisahkan spasi): ", i+1)
		for j := 0; j < m; j++ {
			var rumah int
			fmt.Scan(&rumah)
			rumahKerabat = append(rumahKerabat, rumah)
		}

		sort.Ints(rumahKerabat)
		fmt.Printf("Keluaran daerah %d: ", i+1)
		for _, rumah := range rumahKerabat {
			fmt.Printf("%d ", rumah)
		}
		fmt.Println()
	}
}
