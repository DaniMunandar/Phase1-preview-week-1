package main

import (
	"fmt"
	"os"
	"strconv"
)

// Struct untuk menyimpan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Data teman-teman di kelas
var temanKelas = []Teman{
	{"Teman 1", "Alamat 1", "Pekerjaan 1", "Alasan 1"},
	{"Teman 2", "Alamat 2", "Pekerjaan 2", "Alasan 2"},
	{"Teman 3", "Alamat 3", "Pekerjaan 3", "Alasan 3"},
	// Tambahkan data teman-teman lain di sini
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Penggunaan: go run biodata.go [nomor absen]")
		os.Exit(1)
	}

	nomorAbsen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka.")
		os.Exit(1)
	}

	if nomorAbsen < 1 || nomorAbsen > len(temanKelas) {
		fmt.Printf("Nomor absen harus antara 1 dan %d\n", len(temanKelas))
		os.Exit(1)
	}

	teman := temanKelas[nomorAbsen-1]
	fmt.Printf("Data Teman (Absen %d):\n", nomorAbsen)
	fmt.Printf("Nama: %s\n", teman.Nama)
	fmt.Printf("Alamat: %s\n", teman.Alamat)
	fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
	fmt.Printf("Alasan memilih kelas Golang: %s\n", teman.Alasan)
}
