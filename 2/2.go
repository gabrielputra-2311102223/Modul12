package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Fungsi untuk memisahkan bilangan ganjil dan genap
func separateOddEven(numbers []int) ([]int, []int) {
	var odds, evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	return odds, evens
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan baris bilangan, pisahkan dengan spasi. Tekan Enter untuk setiap baris. Ketik 'selesai' untuk mengakhiri input:")

	var input [][]int

	for {
		fmt.Print("Masukkan bilangan: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if strings.ToLower(text) == "selesai" {
			break
		}

		strNumbers := strings.Fields(text)
		var numbers []int

		for _, strNum := range strNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Println("Input tidak valid, silakan masukkan bilangan bulat.")
				continue
			}
			numbers = append(numbers, num)
		}

		if len(numbers) > 0 {
			input = append(input, numbers)
		}
	}

	// Proses setiap baris input
	for _, numbers := range input {
		odds, evens := separateOddEven(numbers)

		// Urutkan bilangan ganjil secara menurun
		sort.Slice(odds, func(i, j int) bool {
			return odds[i] > odds[j]
		})

		// Urutkan bilangan genap secara menaik
		sort.Ints(evens)

		// Gabungkan hasil
		result := append(odds, evens...)

		// Cetak hasil
		fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
	}
}
