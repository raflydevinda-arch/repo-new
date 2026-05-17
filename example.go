package main

import "fmt"

const MAX int = 100

type Peserta struct {
	id          int
	nama        string
	bidangMinat string
	tanggal     string
	aktif       bool
}

var dataPeserta [MAX]Peserta
var jumlah int

func tambahPeserta() {
	var p Peserta

	fmt.Println("=== Tambah Peserta ===")

	fmt.Print("ID Pendaftaran : ")
	fmt.Scan(&p.id)

	fmt.Print("Nama Lengkap   : ")
	fmt.Scan(&p.nama)

	fmt.Print("Bidang Minat   : ")
	fmt.Scan(&p.bidangMinat)

	fmt.Print("Tanggal Daftar : ")
	fmt.Scan(&p.tanggal)

	p.aktif = true

	dataPeserta[jumlah] = p
	jumlah = jumlah + 1

	fmt.Println("Peserta berhasil ditambahkan")
}

func tampilPeserta() {
	var i int

	fmt.Println("=== Data Peserta ===")

	if jumlah == 0 {

		fmt.Println("Data kosong")

	} else {

		for i = 0; i < jumlah; i++ {

			fmt.Println("-------------------")
			fmt.Println("ID       :", dataPeserta[i].id)
			fmt.Println("Nama     :", dataPeserta[i].nama)
			fmt.Println("Minat    :", dataPeserta[i].bidangMinat)
			fmt.Println("Tanggal  :", dataPeserta[i].tanggal)
			fmt.Println("Aktif    :", dataPeserta[i].aktif)
		}
	}
}

func ubahPeserta() {
	var idCari int
	var i int
	var ketemu bool

	fmt.Println("=== Ubah Peserta ===")

	fmt.Print("Masukkan ID : ")
	fmt.Scan(&idCari)

	ketemu = false

	for i = 0; i < jumlah; i++ {

		if dataPeserta[i].id == idCari {

			fmt.Print("Nama Baru : ")
			fmt.Scan(&dataPeserta[i].nama)

			fmt.Print("Minat Baru : ")
			fmt.Scan(&dataPeserta[i].bidangMinat)

			fmt.Print("Tanggal Baru : ")
			fmt.Scan(&dataPeserta[i].tanggal)

			ketemu = true
		}
	}

	if ketemu == true {

		fmt.Println("Data berhasil diubah")

	} else {

		fmt.Println("ID tidak ditemukan")
	}
}

func hapusPeserta() {
	var idCari int
	var i int
	var j int
	var ketemu bool

	fmt.Println("=== Hapus Peserta ===")

	fmt.Print("Masukkan ID : ")
	fmt.Scan(&idCari)

	ketemu = false

	for i = 0; i < jumlah; i++ {

		if dataPeserta[i].id == idCari {

			for j = i; j < jumlah-1; j++ {

				dataPeserta[j] = dataPeserta[j+1]
			}

			jumlah = jumlah - 1
			ketemu = true
		}
	}

	if ketemu == true {

		fmt.Println("Data berhasil dihapus")

	} else {

		fmt.Println("ID tidak ditemukan")
	}
}

func sequentialSearchNama() {
	var cari string
	var i int
	var ketemu bool

	fmt.Println("=== Sequential Search ===")

	fmt.Print("Cari Nama : ")
	fmt.Scan(&cari)

	ketemu = false

	for i = 0; i < jumlah; i++ {

		if dataPeserta[i].nama == cari {

			fmt.Println("Data ditemukan")
			fmt.Println("ID    :", dataPeserta[i].id)
			fmt.Println("Nama  :", dataPeserta[i].nama)
			fmt.Println("Minat :", dataPeserta[i].bidangMinat)

			ketemu = true
		}
	}

	if ketemu == false {

		fmt.Println("Data tidak ditemukan")
	}
}

