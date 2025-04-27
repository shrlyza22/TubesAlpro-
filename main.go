package main

import (
	"fmt"
	"strings"
)

// Struktur untuk menyimpan ide
type Ide struct {
	noIde        int
	ideNya       string
	pengusul     string
	deskripsiIde string
	status       string
}

// Slice untuk menyimpan daftar ide
const NMAX int = 10

var daftarIde [NMAX]Ide

var n int = 0

// var idePengusul daftarIde

func main() {
	var pilihan int

	for {
		fmt.Println("\n=== Aplikasi Pengelolaan Ide ===")
		fmt.Println("1. Tambah Ide")
		fmt.Println("2. Lihat Semua Ide")
		fmt.Println("3. Cari Ide")
		fmt.Println("4. Hapus Ide")
		fmt.Println("5. Edit Ide")
		fmt.Println("6. Keluar")
		fmt.Println("================================")
		fmt.Print("Pilih menu (1/2/3/4/5/6) ?: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			inputIdeBaru(&daftarIde, &n)
		case 2:
			tampilSemuaIdePengusul(daftarIde, n)
		case 3:
			cariIdenya(daftarIde, n)
		case 4:
			hapusIde(&daftarIde, &n)
		case 5:
			editIde(daftarIde, n)
		case 6:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func inputIdeBaru(daftarIde *[NMAX]Ide, n *int) {
	// I.S  : daftarIde terdefinisi, n terdefinisi
	//  F.S  : Menampilkan ide baru yang ingin ditambahkan ke dalam daftar ide
	var i int
	var tempNomorIde int

	fmt.Println("\n=== Tambah Ide Baru ===")

	for {
		fmt.Print("Masukkan nomor ide: ")
		fmt.Scan(&tempNomorIde)

		if tempNomorIde <= 0 {
			fmt.Println("Nomor ide harus lebih dari 0. Silakan coba lagi.")
			continue
		}

		// Cek kesamaan nomor ide
		duplicate := false
		for i := 0; i < *n; i++ {
			if daftarIde[*n].noIde == tempNomorIde {
				duplicate = true
				break
			}
		}

		if duplicate {
			fmt.Println("-----PERINGATAN-----")
			fmt.Println("Nomor ide sudah ada. Silakan coba lagi!!")
		} else {
			break
		}
	}
	//menambahkan daftar nomor ke dalam array
	daftarIde[*n].noIde = tempNomorIde

	fmt.Print("Judul (gunakan_underscore_untuk_spasi): ")
	fmt.Scan(&daftarIde[i].ideNya)
	fmt.Print("Deskripsi (gunakan_underscore_untuk_spasi): ")
	fmt.Scan(&daftarIde[i].deskripsiIde)
	fmt.Print("Pengusul: ")
	fmt.Scan(&daftarIde[i].pengusul)
	fmt.Print("Kategori: ")
	fmt.Scan(&daftarIde[i].status)

	(*n)++
	fmt.Print("Data berhasil ditambahkan!")
}

func tampilSemuaIdePengusul(daftarIde [NMAX]Ide, n int) {
	// I.S  : daftarIde terdefinisi, n terdefinisi
	// F.S  : Menampilkan semua ide yang diberikan oleh pengusul
	// fmt.Println("\n=== Daftar Semua Ide Pengusul ===")
	if n == 0 {
		fmt.Println("Belum ada ide yang ditambahkan.")
		return
	}
	fmt.Println("\n==== Daftar Semua Ide Pengusul ====")
	for i := 0; i < n; i++ {
		if daftarIde[n].noIde > 0 {

		}
		fmt.Println("============================================")
		fmt.Printf("Ide ke - %d\n", i+1)
		fmt.Println("============================================")
		fmt.Printf("Nomor Ide : %d\n", daftarIde[n].noIde)
		fmt.Println("Judul    :", strings.ReplaceAll(daftarIde[i].ideNya, "_", " "))
		fmt.Println("Deskripsi:", strings.ReplaceAll(daftarIde[i].deskripsiIde, "_", " "))
		fmt.Println("Pengusul :", daftarIde[i].pengusul)
		fmt.Println("Kategori :", daftarIde[i].status)
		fmt.Println("--------------------------------------------")

	}
}

func cariIdenya(daftarIde [NMAX]Ide, n int) {
	// I.S  : daftarIde terdefinisi, n terdefinisi
	// F.S  : Mencari ide yang ingin pengguna temukan berdasarkan kata kunci yang diberikan
	// var keyword string
	// fmt.Println("\n=== Cari Ide ===")
	// fmt.Print("Masukkan kata kunci: ")
	// fmt.Scanln(&keyword, " ")

	// if keyword == " " {
	// 	fmt.Println("Kata kunci tidak boleh kosong.")
	// 	return
	// }

	// fmt.Println("\nHasil Pencarian Ide:")
	// found := false
	// for i := 0; i < n; i++ {
	// 	if strings.Contains(strings.ToLower(daftarIde[i].ideNya), strings.ToLower(keyword)) ||
	// 		strings.Contains(strings.ToLower(daftarIde[i].deskripsiIde), strings.ToLower(keyword)) ||
	// 		strings.Contains(strings.ToLower(daftarIde[i].pengusul), strings.ToLower(keyword)) ||
	// 		strings.Contains(strings.ToLower(daftarIde[i].status), strings.ToLower(keyword)) {

	// 		fmt.Printf(" #%d. %d. %s - %s (Pengusul: %s, Kategori: %s)\n", daftarIde[i].noIde,
	// 			daftarIde[i].ideNya, daftarIde[i].deskripsiIde,
	// 			daftarIde[i].pengusul, daftarIde[i].status)
	// 		found = true
	// 	}
}

//		if !found {
//			fmt.Println("Tidak ada ide yang cocok dengan kata kunci yang di cari Anda.")
//		}

func hapusIde(daftarIde *[NMAX]Ide, n *int) {
	// I.S  : daftarIde terdefinisi, n terdefinisi
	// F.S  : Menghapus ide yang ingin dihapus dari daftar ide

}
func editIde(daftarIde [NMAX]Ide, n int) {
	// I.S  : daftarIde terdefinisi, n terdefinisi
	// F.S  : Mengedit ide yang ingin diedit dari daftar ide

}
