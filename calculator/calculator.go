package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Penggunaan: calculator [angka1] [operator] [angka2]")
		os.Exit(1)
	}

	arg1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator := os.Args[2]
	arg2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Masukan harus berupa angka.")
		os.Exit(1)
	}

	hasil := 0.0

	switch operator {
	case "add":
		hasil = arg1 + arg2
	case "sub":
		hasil = arg1 - arg2
	case "mul":
		hasil = arg1 * arg2
	case "div":
		if arg2 == 0 {
			fmt.Println("Tidak dapat melakukan pembagian oleh nol.")
			os.Exit(1)
		}
		hasil = arg1 / arg2
	default:
		fmt.Println("Operator tidak valid. Gunakan +, -, *, atau /.")
		os.Exit(1)
	}

	fmt.Printf("Hasil: %.2f\n", hasil)
}
