package main

import "fmt"

type Ide struct {
	noIde        int
	ideNya       string
	pengusul     string
	deskripsiIde string
	status       string
}

const NMAX int = 100

type tabIde [NMAX]Ide

func main() {
	var n int = 0
	var data tabIde
	home(&data, &n)
}
func home(A *tabIde, n *int) {
	var opsi, no, idx int
	for opsi != 5 {
		fmt.Println("===================================================")
		fmt.Println("||                                               ||")
		fmt.Println("||       Sheren Aulia Azahra || 103032400036     ||")
		fmt.Println("||       Ririn Nur Aini      || 103032400054     ||")
		fmt.Println("||                                               ||")
		fmt.Println("===================================================")
		fmt.Println("Selamat datang Di Aplikasi Pengelolaan Ide")
		fmt.Println("Pilih Opsi berikut:")
		fmt.Println("1. Menambah Ide")
		fmt.Println("2. Menghapus Ide")
		fmt.Println("3. Mengedit Ide")
		fmt.Println("4. Menampilkan Data")
		fmt.Println("5. Keluar")
		fmt.Println("===================================================")
		fmt.Print("Masukkan Opsi: ")
		fmt.Scan(&opsi)
		if opsi == 1 {
			tambah_ide(&*A, &*n)
		} else if opsi == 2 {
			fmt.Println("Masukkan No Ide yang mau dihapus")
			fmt.Scan(&no)
			ascending_insertion_no_ide(&*A, *n)
			idx = cari_no_ide(*A, *n, no)
			if idx != -1 {
				hapus_ide(&*A, &*n, idx)
			} else {
				fmt.Println("Data tidak ditemukan")
			}
		} else if opsi == 3 {
			fmt.Println("Masukkan No Ide yang mau diedit")
			fmt.Scan(&no)
			idx = cari_no_ide(*A, *n, no)
			if idx != -1 {
				edit_ide(&*A, idx)
			} else {
				fmt.Println("Data tidak ditemukan")
			}
		} else if opsi == 4 {
			main_display(*A, *n)
		} else if opsi < 1 || opsi > 5 {
			fmt.Println("Opsi Invalid")
		}
	}
}
func tambah_ide(A *tabIde, n *int) {
	var no int
	fmt.Println("Masukkan Data yang Ingin ditambahkan")
	fmt.Print("No Ide:")
	fmt.Scan(&no)
	x := cari_no_ide(*A, *n, no)
	for x != -1 {
		fmt.Println("Nomor ide sudah ada")
		fmt.Print("No Ide:")
		fmt.Scan(&no)
		x = cari_no_ide(*A, *n, no)
	}
	A[*n].noIde = no
	fmt.Print("Nama Ide:")
	fmt.Scan(&A[*n].ideNya)
	fmt.Print("Nama Pengusul:")
	fmt.Scan(&A[*n].pengusul)
	fmt.Print("Deskripsi:")
	fmt.Scan(&A[*n].deskripsiIde)
	fmt.Print("Status:")
	fmt.Scan(&A[*n].status)
	*n++
	fmt.Println("Data Berhasil ditambahkan")
}
func cari_no_ide(A tabIde, n int, x int) int {
	// Data Ascending
	var left, mid, right int
	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2
		if A[mid].noIde == x {
			return mid
		} else if x < A[mid].noIde {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func cari_nama_ide(A tabIde, n int, x string) int {
	for i := 0; i < n; i++ {
		if A[i].ideNya == x {
			return i
		}
	}
	return -1
}
func ascending_insertion_no_ide(A *tabIde, n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if A[j-1].noIde > A[j].noIde {
				A[j-1], A[j] = A[j], A[j-1]
			}
			j = j - 1
		}
	}
}
func descending_insertion_no_ide(A *tabIde, n int) {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if A[j-1].noIde < A[j].noIde {
				A[j-1], A[j] = A[j], A[j-1]
			}
			j = j - 1
		}
	}
}
func ascending_selection_nama_ide(A *tabIde, n int) {
	var i, pass, idx_min int
	for pass = 0; pass <= n-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].ideNya < A[idx_min].ideNya {
				idx_min = i
			}
		}
		A[pass], A[idx_min] = A[idx_min], A[pass]
	}
}
func descending_selection_nama_ide(A *tabIde, n int) {
	var i, pass, idx_max int
	for pass = 0; pass <= n-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].ideNya > A[idx_max].ideNya {
				idx_max = i
			}
		}
		A[pass], A[idx_max] = A[idx_max], A[pass]
	}
}
func hapus_ide(A *tabIde, n *int, idx int) {
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		display_ide(*A, idx)
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data Ide Di Atas Telah Dihapus")
	}
}
func display_ide(A tabIde, idx int) {
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		fmt.Println("No Ide          : ", A[idx].noIde)
		fmt.Println("Nama Ide        : ", A[idx].ideNya)
		fmt.Println("Pengusul Ide    : ", A[idx].pengusul)
		fmt.Println("Deskripsi       : ", A[idx].deskripsiIde)
		fmt.Println("Status          : ", A[idx].status)
		fmt.Println()
	}
}
func all_display(A tabIde, n int) {
	for i := 0; i < n; i++ {
		display_ide(A, i)
	}
}
func edit_ide(A *tabIde, idx int) {
	var no int
	fmt.Println("Masukkan Data yang Ingin diedit")
	fmt.Print("No Ide:")
	fmt.Scan(&no)
	x := cari_no_ide(*A, idx, no)
	for x != -1 {
		fmt.Println("Nomor ide sudah ada")
		fmt.Print("No Ide:")
		fmt.Scan(&no)
		x = cari_no_ide(*A, idx, no)
	}
	A[idx].noIde = no
	fmt.Print("Nama Ide:")
	fmt.Scan(&A[idx].ideNya)
	fmt.Print("Nama Pengusul:")
	fmt.Scan(&A[idx].pengusul)
	fmt.Print("Deskripsi:")
	fmt.Scan(&A[idx].deskripsiIde)
	fmt.Print("Status:")
	fmt.Scan(&A[idx].status)
	fmt.Println("Data Berhasil diedit")
}
func main_display(A tabIde, n int) {
	var opsi, idx, no int
	var nama string
	fmt.Println("1. Nomor Ide Menaik")
	fmt.Println("2. Nomor Ide Menurun")
	fmt.Println("3. Nama Ide Menaik")
	fmt.Println("4. Nama Ide Menurun")
	fmt.Println("5. Cari Nama Ide")
	fmt.Println("6. Cari Nomor Ide")
	fmt.Print("Masukkan Opsi :")
	fmt.Scan(&opsi)
	for opsi < 1 || opsi > 6 {
		fmt.Print("Input Invalid")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		ascending_insertion_no_ide(&A, n)
		all_display(A, n)
	} else if opsi == 2 {
		descending_insertion_no_ide(&A, n)
		all_display(A, n)
	} else if opsi == 3 {
		ascending_selection_nama_ide(&A, n)
		all_display(A, n)
	} else if opsi == 4 {
		descending_selection_nama_ide(&A, n)
		all_display(A, n)
	} else if opsi == 5 {
		fmt.Println("Masukkan Nama Ide yang akan dicari")
		fmt.Scan(&nama)
		idx = cari_nama_ide(A, n, nama)
		if idx != -1 {
			display_ide(A, idx)
		} else {
			fmt.Println("Ide tidak ditemukan")
		}
	} else if opsi == 6 {
		fmt.Println("Masukkan No Ide yang akan dicari")
		fmt.Scan(&no)
		idx = cari_no_ide(A, n, no)
		if idx != -1 {
			display_ide(A, idx)
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}
}
