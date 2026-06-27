package main

import "fmt"

const MAKS_MEJA = 50
const MAKS_PELANGGAN = 200
const MAKS_RESERVASI = 100

type Meja struct {
	nomorMeja int
	kapasitas int
	tersedia  bool
}

type Pelanggan struct {
	idPelanggan int
	nama        string
	noTelepon   string
}

type Reservasi struct {
	idReservasi int
	idPelanggan int
	nomorMeja   int
	tanggal     string
	jam         string
	aktif       bool
}

var daftarMeja [MAKS_MEJA]Meja
var jumlahMeja int

var daftarPelanggan [MAKS_PELANGGAN]Pelanggan
var jumlahPelanggan int

var daftarReservasi [MAKS_RESERVASI]Reservasi
var jumlahReservasi int

var counterIdPelanggan int = 1
var counterIdReservasi int = 1

func inputInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func inputString() string {
	var s string
	fmt.Scan(&s)
	return s
}

func tampilHeader() {
	fmt.Println("===================================")
	fmt.Println("+++       ReservaResto          +++")
	fmt.Println("+++ Aplikasi Reservasi Restoran +++")
	fmt.Println("===================================")
}

func tampilSemuaMeja() {
	fmt.Println("\n--- Daftar Meja ---")
	if jumlahMeja == 0 {
		fmt.Println("Belum ada data meja.")
		return
	}
	fmt.Println("No  | Nomor Meja | Kapasitas | Status")
	fmt.Println("----|------------|-----------|--------")
	for i := 0; i < jumlahMeja; i++ {
		status := "Tersedia"
		if !daftarMeja[i].tersedia {
			status = "Terisi"
		}
		fmt.Printf("%-3d | %-10d | %-9d | %s\n", i+1, daftarMeja[i].nomorMeja, daftarMeja[i].kapasitas, status)
	}
}

func tambahMeja() {
	if jumlahMeja >= MAKS_MEJA {
		fmt.Println("Data meja sudah penuh!")
		return
	}
	fmt.Print("Nomor meja: ")
	nomor := inputInt()
	for i := 0; i < jumlahMeja; i++ {
		if daftarMeja[i].nomorMeja == nomor {
			fmt.Println("Nomor meja sudah ada!")
			return
		}
	}
	fmt.Print("Kapasitas kursi: ")
	kapasitas := inputInt()
	daftarMeja[jumlahMeja].nomorMeja = nomor
	daftarMeja[jumlahMeja].kapasitas = kapasitas
	daftarMeja[jumlahMeja].tersedia = true
	jumlahMeja++
	fmt.Println("Meja berhasil ditambahkan!")
}

