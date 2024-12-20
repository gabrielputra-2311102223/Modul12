package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk menghitung median dari slice yang telah diurutkan
func median(sortedData []int) int {
	n := len(sortedData)
	if n%2 == 1 {
		return sortedData[n/2]
	}
	return (sortedData[n/2-1] + sortedData[n/2]) / 2
}

// Fungsi untuk melakukan insertion sort
func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var data []int

	fmt.Println("Masukkan bilangan bulat, pisahkan dengan spasi. Ketik -5313 untuk mengakhiri input:")

	for {
		fmt.Print("Masukkan bilangan: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		strNumbers := strings.Fields(text)
		for _, strNum := range strNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Println("Input tidak valid, silakan masukkan bilangan bulat.")
				continue
			}

			if num == -5313 {
				return
			}

			if num == 0 {
				// Urutkan data yang ada
				insertionSort(data)
				// Cetak median
				fmt.Println(median(data))
			} else {
				data = append(data, num)
			}
		}
	}
}
