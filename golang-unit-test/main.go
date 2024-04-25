package main

import (
	"errors"
	"fmt"
)

const (
	tax = 10
	app = 2000
)

func main() {
	fmt.Println("Total Harga?")
	fmt.Println(HitungHargaTotal(10000, 8000, 4))
	fmt.Println("Pembayaran Berhasil?")
	fmt.Println(PembayaranBarang(80000, "ngutang", false))
}

func HitungHargaTotal(hargaItem, ongkir float64, qty int) (float64, error) {
	if hargaItem <= 0 {
		return 0, errors.New("harga barang tidak boleh nol")
	}

	if qty <= 0 {
		return 0, errors.New("jumlah barang tidak boleh nol")
	}

	hargaAkhirItem := hargaItem * float64(qty)

	if ongkir <= 0 {
		return 0, errors.New("harga ongkir tidak boleh nol")
	}

	hargaSetelahOngkir := hargaAkhirItem + ongkir

	pajak := hargaAkhirItem * tax / 100

	total := hargaSetelahOngkir + pajak + app

	return total, nil
}

func PembayaranBarang(hargaTotal float64, metode string, cicil bool) error {
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	mets := []string{"cod", "transfer", "debit", "credit", "gerai"}
	valid := false
	for _, met := range mets {
		if metode == met {
			valid = true
		}
	}
	if !valid {
		return errors.New("metode tidak dikenali")
	}

	if cicil {
		if metode != "credit" || hargaTotal < 500000 {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	} else {
		if metode == "credit" {
			return errors.New("credit harus dicicil")
		}
	}

	return nil
}
