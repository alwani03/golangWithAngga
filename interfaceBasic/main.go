package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	luasPersegi := Persegi{Sisi: 9}
	hasilLuas := seberapaLuas(luasPersegi)

	fmt.Println("Luas dari Persegi ialah", hasilLuas)

	luasPersegiPanjang := PersegiPanjang{Panjang: 7, Lebar: 7}
	hasilLuas2 := seberapaLuas(luasPersegiPanjang)

	fmt.Println("Luas dari Persegi ialah", hasilLuas2)
}

func seberapaLuas(bangunLuas BangunDatar) int {
	return bangunLuas.HitungLuas()
}