func ubahMeja() {
	fmt.Print("Nomor meja yang diubah: ")
	nomor := inputInt()
	idx := -1
	for i := 0; i < jumlahMeja; i++ {
		if daftarMeja[i].nomorMeja == nomor {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("Meja tidak ditemukan!")
		return
	}
	fmt.Printf("Kapasitas saat ini: %d\n", daftarMeja[idx].kapasitas)
	fmt.Print("Kapasitas baru: ")
	daftarMeja[idx].kapasitas = inputInt()
	fmt.Println("Data meja berhasil diubah!")
}

func hapusMeja() {
	fmt.Print("Nomor meja yang dihapus: ")
	nomor := inputInt()
	idx := -1
	for i := 0; i < jumlahMeja; i++ {
		if daftarMeja[i].nomorMeja == nomor {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("Meja tidak ditemukan!")
		return
	}
	// geser elemen ke kiri
	for i := idx; i < jumlahMeja-1; i++ {
		daftarMeja[i] = daftarMeja[i+1]
	}
	jumlahMeja--
	fmt.Println("Meja berhasil dihapus!")
}

func sequentialSearch(cariNomor bool, nilai int) int {
	// cek satu per satu dari depan 
	for i := 0; i < jumlahMeja; i++ {
		if cariNomor && daftarMeja[i].nomorMeja == nilai {
			return i
		} else if !cariNomor && daftarMeja[i].kapasitas == nilai {
			return i
		}
	}
	return -1
}

func binarySearch(nilai int, cariNomor bool) int {
	kiri := 0
	kanan := jumlahMeja - 1
	// bagi dua ruang pencarian tiap iterasi
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		var p int
		if cariNomor {
			p = daftarMeja[tengah].nomorMeja
		} else {
			p = daftarMeja[tengah].kapasitas
		}
		if p == nilai {
			return tengah
		} else if p < nilai {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func menuCariMeja() {
	fmt.Println("\n--- Cari Meja ---")
	fmt.Println("1. Berdasarkan Nomor Meja")
	fmt.Println("2. Berdasarkan Kapasitas")
	fmt.Print("Pilihan: ")
	pilihanCari := inputInt()
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (data harus terurut)")
	fmt.Print("Metode: ")
	metode := inputInt()
	fmt.Print("Nilai yang dicari: ")
	nilai := inputInt()

	cariNomor := pilihanCari == 1
	idx := -1
	namaMetode := ""

	if metode == 1 {
		namaMetode = "Sequential Search"
		idx = sequentialSearch(cariNomor, nilai)
	} else if metode == 2 {
		namaMetode = "Binary Search"
		idx = binarySearch(nilai, cariNomor)
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Printf("Metode: %s\n", namaMetode)
	if idx == -1 {
		fmt.Println("Meja tidak ditemukan!")
	} else {
		status := "Tersedia"
		if !daftarMeja[idx].tersedia {
			status = "Terisi"
		}
		fmt.Printf("Meja ditemukan! Nomor: %d | Kapasitas: %d | Status: %s\n",
			daftarMeja[idx].nomorMeja, daftarMeja[idx].kapasitas, status)
	}
}

func selectionSort() {
	// cari minimum, tukar ke posisi i
	for i := 0; i < jumlahMeja-1; i++ {
		idxMin := i
		for j := i + 1; j < jumlahMeja; j++ {
			if daftarMeja[j].kapasitas < daftarMeja[idxMin].kapasitas {
				idxMin = j
			}
		}
		if idxMin != i {
			daftarMeja[i], daftarMeja[idxMin] = daftarMeja[idxMin], daftarMeja[i]
		}
	}
	fmt.Println("Diurutkan dengan Selection Sort!")
}

func insertionSort() {
	// sisipkan elemen ke posisi yang tepat
	for i := 1; i < jumlahMeja; i++ {
		kunci := daftarMeja[i]
		j := i - 1
		for j >= 0 && daftarMeja[j].kapasitas > kunci.kapasitas {
			daftarMeja[j+1] = daftarMeja[j]
			j--
		}
		daftarMeja[j+1] = kunci
	}
	fmt.Println("Diurutkan dengan Insertion Sort!")
}

func menuUrutMeja() {
	fmt.Println("\n--- Urutkan Meja (Kapasitas Terkecil ke Terbesar) ---")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilihan: ")
	pilihan := inputInt()
	if pilihan == 1 {
		selectionSort()
	} else if pilihan == 2 {
		insertionSort()
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}
	tampilSemuaMeja()
}

func tampilSemuaPelanggan() {
	fmt.Println("\n--- Daftar Pelanggan ---")
	if jumlahPelanggan == 0 {
		fmt.Println("Belum ada data pelanggan.")
		return
	}
	fmt.Println("ID  | Nama                 | No. Telepon")
	fmt.Println("----|----------------------|------------")
	for i := 0; i < jumlahPelanggan; i++ {
		fmt.Printf("%-3d | %-20s | %s\n",
			daftarPelanggan[i].idPelanggan, daftarPelanggan[i].nama, daftarPelanggan[i].noTelepon)
	}
}

func tambahPelanggan() {
	if jumlahPelanggan >= MAKS_PELANGGAN {
		fmt.Println("Data pelanggan sudah penuh!")
		return
	}
	fmt.Print("Nama: ")
	nama := inputString()
	fmt.Print("No. Telepon: ")
	noTelp := inputString()
	daftarPelanggan[jumlahPelanggan].idPelanggan = counterIdPelanggan
	daftarPelanggan[jumlahPelanggan].nama = nama
	daftarPelanggan[jumlahPelanggan].noTelepon = noTelp
	jumlahPelanggan++
	counterIdPelanggan++
	fmt.Println("Pelanggan berhasil ditambahkan!")
}

func ubahPelanggan() {
	fmt.Print("ID pelanggan yang diubah: ")
	id := inputInt()
	idx := -1
	for i := 0; i < jumlahPelanggan; i++ {
		if daftarPelanggan[i].idPelanggan == id {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("Pelanggan tidak ditemukan!")
		return
	}
	fmt.Printf("Nama saat ini: %s\n", daftarPelanggan[idx].nama)
	fmt.Print("Nama baru: ")
	daftarPelanggan[idx].nama = inputString()
	fmt.Print("No. Telepon baru: ")
	daftarPelanggan[idx].noTelepon = inputString()
	fmt.Println("Data pelanggan berhasil diubah!")
}

func hapusPelanggan() {
	fmt.Print("ID pelanggan yang dihapus: ")
	id := inputInt()
	idx := -1
	for i := 0; i < jumlahPelanggan; i++ {
		if daftarPelanggan[i].idPelanggan == id {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("Pelanggan tidak ditemukan!")
		return
	}
	// geser elemen ke kiri (alan)
	for i := idx; i < jumlahPelanggan-1; i++ {
		daftarPelanggan[i] = daftarPelanggan[i+1]
	}
	jumlahPelanggan--
	fmt.Println("Pelanggan berhasil dihapus!")
}

func tampilSemuaReservasi() {
	fmt.Println("\n--- Daftar Reservasi ---")
	fmt.Println("ID  | ID Pelanggan | Meja | Tanggal    | Jam  ")
	fmt.Println("----|--------------|------|------------|------")
	ada := false
	for i := 0; i < jumlahReservasi; i++ {
		if daftarReservasi[i].aktif {
			fmt.Printf("%-3d | %-12d | %-4d | %-10s | %s\n",
				daftarReservasi[i].idReservasi, daftarReservasi[i].idPelanggan,
				daftarReservasi[i].nomorMeja, daftarReservasi[i].tanggal, daftarReservasi[i].jam)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Belum ada reservasi aktif.")
	}
}

func cariIndexPelanggan(id int) int {
	for i := 0; i < jumlahPelanggan; i++ {
		if daftarPelanggan[i].idPelanggan == id {
			return i
		}
	}
	return -1
}

func cariIndexMeja(nomor int) int {
	for i := 0; i < jumlahMeja; i++ {
		if daftarMeja[i].nomorMeja == nomor {
			return i
		}
	}
	return -1
}

func tambahReservasi() {
	if jumlahReservasi >= MAKS_RESERVASI {
		fmt.Println("Data reservasi sudah penuh!")
		return
	}
	fmt.Print("ID Pelanggan: ")
	idPelanggan := inputInt()
	if cariIndexPelanggan(idPelanggan) == -1 {
		fmt.Println("Pelanggan tidak ditemukan!")
		return
	}
	fmt.Print("Nomor Meja: ")
	nomorMeja := inputInt()
	idxMeja := cariIndexMeja(nomorMeja)
	if idxMeja == -1 {
		fmt.Println("Meja tidak ditemukan!")
		return
	}
	if !daftarMeja[idxMeja].tersedia {
		fmt.Println("Meja sedang terisi!")
		return
	}
	fmt.Print("Tanggal (dd-mm-yyyy): ")
	tanggal := inputString()
	fmt.Print("Jam (hh:mm): ")
	jam := inputString()
	daftarReservasi[jumlahReservasi].idReservasi = counterIdReservasi
	daftarReservasi[jumlahReservasi].idPelanggan = idPelanggan
	daftarReservasi[jumlahReservasi].nomorMeja = nomorMeja
	daftarReservasi[jumlahReservasi].tanggal = tanggal
	daftarReservasi[jumlahReservasi].jam = jam
	daftarReservasi[jumlahReservasi].aktif = true
	jumlahReservasi++
	counterIdReservasi++
	daftarMeja[idxMeja].tersedia = false
	fmt.Println("Reservasi berhasil dibuat!")
}

func batalReservasi() {
	fmt.Print("ID reservasi yang dibatalkan: ")
	id := inputInt()
	idx := -1
	for i := 0; i < jumlahReservasi; i++ {
		if daftarReservasi[i].idReservasi == id && daftarReservasi[i].aktif {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("Reservasi tidak ditemukan!")
		return
	}
	idxMeja := cariIndexMeja(daftarReservasi[idx].nomorMeja)
	if idxMeja != -1 {
		daftarMeja[idxMeja].tersedia = true
	}
	daftarReservasi[idx].aktif = false
	fmt.Println("Reservasi berhasil dibatalkan!")
}

func tampilStatistik() {
	fmt.Println("\n=================================")
	fmt.Println("+++       ReservaResto          +++")
	fmt.Println("+++         Statistik           +++")
	fmt.Println("=================================")
	if jumlahReservasi == 0 {
		fmt.Println("Belum ada data reservasi.")
		return
	}

	// hitung reservasi per tanggal pakai array paralel (alan)
	var tanggalList [MAKS_RESERVASI]string
	var countHari [MAKS_RESERVASI]int
	jmlTgl := 0
	for i := 0; i < jumlahReservasi; i++ {
		if !daftarReservasi[i].aktif {
			continue
		}
		tgl := daftarReservasi[i].tanggal
		ketemu := false
		for j := 0; j < jmlTgl; j++ {
			if tanggalList[j] == tgl {
				countHari[j]++
				ketemu = true
			}
		}
		if !ketemu {
			tanggalList[jmlTgl] = tgl
			countHari[jmlTgl] = 1
			jmlTgl++
		}
	}
	fmt.Println("\n[Reservasi Per Hari]")
	fmt.Println("Tanggal     | Jumlah")
	fmt.Println("------------|-------")
	for i := 0; i < jmlTgl; i++ {
		fmt.Printf("%-11s | %d\n", tanggalList[i], countHari[i])
	}

	// hitung meja paling sering dipesan (alan)
	var mejaList [MAKS_MEJA]int
	var countMeja [MAKS_MEJA]int
	jmlMejaStats := 0
	for i := 0; i < jumlahReservasi; i++ {
		if !daftarReservasi[i].aktif {
			continue
		}
		nm := daftarReservasi[i].nomorMeja
		ketemu := false
		for j := 0; j < jmlMejaStats; j++ {
			if mejaList[j] == nm {
				countMeja[j]++
				ketemu = true
			}
		}
		if !ketemu {
			mejaList[jmlMejaStats] = nm
			countMeja[jmlMejaStats] = 1
			jmlMejaStats++
		}
	}
	idxTop := 0
	for i := 1; i < jmlMejaStats; i++ {
		if countMeja[i] > countMeja[idxTop] {
			idxTop = i
		}
	}
	fmt.Println("\n[Meja Terpopuler]")
	if jmlMejaStats > 0 {
		fmt.Printf("Nomor Meja : %d\n", mejaList[idxTop])
		fmt.Printf("Dipesan    : %d kali\n", countMeja[idxTop])
	} else {
		fmt.Println("Belum ada data.")
	}
	fmt.Println("=================================")
}

func menuMeja() {
	for {
		fmt.Println("\n=== Menu Meja ===")
		fmt.Println("1. Lihat Semua Meja")
		fmt.Println("2. Tambah Meja")
		fmt.Println("3. Ubah Meja")
		fmt.Println("4. Hapus Meja")
		fmt.Println("5. Cari Meja")
		fmt.Println("6. Urutkan Meja")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		pilihan := inputInt()
		if pilihan == 1 {
			tampilSemuaMeja()
		} else if pilihan == 2 {
			tambahMeja()
		} else if pilihan == 3 {
			ubahMeja()
		} else if pilihan == 4 {
			hapusMeja()
		} else if pilihan == 5 {
			menuCariMeja()
		} else if pilihan == 6 {
			menuUrutMeja()
		} else if pilihan == 0 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuPelanggan() {
	for {
		fmt.Println("\n=== Menu Pelanggan ===")
		fmt.Println("1. Lihat Semua Pelanggan")
		fmt.Println("2. Tambah Pelanggan")
		fmt.Println("3. Ubah Pelanggan")
		fmt.Println("4. Hapus Pelanggan")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		pilihan := inputInt()
		if pilihan == 1 {
			tampilSemuaPelanggan()
		} else if pilihan == 2 {
			tambahPelanggan()
		} else if pilihan == 3 {
			ubahPelanggan()
		} else if pilihan == 4 {
			hapusPelanggan()
		} else if pilihan == 0 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuReservasi() {
	for {
		fmt.Println("\n=== Menu Reservasi ===")
		fmt.Println("1. Lihat Semua Reservasi")
		fmt.Println("2. Buat Reservasi")
		fmt.Println("3. Batalkan Reservasi")
		fmt.Println("0. Kembali")
		fmt.Print("Pilihan: ")
		pilihan := inputInt()
		if pilihan == 1 {
			tampilSemuaReservasi()
		} else if pilihan == 2 {
			tambahReservasi()
		} else if pilihan == 3 {
			batalReservasi()
		} else if pilihan == 0 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func main() {
	tampilHeader()
	for {
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("1. Manajemen Meja")
		fmt.Println("2. Manajemen Pelanggan")
		fmt.Println("3. Manajemen Reservasi")
		fmt.Println("4. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")
		pilihan := inputInt()
		if pilihan == 1 {
			menuMeja()
		} else if pilihan == 2 {
			menuPelanggan()
		} else if pilihan == 3 {
			menuReservasi()
		} else if pilihan == 4 {
			tampilStatistik()
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan ReservaResto!")
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

