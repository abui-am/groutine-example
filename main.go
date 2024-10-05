package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Fungsi yang dijalankan oleh setiap goroutine yang mensimulasikan orang bangun
func wakeUp(person string, sleepTime time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(sleepTime)
	fmt.Println(person, "has woken up!")
	fmt.Println(person, "has slept for", sleepTime)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Menetapkan seed untuk nilai acak
	var wg sync.WaitGroup
	people := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	start := time.Now()

	for _, person := range people {
		wg.Add(1) // Menambahkan counter WaitGroup
		// Memulai goroutine dengan waktu tidur acak sebelum bangun, waktu tidur dari 1 sampai 10 detik
		go wakeUp(person, time.Duration(rand.Intn(9)+1)*time.Second, &wg)
	}

	wg.Wait() // Menunggu semua goroutine selesai
	fmt.Println("Everyone has woken up!")
	fmt.Println("Total time spent sleeping:", time.Since(start))
}
