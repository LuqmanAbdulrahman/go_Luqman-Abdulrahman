package main

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()
}

// Menggunakan huruf kapital pada awal kata untuk nama struct dan method.
// Menggunakan spasi antara nama variabel dan tipe datanya.
// Menggunakan tanda kurung kurawal pada setiap blok fungsi dan loop.
// Memberikan nama variabel yang lebih jelas dan deskriptif.
// Menggunakan operator += pada fungsi TambahKecepatan.
// Menggunakan nama mobilCepat dan mobilLamban untuk variabel mobil.
