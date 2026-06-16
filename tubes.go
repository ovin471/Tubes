package main

import "fmt"

const NMAX = 100

type Data struct {
	Nama, status       string
	riwayat_trasnsaksi int
}

type arrData [NMAX]Data

func main() {
	var data arrData
	var n, Pilihan, Pilihan2, Pilihan3, JumlahNominal int

	fmt.Print("Silahkan Masukkan Banyak Data : ")
	fmt.Scan(&n)

	fmt.Print("Berapa Total Nominal Untuk KAS (Per Bulan) : ")
	fmt.Scan(&JumlahNominal)

	fmt.Println("Silahkan Masukkan Data Awal")
	InputData(&data, n, JumlahNominal)

	fmt.Println("")
	fmt.Println("================================================")
	fmt.Printf("| %-44s |\n", "Opsi Yang Bisa Dipilih")
	fmt.Println("|==============================================|")
	fmt.Printf("| %-2d | %-39s |\n", 1, "Tambahan Data")
	fmt.Printf("| %-2d | %-39s |\n", 2, "Ubah Data")
	fmt.Printf("| %-2d | %-39s |\n", 3, "MenghapusData")
	fmt.Printf("| %-2d | %-39s |\n", 4, "Menambah Nominal")
	fmt.Printf("| %-2d | %-39s |\n", 5, "Cetak Data")
	fmt.Println("================================================")
	fmt.Printf("Silahkan Masukkan Pilihan : ")
	fmt.Scan(&Pilihan)
	fmt.Println("================================================")
	fmt.Println("")

	for Pilihan != 0 {
		switch Pilihan {
		case 1:
			fmt.Println("Silahkan Masukkan Tambahan Data")
			tambahData(&data, &n, JumlahNominal)
		case 2:
			mengubahData(&data, n, JumlahNominal)
		case 3:
			MenghapusData(&data, &n, JumlahNominal)
		case 4:
			menambahnominal(&data, n, JumlahNominal)
		case 5:
			fmt.Println("")
			fmt.Println("====================================================")
			fmt.Printf("| %-48s |\n", "Opsi Yang Bisa Dipilih")
			fmt.Println("|==================================================|")
			fmt.Printf("| %-3d| %-43s |\n", 1, "Search Data")
			fmt.Printf("| %-3d| %-43s |\n", 2, "Sort Data")
			fmt.Printf("| %-3d| %-43s |\n", 0, "Kembali Ke Menu Utama")
			fmt.Println("====================================================")
			fmt.Printf("Silahkan Masukkan Pilihan : ")
			fmt.Scan(&Pilihan2)
			fmt.Println("====================================================")
			fmt.Println("")

			for Pilihan2 != 0 {
				switch Pilihan2 {
				case 1:
					fmt.Println("")
					fmt.Println("========================================================")
					fmt.Printf("| %-52s |\n", "Metode Apa Yang Anda Ingin Digunakan??")
					fmt.Println("|======================================================|")
					fmt.Printf("| %2d | %-47s |\n", 1, "Sequential Search")
					fmt.Printf("| %2d | %-47s |\n", 2, "Binary Search")
					fmt.Println("========================================================")
					fmt.Printf("Silahkan Masukkan Pilihan : ")
					fmt.Scan(&Pilihan3)
					fmt.Println("========================================================")
					fmt.Println("")

					switch Pilihan3 {
					case 1:
						SequentialSearch(&data, n, JumlahNominal)
					case 2:
						BinarySearch(&data, n, JumlahNominal)
					}
				case 2:
					fmt.Println("")
					fmt.Println("========================================================")
					fmt.Printf("| %-52s |\n", "Metode Apa Yang Anda Ingin Digunakan??")
					fmt.Println("|======================================================|")
					fmt.Printf("| %2d | %-47s |\n", 1, "Insertion Sort")
					fmt.Printf("| %2d | %-47s |\n", 2, "Selection Sort")
					fmt.Println("========================================================")
					fmt.Printf("Silahkan Masukkan Pilihan : ")
					fmt.Scan(&Pilihan3)
					fmt.Println("========================================================")
					fmt.Println("")

					switch Pilihan3 {
					case 1:
						InsertionSort(&data, n, JumlahNominal)
					case 2:
						Selection(&data, n, JumlahNominal)
					}
				}

				fmt.Println("")
				fmt.Println("====================================================")
				fmt.Printf("| %-48s |\n", "Opsi Yang Bisa Dipilih")
				fmt.Println("|==================================================|")
				fmt.Printf("| %-3d| %-43s |\n", 1, "Search Data")
				fmt.Printf("| %-3d| %-43s |\n", 2, "Sort Data")
				fmt.Printf("| %-3d| %-43s |\n", 0, "Kembali Ke Menu Utama")
				fmt.Println("====================================================")
				fmt.Printf("Silahkan Masukkan Pilihan : ")
				fmt.Scan(&Pilihan2)
				fmt.Println("====================================================")
				fmt.Println("")
			}
		}
		fmt.Println("")
		fmt.Println("================================================")
		fmt.Printf("| %-44s |\n", "Opsi Yang Bisa Dipilih")
		fmt.Println("|==============================================|")
		fmt.Printf("| %-2d | %-39s |\n", 1, "Tambahan Data")
		fmt.Printf("| %-2d | %-39s |\n", 2, "Ubah Data")
		fmt.Printf("| %-2d | %-39s |\n", 3, "MenghapusData")
		fmt.Printf("| %-2d | %-39s |\n", 4, "Menambah Nominal")
		fmt.Printf("| %-2d | %-39s |\n", 5, "Cetak Data")
		fmt.Printf("| %-2d | %-39s |\n", 0, "Mengakhiri")
		fmt.Println("================================================")
		fmt.Printf("Silahkan Masukkan Pilihan : ")
		fmt.Scan(&Pilihan)
		fmt.Println("================================================")
		fmt.Println("")
	}

}