func selectionSortID() {
	var i int
	var j int
	var min int
	var temp Peserta

	for i = 0; i < jumlah-1; i++ {

		min = i

		for j = i + 1; j < jumlah; j++ {

			if dataPeserta[j].id < dataPeserta[min].id {

				min = j
			}
		}

		temp = dataPeserta[i]
		dataPeserta[i] = dataPeserta[min]
		dataPeserta[min] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan ID")
}

func insertionSortNama() {
	var i int
	var j int
	var temp Peserta

	for i = 1; i < jumlah; i++ {

		temp = dataPeserta[i]
		j = i - 1

		for j >= 0 && dataPeserta[j].nama > temp.nama {

			dataPeserta[j+1] = dataPeserta[j]
			j = j - 1
		}

		dataPeserta[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan Nama")
}

func binarySearchNama() {
	var kiri int
	var kanan int
	var tengah int
	var cari string
	var ketemu bool

	insertionSortNama()

	fmt.Println("=== Binary Search ===")

	fmt.Print("Cari Nama : ")
	fmt.Scan(&cari)

	kiri = 0
	kanan = jumlah - 1
	ketemu = false

	for kiri <= kanan {

		tengah = (kiri + kanan) / 2

		if dataPeserta[tengah].nama == cari {

			fmt.Println("Data ditemukan")
			fmt.Println("ID    :", dataPeserta[tengah].id)
			fmt.Println("Nama  :", dataPeserta[tengah].nama)
			fmt.Println("Minat :", dataPeserta[tengah].bidangMinat)

			ketemu = true
			kiri = kanan + 1

		} else if dataPeserta[tengah].nama < cari {

			kiri = tengah + 1

		} else {

			kanan = tengah - 1
		}
	}

	if ketemu == false {

		fmt.Println("Data tidak ditemukan")
	}
}

func statistik() {
	var i int
	var desain int
	var coding int
	var bisnis int
	var totalAktif int

	for i = 0; i < jumlah; i++ {

		if dataPeserta[i].bidangMinat == "Desain" {

			desain = desain + 1
		}

		if dataPeserta[i].bidangMinat == "Coding" {

			coding = coding + 1
		}

		if dataPeserta[i].bidangMinat == "Bisnis" {

			bisnis = bisnis + 1
		}

		if dataPeserta[i].aktif == true {

			totalAktif = totalAktif + 1
		}
	}

	fmt.Println("=== Statistik ===")
	fmt.Println("Desain      :", desain)
	fmt.Println("Coding      :", coding)
	fmt.Println("Bisnis      :", bisnis)
	fmt.Println("Total Aktif :", totalAktif)
}

func menu() {

	fmt.Println()
	fmt.Println("===== KURSUSIN =====")
	fmt.Println("1. Tambah Peserta")
	fmt.Println("2. Tampil Peserta")
	fmt.Println("3. Ubah Peserta")
	fmt.Println("4. Hapus Peserta")
	fmt.Println("5. Sequential Search")
	fmt.Println("6. Binary Search")
	fmt.Println("7. Selection Sort ID")
	fmt.Println("8. Insertion Sort Nama")
	fmt.Println("9. Statistik")
	fmt.Println("0. Keluar")
}

func main() {
	var pilih int

	pilih = -1

	for pilih != 0 {

		menu()

		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			tambahPeserta()

		} else if pilih == 2 {

			tampilPeserta()

		} else if pilih == 3 {

			ubahPeserta()

		} else if pilih == 4 {

			hapusPeserta()

		} else if pilih == 5 {

			sequentialSearchNama()

		} else if pilih == 6 {

			binarySearchNama()

		} else if pilih == 7 {

			selectionSortID()

		} else if pilih == 8 {

			insertionSortNama()

		} else if pilih == 9 {

			statistik()

		} else if pilih == 0 {

			fmt.Println("Program selesai")

		} else {

			fmt.Println("Menu tidak tersedia")
		}
	}
}
