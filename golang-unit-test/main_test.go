package main

import "testing"

func TestHitungHargaTotal(t *testing.T) {
	type args struct {
		hargaItem float64
		ongkir    float64
		qty       int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "failed in harga item",
			args: args{
				hargaItem: 0,
				ongkir:    1,
				qty:       1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failed in harga qty",
			args: args{
				hargaItem: 1,
				ongkir:    1,
				qty:       0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failed in harga ongkir",
			args: args{
				hargaItem: 1,
				ongkir:    0,
				qty:       1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				hargaItem: 10000,
				ongkir:    8000,
				qty:       4,
			},
			want:    54000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HitungHargaTotal(tt.args.hargaItem, tt.args.ongkir, tt.args.qty)
			if (err != nil) != tt.wantErr {
				t.Errorf("HitungHargaTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HitungHargaTotal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPembayaranBarang(t *testing.T) {
	type args struct {
		hargaTotal float64
		metode     string
		cicil      bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "failed in harga total",
			args: args{
				hargaTotal: 0,
				metode:     "cod",
				cicil:      true,
			},
			wantErr: true,
		},
		{
			name: "failed in metode pembayaran",
			args: args{
				hargaTotal: 540000,
				metode:     "ngutang",
				cicil:      true,
			},
			wantErr: true,
		},
		{
			name: "failed in cicil true but metode pembayaran not credit",
			args: args{
				hargaTotal: 540000,
				metode:     "cod",
				cicil:      true,
			},
			wantErr: true,
		},
		{
			name: "failed in cicil true, metode pembayaran credit, but harga total less than 500000",
			args: args{
				hargaTotal: 499999,
				metode:     "credit",
				cicil:      true,
			},
			wantErr: true,
		},
		{
			name: "failed in not cicil, but metode pembayaran credit",
			args: args{
				hargaTotal: 540000,
				metode:     "credit",
				cicil:      false,
			},
			wantErr: true,
		},
		{
			name: "pass in cicil, metode pembayaran credit, and harga total >= 500000",
			args: args{
				hargaTotal: 540000,
				metode:     "credit",
				cicil:      true,
			},
			wantErr: false,
		},
		{
			name: "pass in not cicil, metode pembayaran cod",
			args: args{
				hargaTotal: 10000,
				metode:     "cod",
				cicil:      false,
			},
			wantErr: false,
		},
		{
			name: "pass in not cicil, metode pembayaran transfer",
			args: args{
				hargaTotal: 10000,
				metode:     "transfer",
				cicil:      false,
			},
			wantErr: false,
		},
		{
			name: "pass in not cicil, metode pembayaran gerai",
			args: args{
				hargaTotal: 10000,
				metode:     "gerai",
				cicil:      false,
			},
			wantErr: false,
		},
		{
			name: "pass in not cicil, metode pembayaran debit",
			args: args{
				hargaTotal: 10000,
				metode:     "debit",
				cicil:      false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PembayaranBarang(tt.args.hargaTotal, tt.args.metode, tt.args.cicil)
			if (err != nil) != tt.wantErr {
				t.Errorf("PembayaranBarang() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
