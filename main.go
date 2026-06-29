package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, SonarCloud!")

	result := Add(5, 7)
	fmt.Printf("Hasil penjumlahan: %d\n", result)

	// Pastikan TIDAK ADA string rahasia yang di-hardcode di sini
	// secretToken := os.Getenv("APP_SECRET_TOKEN")
	// if secretToken != "" {
	// 	fmt.Println("Token berhasil dimuat.")
	// }
}
