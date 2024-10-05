package main

import (
	"fmt"
	"sync"
)

func sumPart(numbers []int, start, end int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Menandai bahwa goroutine ini selesai
	sum := 0
	for i := start; i < end; i++ {
		sum += numbers[i] // Menjumlahkan bagian dari array ini
	}
	resultChan <- sum // Kirim hasil ke channel
}

func main() {
	numbers := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = i + 2 // Mengisi array dengan angka dari 2 sampai 1001
	}

	totalGoroutines := 7
	partSize := len(numbers) / totalGoroutines
	resultChan := make(chan int, totalGoroutines)
	var wg sync.WaitGroup

	for i := 0; i < totalGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		if i == totalGoroutines-1 {
			end = len(numbers) // Mengatasi sisa, memberikan ke goroutine terakhir
		}
		wg.Add(1)
		go sumPart(numbers, start, end, resultChan, &wg) // Memulai goroutine
	}

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for result := range resultChan {
		println("Penjumlahan:", result)
		totalSum += result // Menjumlahkan semua hasil dari channel
	}

	fmt.Println("The total sum is:", totalSum) // Menampilkan jumlah total
}