func cetakData(A arrData, n, JumlahNominal int) {
	var TotalSaldo, JumlahBelumLunas int

	for i := 1; i <= n; i++ {
		if A[i].riwayat_trasnsaksi < (n * (JumlahNominal * 6)) {
			A[i].status = "Belum Lunas"
		} else {
			A[i].status = "Lunas"
		}
	}

	fmt.Printf("|=====|=========|=====================|=============|\n")
	fmt.Printf("| %-3s | %-7s | %-19s | %-12s|\n", "No", "Nama", "Riwayat Transaksi", "Status")
	fmt.Printf("|=====|=========|=====================|=============|\n")
	for i := 1; i <= n; i++ {
		fmt.Printf("| %-3d | %-7s | %-19d | %-12s|\n",
			i,
			A[i].Nama,
			A[i].riwayat_trasnsaksi,
			A[i].status)
	}
	fmt.Printf("|===================================================|\n")

	for i := 1; i <= n; i++ {
		TotalSaldo = TotalSaldo + A[i].riwayat_trasnsaksi
	}

	fmt.Printf("| Total Saldo KAS : %-31d |\n", TotalSaldo)

	for i := 1; i <= n; i++ {
		if A[i].status == "Belum Lunas" {
			JumlahBelumLunas = JumlahBelumLunas + 1
		}
	}
	fmt.Printf("|===================================================|\n")
	fmt.Printf("| Jumlah Mahasiswa Yang Belum Lunas : %-13d |\n", JumlahBelumLunas)
	fmt.Printf("|===================================================|\n")
	fmt.Printf("| Jumlah Kas Selama 6 Bulan : %-21d |\n", (JumlahNominal*6)-TotalSaldo)
	fmt.Printf("|===================================================|\n")
}

