// main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unit-testing/transformer"
)

// main adalah fungsi entry point aplikasi.
// Ini mengilustrasikan bagaimana logika transformasi pesan dapat digunakan
// dalam konteks aplikasi command-line sederhana.
// Fungsi ini bertanggung jawab untuk:
// 1. Membaca input (misalnya dari file atau stdin).
// 2. Memanggil logika bisnis (TransformMessage).
// 3. Menulis output (misalnya ke stdout).
func main() {
	log.Println("Aplikasi Transformer Pesan Go dimulai.")

	// Baca input dari file atau stdin
	var input []byte
	var err error

	if len(os.Args) > 1 {
		// Jika ada argumen baris perintah, asumsikan itu adalah path file input
		filePath := os.Args[1]
		input, err = ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Gagal membaca file input %s: %v", filePath, err)
		}
		log.Printf("Membaca input dari file: %s", filePath)
	} else {
		// Jika tidak ada argumen, baca dari stdin
		log.Println("Membaca input dari stdin. Masukkan JSON dan tekan Ctrl+D (Unix) / Ctrl+Z Enter (Windows) untuk mengakhiri.")
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Gagal membaca dari stdin: %v", err)
		}
	}

	if len(input) == 0 {
		log.Println("Tidak ada input yang diterima. Keluar.")
		return
	}

	// Panggil logika bisnis transformasi pesan
	// Ini adalah pemisahan concern: main.go mengelola I/O dan orkestrasi,
	// sementara transformer.go berisi logika inti.
	transformedOutput, err := transformer.TransformMessage(input)
	if err != nil {
		log.Fatalf("Gagal mentransformasi pesan: %v", err)
	}

	// Tulis output ke stdout
	fmt.Println("\n--- Pesan Asli ---")
	fmt.Println(string(input))
	fmt.Println("\n--- Pesan Ter-transformasi ---")
	fmt.Println(string(transformedOutput))

	log.Println("Aplikasi Transformer Pesan Go selesai.")
}
