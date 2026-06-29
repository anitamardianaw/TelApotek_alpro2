package main

import "fmt"

const NMAX int = 99

type Obat struct {
	ID             int
	Nama           string
	Gejala         string
	Kategori       string
	Stok           int
	TanggalExpired string
}

type Apotek [NMAX]Obat

func main() {
	var data Apotek
	var total, pilihan, awal int

	welcome()
	fmt.Println("1. Gunakan data dummy")
	fmt.Println("2. Input data sendiri")
	fmt.Print("Pilih: ")
	fmt.Scan(&awal)

	for awal != 1 && awal != 2 { // validate pilihan awal
		fmt.Println("Pilihan tidak valid.")
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&awal)
	}

	if awal == 1 {
		loadDummyData(&data, &total)
		fmt.Println("Data dummy berhasil dimuat.")
	} else {
		tambahObat(&data, &total)
	}

	pilihan = -1
	for pilihan != 0 {
		menu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahObat(&data, &total)
		case 2:
			tampilkanObat(data, total)
		case 3:
			ubahObat(&data, total)
		case 4:
			hapusObat(&data, &total)
		case 5:
			menuCari(data, total)
		case 6:
			menuSort(&data, total)
		case 7:
			menuLaporan(data, total)
		case 0:
			fmt.Println("Terima kasih sudah menggunakan TelApotek!")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func welcome() {
	var nama string

	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scan(&nama)
	fmt.Println("\nHalo", nama, "selamat datang di TelApotek!")
	fmt.Println("Silakan masukkan data obat terlebih dahulu.")
	fmt.Println()
}

func menu() {
	fmt.Println("╔═════════════════════════════════════════╗")
	fmt.Println("║                TelApotek                ║")
	fmt.Println("║       Aplikasi Manajemen Stok Obat      ║")
	fmt.Println("╠═════════════════════════════════════════╣")
	fmt.Println("║ 1. Tambah Obat                          ║")
	fmt.Println("║ 2. Tampilkan Obat                       ║")
	fmt.Println("║ 3. Ubah Data Obat                       ║")
	fmt.Println("║ 4. Hapus Data Obat                      ║")
	fmt.Println("║ 5. Cari Obat                            ║")
	fmt.Println("║ 6. Urutkan Obat                         ║")
	fmt.Println("║ 7. Laporan Stok                         ║")
	fmt.Println("║ 0. Exit                                 ║")
	fmt.Println("╚═════════════════════════════════════════╝")
	fmt.Print("Pilih menu: ")
}

func tambahObat(A *Apotek, total *int) {
	var jumlah, i int

	fmt.Print("Masukkan jumlah obat yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)

	for jumlah < 1 || *total+jumlah > NMAX { // validate jumlah obat
		fmt.Println("Jumlah tidak valid.")
		fmt.Print("Masukkan jumlah obat: ")
		fmt.Scan(&jumlah)
	}

	i = 0
	for i < jumlah {
		(*A)[*total].ID = *total + 1

		fmt.Println("\nData obat ke", *total+1)
		fmt.Print("Nama obat: ")
		fmt.Scan(&(*A)[*total].Nama)

		fmt.Print("Gejala (gunakan _ jika lebih dari satu kata): ")
		fmt.Scan(&(*A)[*total].Gejala)

		pilihKategori(A, *total)

		fmt.Print("Stok: ")
		fmt.Scan(&(*A)[*total].Stok)
		for (*A)[*total].Stok < 0 { // validate stok tidak negatif
			fmt.Print("Stok: ")
			fmt.Scan(&(*A)[*total].Stok)
		}

		fmt.Print("Tanggal expired (YYYY-MM-DD): ")
		fmt.Scan(&(*A)[*total].TanggalExpired)
		for len((*A)[*total].TanggalExpired) != 10 { // validate format tanggal
			fmt.Println("Format tanggal harus YYYY-MM-DD.")
			fmt.Print("Tanggal expired (YYYY-MM-DD): ")
			fmt.Scan(&(*A)[*total].TanggalExpired)
		}

		*total = *total + 1
		i++
	}

	fmt.Println("Data obat berhasil ditambahkan.")
}

func pilihKategori(A *Apotek, idx int) {
	var pilih int

	fmt.Println("Pilih bentuk/kategori obat:")
	fmt.Println("1. Tablet")
	fmt.Println("2. Kapsul")
	fmt.Println("3. Sirup")
	fmt.Println("4. Liquid")
	fmt.Println("5. Salep")
	fmt.Println("6. Tetes")
	fmt.Print("Pilih kategori: ")
	fmt.Scan(&pilih)

	for pilih < 1 || pilih > 6 { // validate pilihan kategori
		fmt.Println("Pilihan tidak valid.")
		fmt.Print("Pilih kategori: ")
		fmt.Scan(&pilih)
	}

	switch pilih {
	case 1:
		(*A)[idx].Kategori = "Tablet"
	case 2:
		(*A)[idx].Kategori = "Kapsul"
	case 3:
		(*A)[idx].Kategori = "Sirup"
	case 4:
		(*A)[idx].Kategori = "Liquid"
	case 5:
		(*A)[idx].Kategori = "Salep"
	case 6:
		(*A)[idx].Kategori = "Tetes"
	}
}

func tampilkanObat(A Apotek, total int) {
	var i int

	if total == 0 {
		fmt.Println("Belum ada data obat.")
	} else { // display obat in table format
		fmt.Println()
		fmt.Println("===================================================================================")
		fmt.Printf("%-4s %-20s %-20s %-15s %-8s %-15s\n", "ID", "Nama", "Gejala", "Kategori", "Stok", "Expired")
		fmt.Println("-----------------------------------------------------------------------------------")

		i = 0
		for i < total {
			fmt.Printf("%-4d %-20s %-20s %-15s %-8d %-15s\n", A[i].ID, A[i].Nama, A[i].Gejala, A[i].Kategori, A[i].Stok, A[i].TanggalExpired)
			i++
		}
		fmt.Println("===================================================================================")
	}
}

func ubahObat(A *Apotek, total int) { // modify obat based on ID
	var id, idx, pilih int

	tampilkanObat(*A, total)
	fmt.Print("Masukkan ID obat yang ingin diubah: ")
	fmt.Scan(&id)

	selectionSortID(A, total)
	idx = binarySearchID(*A, total, id)

	if idx != -1 {
		fmt.Println()
		fmt.Println("Ubah Data Obat")
		fmt.Println("1. Nama")
		fmt.Println("2. Gejala")
		fmt.Println("3. Kategori")
		fmt.Println("4. Stok")
		fmt.Println("5. Tanggal Expired")
		fmt.Print("Pilih data yang ingin diubah: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Nama baru: ")
			fmt.Scan(&(*A)[idx].Nama)
		case 2:
			fmt.Print("Gejala baru (gunakan _ jika lebih dari satu kata): ")
			fmt.Scan(&(*A)[idx].Gejala)
		case 3:
			pilihKategori(A, idx)
		case 4:
			fmt.Print("Stok baru: ")
			fmt.Scan(&(*A)[idx].Stok)
			for (*A)[idx].Stok < 0 {
				fmt.Println("Stok tidak boleh negatif.")
				fmt.Print("Stok baru: ")
				fmt.Scan(&(*A)[idx].Stok)
			}
		case 5:
			fmt.Print("Tanggal expired baru (YYYY-MM-DD): ")
			fmt.Scan(&(*A)[idx].TanggalExpired)
			for len((*A)[idx].TanggalExpired) != 10 {
				fmt.Println("Format tanggal harus YYYY-MM-DD.")
				fmt.Print("Tanggal expired baru (YYYY-MM-DD): ")
				fmt.Scan(&(*A)[idx].TanggalExpired)
			}
		default:
			fmt.Println("Pilihan tidak valid.")
		}

		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Obat tidak ditemukan.")
	}
}

func hapusObat(A *Apotek, total *int) { // hapus obat based on ID
	var id, idx, i int

	tampilkanObat(*A, *total)
	fmt.Print("Masukkan ID obat yang ingin dihapus: ")
	fmt.Scan(&id)

	selectionSortID(A, *total)
	idx = binarySearchID(*A, *total, id)

	if idx != -1 {
		i = idx
		for i < *total-1 {
			(*A)[i] = (*A)[i+1]
			(*A)[i].ID = i + 1
			i++
		}

		*total = *total - 1
		fmt.Println("Data obat berhasil dihapus.")
	} else {
		fmt.Println("Obat tidak ditemukan.")
	}
}

func menuCari(A Apotek, total int) {
	var pilih int

	fmt.Println("\nCari Obat")
	fmt.Println("1. Cari berdasarkan Gejala")
	fmt.Println("2. Cari berdasarkan ID")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih menu cari: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		cariGejala(A, total)
	case 2:
		cariID(A, total)
	case 0:
		fmt.Println("Kembali ke menu utama.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func cariGejala(A Apotek, total int) { // sequential search
	var gejala string
	var i int
	var found bool

	fmt.Print("Masukkan gejala yang dicari: ")
	fmt.Scan(&gejala)

	found = false
	i = 0

	fmt.Println("\nHasil pencarian:")
	for i < total {
		if A[i].Gejala == gejala {
			fmt.Println(A[i].ID, A[i].Nama, A[i].Kategori, A[i].Stok, A[i].TanggalExpired)
			found = true
		}
		i++
	}
	if !found {
		fmt.Println("Obat dengan gejala tersebut tidak ditemukan.")
	}
}

func binarySearchID(A Apotek, total int, id int) int { // get the idx obat from ID
	var left, right, mid int

	left = 0
	right = total - 1

	for left <= right {
		mid = (left + right) / 2

		if A[mid].ID == id {
			return mid
		} else if A[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func cariID(A Apotek, total int) { // binary search
	var id, idx int

	selectionSortID(&A, total)
	fmt.Print("Masukkan ID obat: ")
	fmt.Scan(&id)

	idx = binarySearchID(A, total, id)

	if idx != -1 {
		fmt.Println("Obat ditemukan:")
		fmt.Println("ID              :", A[idx].ID)
		fmt.Println("Nama            :", A[idx].Nama)
		fmt.Println("Gejala          :", A[idx].Gejala)
		fmt.Println("Kategori        :", A[idx].Kategori)
		fmt.Println("Stok            :", A[idx].Stok)
		fmt.Println("Tanggal Expired :", A[idx].TanggalExpired)
	} else {
		fmt.Println("Obat tidak ditemukan.")
	}
}

func menuSort(A *Apotek, total int) {
	var pilih int

	fmt.Println("\nUrutkan Obat")
	fmt.Println("1. Urutkan Expired Terdekat")
	fmt.Println("2. Urutkan Stok Terkecil")
	fmt.Println("3. Urutkan Berdasarkan ID")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih menu sort: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		selectionSortExpired(A, total)
		fmt.Println("Data berhasil diurutkan berdasarkan expired.")
		tampilkanObat(*A, total)
	case 2:
		insertionSortStok(A, total)
		fmt.Println("Data berhasil diurutkan berdasarkan stok.")
		tampilkanObat(*A, total)
	case 3:
		selectionSortID(A, total)
		fmt.Println("Data berhasil diurutkan berdasarkan ID.")
		tampilkanObat(*A, total)
	case 0:
		fmt.Println("Kembali ke menu utama.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func selectionSortExpired(A *Apotek, total int) { // ascending sort by expired date
	var pass, i, idx int
	var temp Obat

	pass = 0
	for pass < total-1 {
		idx = pass
		i = pass + 1

		for i < total {
			if (*A)[i].TanggalExpired < (*A)[idx].TanggalExpired {
				idx = i
			}
			i++
		}

		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp

		pass++
	}
}

func insertionSortStok(A *Apotek, total int) { // ascending sort by stock
	var pass, i int
	var temp Obat

	pass = 1
	for pass < total {
		temp = (*A)[pass]
		i = pass

		for i > 0 && temp.Stok < (*A)[i-1].Stok {
			(*A)[i] = (*A)[i-1]
			i--
		}

		(*A)[i] = temp
		pass++
	}
}

func selectionSortID(A *Apotek, total int) { // ascending sort by ID
	var pass, i, idx int
	var temp Obat

	pass = 0
	for pass < total-1 {
		idx = pass
		i = pass + 1

		for i < total {
			if (*A)[i].ID < (*A)[idx].ID {
				idx = i
			}
			i++
		}

		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp

		pass++
	}
}

func menuLaporan(A Apotek, total int) {
	var pilih int

	fmt.Println("\nLaporan Stok")
	fmt.Println("1. Obat Stok Terbanyak")
	fmt.Println("2. Obat Stok Terkecil")
	fmt.Println("3. Total Seluruh Stok")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih menu laporan: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		tampilMaxStok(A, total)
	case 2:
		tampilMinStok(A, total)
	case 3:
		fmt.Println("Total seluruh stok obat:", totalStokRekursif(A, total))
	case 0:
		fmt.Println("Kembali ke menu utama.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func findMaxStok(A Apotek, total int) int {
	var i, idx int

	idx = 0
	i = 1

	for i < total {
		if A[i].Stok > A[idx].Stok {
			idx = i
		}
		i++
	}

	return idx
}

func tampilMaxStok(A Apotek, total int) {
	var idx int

	if total == 0 {
		fmt.Println("Belum ada data obat.")
	} else {
		idx = findMaxStok(A, total)
		fmt.Println("Obat dengan stok terbanyak:")
		fmt.Println(A[idx].Nama, "Stok:", A[idx].Stok)
	}
}

func findMinStok(A Apotek, total int) int {
	var i, idx int

	idx = 0
	i = 1

	for i < total {
		if A[i].Stok < A[idx].Stok {
			idx = i
		}
		i++
	}

	return idx
}

func tampilMinStok(A Apotek, total int) {
	var idx int

	if total == 0 {
		fmt.Println("Belum ada data obat.")
	} else {
		idx = findMinStok(A, total)
		fmt.Println("Obat dengan stok terkecil:")
		fmt.Println(A[idx].Nama, "Stok:", A[idx].Stok)
	}
}

func totalStokRekursif(A Apotek, total int) int {
	if total == 0 {
		return 0
	}

	return A[total-1].Stok + totalStokRekursif(A, total-1)
}

func loadDummyData(A *Apotek, total *int) {
	(*A)[0] = Obat{1, "Paracetamol", "Demam", "Tablet", 20, "2027-01-10"}
	(*A)[1] = Obat{2, "Bodrex", "Pusing", "Tablet", 15, "2027-05-12"}
	(*A)[2] = Obat{3, "OBH", "Batuk", "Sirup", 8, "2027-11-20"}
	(*A)[3] = Obat{4, "Promag", "Maag", "Tablet", 12, "2026-08-02"}
	(*A)[4] = Obat{5, "Insto", "Mata_merah", "Tetes", 5, "2028-12-30"}

	(*A)[5] = Obat{6, "Betadine", "Luka", "Liquid", 10, "2028-03-15"}
	(*A)[6] = Obat{7, "Sanmol", "Demam", "Sirup", 25, "2026-09-22"}
	(*A)[7] = Obat{8, "Counterpain", "Nyeri_otot", "Salep", 7, "2027-06-10"}
	(*A)[8] = Obat{9, "Diapet", "Diare", "Kapsul", 9, "2027-08-19"}
	(*A)[9] = Obat{10, "Panadol", "Demam", "Tablet", 30, "2026-12-31"}

	*total = 10
}