func InputData(A *arrData, n, JumlahNominal int) {
	for i := 1; i <= n; i++ {
		fmt.Scan(
			&A[i].Nama,
			&A[i].riwayat_trasnsaksi)
	}

	for i := 1; i <= n; i++ {
		if A[i].riwayat_trasnsaksi < JumlahNominal {
			A[i].status = "Belum Lunas"
		} else {
			A[i].status = "Lunas"
		}
	}
}

func tambahData(A *arrData, n *int, JumlahNominal int) {
	var jumlahPenambahanData int
	fmt.Scan(&jumlahPenambahanData)

	for i := *n + 1; i <= *n+jumlahPenambahanData; i++ {
		fmt.Scan(
			&A[i].Nama,
			&A[i].riwayat_trasnsaksi)
	}

	for i := *n + 1; i <= *n+jumlahPenambahanData; i++ {
		if (*A)[i].riwayat_trasnsaksi < JumlahNominal {
			(*A)[i].status = "Belum Lunas"
		} else {
			(*A)[i].status = "Lunas"
		}
	}

	*n = *n + jumlahPenambahanData
}

func mengubahData(A *arrData, n, JumlahNominal int) {
	var CariNama, Nama string
	var banyakdata, nominal, hasil int
	hasil = -1

	fmt.Print("Jumlah Data yang ingin di ubah : ")
	fmt.Scan(&banyakdata)

	for a := 1; a <= banyakdata; a++ {
		fmt.Println("Masukkan nama yang di cari :")
		fmt.Scan(&CariNama)
		for i := 1; i <= n; i++ {
			if A[i].Nama == CariNama {
				fmt.Println("Nama yang di cari ditemukan", A[i].Nama, A[i].riwayat_trasnsaksi, A[i].status)
				fmt.Println("Massukkan Data Baru :")
				fmt.Print("Nama : ")
				fmt.Scan(&Nama)
				fmt.Print("Massukkan Riwayat Trasnsaksi : ")
				fmt.Scan(&nominal)
				A[i].Nama = Nama
				A[i].riwayat_trasnsaksi = nominal
				hasil = 0
			} else {
				hasil = 1
			}
		}
	}

	if hasil == 1 {
		fmt.Println("Nama Yang Di Cari Tidak Ditemukan ")
	}

	for i := 1; i <= n; i++ {
		if A[i].riwayat_trasnsaksi < JumlahNominal {
			A[i].status = "Belum Lunas"
		} else {
			A[i].status = "Lunas"
		}
	}

}

//Function untuk mengurutan Data dengan metode Insertion Sort
func InsertionSort(A *arrData, n, JumlahNominal int) {
	var pass, i int
	var temp Data
	pass = 2

	for pass <= n {
		i = pass
		temp = (*A)[pass]
		for i > 1 && (temp.riwayat_trasnsaksi) < ((*A)[i-1].riwayat_trasnsaksi) {
			(*A)[i] = (*A)[i-1]
			i = i - 1
		}
		(*A)[i] = temp
		pass = pass + 1
	}

	fmt.Printf("|===================================================|\n")
	fmt.Printf("|%-51s|\n", "Data Setelah Diurutkan Dengan Insertion Sort")

	cetakData(*A, n, JumlahNominal)
}

//Function untuk mengurutan Data dengan metode Selection Sort
func Selection(A *arrData, n, JumlahNominal int) {
	var posisi int
	var temp Data

	for a := 1; a <= n; a++ {
		posisi = a

		for j := a + 1; j <= n; j++ {
			if (*A)[j].riwayat_trasnsaksi > (*A)[posisi].riwayat_trasnsaksi {
				posisi = j
			} else if (*A)[j].riwayat_trasnsaksi == (*A)[posisi].riwayat_trasnsaksi {
				if (*A)[j].Nama < (*A)[posisi].Nama {
					posisi = j
				}
			}
		}

		temp = (*A)[posisi]
		(*A)[posisi] = (*A)[a]
		(*A)[a] = temp
	}

	fmt.Printf("|===================================================|\n")
	fmt.Printf("|%-51s|\n", "Data Setelah Diurutkan Dengan Selection Sort")

	cetakData(*A, n, JumlahNominal)
}

