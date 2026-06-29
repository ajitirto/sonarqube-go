package main

import (
	"fmt"
	"os"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, SonarCloud!")

	result := Add(5, 7)
	fmt.Printf("Hasil penjumlahan: %d\n", result)

	secretToken := os.Getenv("APP_SECRET_TOKEN")
	if secretToken != "" {
		fmt.Println("Token berhasil dimuat dari environment variable.")
	} else {
		fmt.Println("Peringatan: APP_SECRET_TOKEN tidak diatur.")
	}
}
