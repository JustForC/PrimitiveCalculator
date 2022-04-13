package main

import (
	"fmt"

	"github.com/JustForC/PrimitiveCalculator/calculate"
	"github.com/JustForC/PrimitiveCalculator/model"
)

func main() {
	var number model.Number
	var symbol string

	for i := 0; i < 2; {
		if i == 0 {
			fmt.Println("Masukan Angka Pertama")
			fmt.Scanln(&number.NumberOne)

			fmt.Println("Jenis Kalkulasi (Symbol)")
			fmt.Scanln(&symbol)

			fmt.Println("Masukan Angka Kedua")
			fmt.Scanln(&number.NumberTwo)

			calculate.CheckSymbol(symbol, &number)
		}
		if i == 1 {
			number.NumberOne = number.Calculation
			fmt.Println("Jenis Kalkulasi (Symbol)")
			fmt.Scanln(&symbol)

			fmt.Println("Masukan Angka Kedua")
			fmt.Scanln(&number.NumberTwo)

			calculate.CheckSymbol(symbol, &number)
		}

		fmt.Println("Hasil :", number.Calculation)

		fmt.Println("Mulai Ulang Kalkulator [0]")
		fmt.Println("Kalkulasikan Hasil Kembali [1]")
		fmt.Println("Akhiri Program [2]")
		fmt.Scanln(&i)
	}

}