//Function untuk mencari data dengan metode Sequential Search
func SequentialSearch(A *arrData, n, JumlahNominal int) {
	var ketemu, k int
	var Namayangdicari string

	cetakData(*A, n, JumlahNominal)

	fmt.Print("Masukkan Nama Yang Mau Dicari : ")
	fmt.Scan(&Namayangdicari)
	fmt.Println("|====================================================|")

	ketemu = -1
	k = 1
	for ketemu == -1 && k <= n {
		if A[k].Nama == Namayangdicari {
			ketemu = k
		}
		k = k + 1
	}

	if ketemu != -1 {
		fmt.Println("|====================================================|")
		fmt.Println("| Data yang di cari Ditemukan!!                      |")
		fmt.Println("|====================================================|")
		fmt.Printf("| %-5d | %-10s | %-10d| %-17s |\n", ketemu+1, A[ketemu].Nama, A[ketemu].riwayat_trasnsaksi, A[ketemu].status)
		fmt.Println("|====================================================|")
	} else {
		fmt.Println("|====================================================|")
		fmt.Println("| Data yang di cari Tidak Ditemukan!!!")
		fmt.Println("|===================================================|")
	}
}

//Function untuk mencari data dengan metode Binary Search
func BinarySearch(A *arrData, n, JumlahNominal int) {
	var kiri, kanan, tengah, hasil int
	var YangDicari string
	kiri = 0
	kanan = n
	hasil = -1

	cetakData(*A, n, JumlahNominal)
	fmt.Println("")

	fmt.Print("Masukkan Nama Yang Mau Dicari :")
	fmt.Scan(&YangDicari)
	fmt.Println("|====================================================|")

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if YangDicari == A[tengah].Nama {
			hasil = tengah
			break
		} else if YangDicari < A[tengah].Nama {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if hasil != -1 {
		fmt.Printf("|%-52s|\n", "Data Yang Di Cari Ditemukan!!")
		fmt.Println("|====================================================|")
		fmt.Printf("| %-5d | %-10s | %-10d| %-17s |\n", hasil+1, A[hasil].Nama, A[hasil].riwayat_trasnsaksi, A[hasil].status)
		fmt.Println("|====================================================|")
	} else {
		fmt.Println("|====================================================|")
		fmt.Println("Data Yang Di Cari Tidak Ditemukan Di Dalam Data!!")
		fmt.Println("|====================================================|")
	}

}

func MenghapusData(A *arrData, n *int, JumlahNominal int) {
	var NomorData int
	cetakData(*A, *n, JumlahNominal)

	fmt.Println("")
	fmt.Print("Masukkan Nomor Data Yang Ingin Dihapus : ")
	fmt.Scan(&NomorData)

	for i := NomorData; i < *n; i++ {
		(*A)[i] = (*A)[i+1]
	}

	*n = *n - 1

	fmt.Printf("|===================================================|\n")
	fmt.Printf("|%-51s|\n", "Hasil Data Yang Telah Di Ubah")

	cetakData(*A, *n, JumlahNominal)
}

func menambahnominal(A *arrData, n, JumlahNominal int) {
	var NomorData, TambahIsiData int

	cetakData(*A, n, JumlahNominal)

	fmt.Println("")
	fmt.Print("Masukkan Nomor Data Yang Ingin Di Tambah Pada Nominalnya : ")
	fmt.Scan(&NomorData)

	for i := 1; i <= n; i++ {
		if NomorData == i {
			fmt.Println("Masukkan Nomilah Yang Masuk")
			fmt.Scan(&TambahIsiData)
			A[i].riwayat_trasnsaksi = A[i].riwayat_trasnsaksi + TambahIsiData
		}
	}

	fmt.Printf("|===================================================|\n")
	fmt.Printf("|%-51s|\n", "Data Setelah Nominal Yang Massuk")
	cetakData(*A, n, JumlahNominal)
}
