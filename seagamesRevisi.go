// Pengguna bisa melakukan penambahan, pengeditan dan penghapusan negara peserta seagames
// pengguna bisa melakukan penambahan, sekaligus pengubahan, dan penghapusan data medali yang diperoleh negara peserta
// pengguna bisa menampilkan daftar peringkat negara peserta terurut berdasarkan emas tertinggi, selanjutnya perak, dan perunggu
package main

import "fmt"

const NMAX int = 100

type seagames struct {
	Gmedal, Smedal, Bmedal, length int
	CountryName                    string
}
type tabCtry [NMAX]seagames

func main() {
	var no, nC int
	var idx int
	var country tabCtry
	for no != 7 {
		menu()
		fmt.Print("Pilih nomor yang anda inginkan: ")
		fmt.Scan(&no)
		if no == 1 {
			addCountry(&country, &nC)
		} else if no == 2 {
			editCountry(&country, &nC)
		} else if no == 3 {
			addMedal(&country, nC, idx)
		} else if no == 4 {
			resetMedal(&country, nC, idx)
		} else if no == 5 {
			leaderboard(country, nC)
		} else if no == 6 {
			tampilNegara(country, nC)
		}
	}
	fmt.Println("Terimakasih sudah menggunakan aplikasi ini")
}

func menu() {
	fmt.Println("-----------------------------------------------------")
	fmt.Println("         Selamat datang di aplikasi sea games        ")
	fmt.Println("         1. Tambah Negara Peserta                   ")
	fmt.Println("         2. Edit Negara Peserta                     ")
	fmt.Println("         3. Input Medali                            ")
	fmt.Println("         4. Reset Medali                            ")
	fmt.Println("         5. Leaderboard                             ")
	fmt.Println("         6. Cari Negara Berdasarkan Medal           ")
	fmt.Println("         7. Keluar                                  ")
	fmt.Println("-----------------------------------------------------")
}

func addCountry(C *tabCtry, nC *int) {
	fmt.Println("-----------------------------------------------------")
	var nAdd int
	var answer string
	fmt.Print("Berapa negara yang akan anda tambah? ")
	fmt.Scan(&nAdd)
	fmt.Println("Silahkan tambah negara yang ingin anda masukkan")
	for i := 0; i < nAdd; i++ {
		fmt.Scan(&C[*nC].CountryName)
		*nC++
	}
	fmt.Println("-----------------------------------------------------")
	fmt.Println(" Apakah anda ingin menambahkan negara yang lainnya?  ")
	fmt.Println("                      Ya/Tidak                       ")
	fmt.Println("    (Tolong sesuaikan kapital sesuai instruksi)      ")
	fmt.Scan(&answer)

	if answer == "Ya" {
		addCountry(C, nC)
	}
}

func editCountry(C *tabCtry, nC *int) {
	var YN string
	var idx, answer int
	if *nC == 0 {
		fmt.Println("Tidak ada negara peserta yang bisa di edit")
		fmt.Println("kembali ke menu?")
		fmt.Println("    Ya/Tidak    ")
		fmt.Scan(&YN)
	} else {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("            Pilih negara yang akan di edit           ")
		for i := 0; i < *nC; i++ {
			C[i].length = len(C[i].CountryName)
		}
		for k := 1; k < *nC; k++ {
			temp := C[k]
			j := k - 1
			for j >= 0 && C[j].length > temp.length {
				C[j+1] = C[j]
				j--
			}
			C[j+1] = temp
		}
		for i := 0; i < *nC; i++ {
			fmt.Print(i+1, ". ", C[i].CountryName, "\n")
		}
		fmt.Print(*nC+1, ". kembali ke menu\n")
		fmt.Scan(&idx)
		idx -= 1
		if idx == *nC {
			return
		} else {
			fmt.Println("-----------------------------------------------------")
			fmt.Println("              Pilih apa yang akan di edit            ")
			fmt.Println("              1. Ubah Nama                           ")
			fmt.Println("              2. Hapus negara peserta                ")
			fmt.Println("              3. kembali                             ")
			fmt.Println("-----------------------------------------------------")
			fmt.Scan(&answer)
			if answer == 1 {
				renameCountry(C, *nC, idx)
			} else if answer == 2 {
				delCountry(C, nC, &idx)
			}
		}
	}
}

func delCountry(C *tabCtry, nC *int, idx *int) {
	var YN string
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Apakah anda yakin ingin menghapus negara peserta ini?")
	fmt.Println("                      Ya/Tidak                       ")
	fmt.Println("    (Tolong sesuaikan kapital sesuai instruksi)      ")
	fmt.Scan(&YN)
	if YN == "Ya" {
		for j := *idx + 1; j < *nC; j++ {
			C[j-1] = C[j]
		}
		*nC--
		fmt.Println("-----------------------------------------------------")
		fmt.Println("           Negara peserta berhasil dihapus           ")
		fmt.Println("-----------------------------------------------------")
		fmt.Println("    Apakah anda ingin mengedit negara peserta lain?  ")
		fmt.Println("                      Ya/Tidak                       ")
		fmt.Println("     (Tolong sesuaikan kapital sesuai instruksi)     ")
		fmt.Scan(&YN)
		if YN == "Ya" {
			editCountry(C, nC)
		}
	} else if YN == "Tidak" {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("    Apakah anda ingin mengedit negara peserta lain?  ")
		fmt.Println("                      Ya/Tidak                       ")
		fmt.Println("     (Tolong sesuaikan kapital sesuai instruksi)     ")
		fmt.Scan(&YN)
		if YN == "Ya" {
			editCountry(C, nC)
		}
	}
}

func leaderboard(C tabCtry, nC int) {
	var YN string
	var show int
	fmt.Println("-----------------------------------------------------")
	fmt.Println("          Pilih tampilan leaderboard                 ")
	fmt.Println("          1. Berdasarkan emas tertinggi              ")
	fmt.Println("          2. Berdasarkan perak tertinggi             ")
	fmt.Println("          3. Berdasarkan perunggu tertinggi          ")
	fmt.Println("          4. Berdasarkan keseluruhan tertinggi       ")
	fmt.Println("          5. Kembali                                 ")
	fmt.Print("Input: ")
	fmt.Scan(&show)
	if show == 5 {
		return
	}
	fmt.Println("-----------------------------------------------------")
	if nC == 0 {
		fmt.Println("              Data negara masih kosong               ")
		fmt.Println("-----------------------------------------------------")
	} else {
		if show == 1 {
			sortingGold(&C, nC)
			fmt.Println("Berikut adalah Peringkat negara yang memperoleh medali emas tertinggi")
			fmt.Printf("%-20s %20s %20s %20s \n", "Nama Negara", "Medali emas", "Medali perak", "Medali Perunggu")
			for i := 0; i < nC; i++ {
				fmt.Printf("%-20s %15d %19d %18d \n", C[i].CountryName, C[i].Gmedal, C[i].Smedal, C[i].Bmedal)
			}
		} else if show == 2 {
			sortingSilver(&C, nC)
			fmt.Println("Berikut adalah Peringkat negara yang memperoleh medali Perak tertinggi")
			fmt.Printf("%-20s %20s %20s %20s \n", "Nama Negara", "Medali emas", "Medali perak", "Medali Perunggu")
			for i := 0; i < nC; i++ {
				fmt.Printf("%-20s %15d %19d %18d \n", C[i].CountryName, C[i].Gmedal, C[i].Smedal, C[i].Bmedal)
			}
		} else if show == 3 {
			sortingBronze(&C, nC)
			fmt.Println("Berikut adalah Peringkat negara yang memperoleh medali perunggu tertinggi")
			fmt.Printf("%-20s %20s %20s %20s \n", "Nama Negara", "Medali emas", "Medali perak", "Medali Perunggu")
			for i := 0; i < nC; i++ {
				fmt.Printf("%-20s %15d %19d %18d \n", C[i].CountryName, C[i].Gmedal, C[i].Smedal, C[i].Bmedal)
			}
		} else if show == 4 {
			sortingMedal(&C, nC)
			fmt.Println("Berikut adalah peringkat negara yang memperoleh medali tertinggi secara keseluruhan")
			fmt.Printf("%-20s %20s %20s %20s \n", "Nama Negara", "Medali emas", "Medali perak", "Medali Perunggu")
			for i := 0; i < nC; i++ {
				fmt.Printf("%-20s %15d %19d %18d \n", C[i].CountryName, C[i].Gmedal, C[i].Smedal, C[i].Bmedal)
			}
		}
		fmt.Println("-----------------------------------------------------")
	}
	fmt.Println("                  kembali ke menu?                   ")
	fmt.Println("                      Ya/Tidak                       ")
	fmt.Println("     (Tolong sesuaikan kapital sesuai instruksi)     ")
	fmt.Scan(&YN)
	if YN == "Ya" {
		return
	} else if YN == "Tidak" {
		leaderboard(C, nC)
	}
}

func renameCountry(C *tabCtry, nC, idx int) {
	var rename, YN string
	fmt.Print("Silahkan ubah nama: ")
	fmt.Scan(&rename)
	if rename == C[idx].CountryName {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Nama negara peserta masih sama, tidak ada yang diubah")
	} else {
		C[idx].CountryName = rename
		fmt.Println("-----------------------------------------------------")
		fmt.Println("              Nama negara berhasil diubah            ")
	}
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Apakah anda ingin mengedit negara peserta yang lainnya?")
	fmt.Println("                      Ya/Tidak                       ")
	fmt.Println("     (Tolong sesuaikan kapital sesuai instruksi)     ")
	fmt.Scan(&YN)
	if YN == "Ya" {
		editCountry(C, &nC)
	}
}

func addMedal(C *tabCtry, nC, idx int) {
	var i, medal, nMedal int
	var YN string
	if nC == 0 {
		fmt.Println("              Data negara masih kosong               ")
	} else {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("      Pilih negara yang akan di tambahkan medal      ")
		for i := 0; i < nC; i++ {
			C[i].length = len(C[i].CountryName)
		}
		for k := 1; k < nC; k++ {
			temp := C[k]
			j := k - 1
			for j >= 0 && C[j].length > temp.length {
				C[j+1] = C[j]
				j--
			}
			C[j+1] = temp
		}
		for i := 0; i < nC; i++ {
			fmt.Print(i+1, ". ", C[i].CountryName, "\n")
		}
		fmt.Print(nC+1, ". kembali ke menu\n")
		fmt.Scan(&idx)
		idx -= 1
		if idx == nC {
			return
		} else {
			fmt.Println("-----------------------------------------------------")
			fmt.Println("           Berapa medali yang akan ditambah?         ")
			fmt.Scan(&nMedal)
			fmt.Println("-----------------------------------------------------")
			fmt.Println("       Pilih jenis medali yang akan ditambah      ")
			fmt.Println("       1. Emas                                    ")
			fmt.Println("       2. Perak                                   ")
			fmt.Println("       3. Perunggu                                ")
			fmt.Println("-----------------------------------------------------")
			for i < nMedal {
				fmt.Scan(&medal)
				if medal == 1 {
					C[idx].Gmedal++
				} else if medal == 2 {
					C[idx].Smedal++
				} else if medal == 3 {
					C[idx].Bmedal++
				}
				i++
			}
		}
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Medali berhasil ditambahkan")
	}
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Apakah anda ingin menambahkan medali negara peserta yang lainnya?")
	fmt.Println("                      Ya/Tidak                       ")
	fmt.Println("     (Tolong sesuaikan kapital sesuai instruksi)     ")
	fmt.Scan(&YN)
	if YN == "Ya" {
		addMedal(C, nC, idx)
	}
}

func resetMedal(C *tabCtry, nC, idx int) {
	if nC == 0 {
		fmt.Println("              Data negara masih kosong               ")
	} else {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("      Pilih negara yang akan di reset medalinya      ")
		for i := 0; i < nC; i++ {
			C[i].length = len(C[i].CountryName)
		}
		for k := 1; k < nC; k++ {
			temp := C[k]
			j := k - 1
			for j >= 0 && C[j].length > temp.length {
				C[j+1] = C[j]
				j--
			}
			C[j+1] = temp
		}
		for i := 0; i < nC; i++ {
			fmt.Print(i+1, ". ", C[i].CountryName, "\n")
		}
		fmt.Print(nC+1, ". kembali ke menu\n")
		fmt.Scan(&idx)
		idx -= 1
		if idx == nC {
			return
		} else {
			C[idx].Gmedal = 0
			C[idx].Smedal = 0
			C[idx].Bmedal = 0
			fmt.Println("Medali berhasil di reset")
		}
	}
}

func sortingMedal(C *tabCtry, n int) {
	var i, idx int
	for i = 1; i < n; i++ {
		idx = i - 1
		maxGold(C, n, i, &idx)
		swap(C, idx, i)
	}
	for i = 1; i < n; i++ {
		idx = i - 1
		maxSilverAll(C, n, i, &idx)
		swap(C, idx, i)
	}
	for i = 1; i < n; i++ {
		idx = i - 1
		maxBronzeAll(C, n, i, &idx)
		swap(C, idx, i)
	}
}

func sortingGold(C *tabCtry, n int) {
	var i, idx int
	for i = 1; i < n; i++ {
		idx = i - 1
		maxGold(C, n, i, &idx)
		swap(C, idx, i)
	}
}

func sortingSilver(C *tabCtry, n int) {
	var i, idx int
	for i = 1; i < n; i++ {
		idx = i - 1
		maxSilv(C, n, i, &idx)
		swap(C, idx, i)
	}
}

func sortingBronze(C *tabCtry, n int) {
	var i, idx int
	for i = 1; i < n; i++ {
		idx = i - 1
		maxBR(C, n, i, &idx)
		swap(C, idx, i)
	}
}

func swap(C *tabCtry, idx, j int) {
	temp := C[idx]
	C[idx] = C[j-1]
	C[j-1] = temp
}

func maxGold(C *tabCtry, n, j int, idx *int) {
	for i := j; i < n; i++ {
		if C[*idx].Gmedal < C[i].Gmedal {
			*idx = i
		}
	}
}

func maxSilv(C *tabCtry, n, j int, idx *int) {
	for i := j; i < n; i++ {
		if C[*idx].Smedal < C[i].Smedal {
			*idx = i
		}
	}
}

func maxBR(C *tabCtry, n, j int, idx *int) {
	for i := j; i < n; i++ {
		if C[*idx].Bmedal < C[i].Bmedal {
			*idx = i
		}
	}
}

func maxSilverAll(C *tabCtry, n, j int, idx *int) {
	for i := j; i < n; i++ {
		if C[*idx].Gmedal == C[i].Gmedal {
			if C[*idx].Smedal < C[i].Smedal {
				*idx = i
			}
		}
	}
}
func maxBronzeAll(C *tabCtry, n, j int, idx *int) {
	for i := j; i < n; i++ {
		if C[*idx].Gmedal == C[i].Gmedal {
			if C[*idx].Smedal == C[i].Smedal {
				if C[*idx].Bmedal < C[i].Bmedal {
					*idx = i
				}
			}
		}
	}
}

func tampilNegara(C tabCtry, nC int) {
	var idx int
	var many int
	if nC == 0 {
		fmt.Println("              Data negara masih kosong               ")
	} else {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("       Pilih jenis medali yang akan ditambah      ")
		fmt.Println("       1. Emas                                    ")
		fmt.Println("       2. Perak                                   ")
		fmt.Println("       3. Perunggu                                ")
		fmt.Println("       4. kembali                                ")
		fmt.Println("-----------------------------------------------------")
		fmt.Scan(&idx)
		if idx == 4 {
			return
		} else if idx == 1 {
			fmt.Println("-----------------------------------------------------")
			fmt.Println("         berapa medali emas yang anda cari?          ")
			fmt.Scan(&many)
			sortingGold(&C, nC)
			hasil := GoldSearch(C, nC, many)
			fmt.Println("-----------------------------------------------------")
			if hasil == -1 {
				fmt.Println("Maaf, tidak ada negara yang mendapatkan", many, "medali emas")
			} else {
				fmt.Println("Baiklah, berikut adalah negara yang mendapatkan", many, "medali emas yang pertama kami dapatkan")
				fmt.Println("Negara", C[hasil].CountryName, "dengan", many, "medali emas")
			}
		} else if idx == 2 {
			fmt.Println("-----------------------------------------------------")
			fmt.Println("         berapa medali perak yang anda cari?          ")
			fmt.Scan(&many)
			sortingSilver(&C, nC)
			hasil := SilvSearch(C, nC, many)
			fmt.Println("-----------------------------------------------------")
			if hasil == -1 {
				fmt.Println("Maaf, tidak ada negara yang mendapatkan", many, "medali perak")
			} else {
				fmt.Println("Baiklah, berikut adalah negara yang mendapatkan", many, "medali perak yang pertama kami dapatkan")
				fmt.Println("Negara", C[hasil].CountryName, "dengan", many, "medali perak")
			}
		} else if idx == 3 {
			fmt.Println("-----------------------------------------------------")
			fmt.Println("         berapa medali perunggu yang anda cari?          ")
			fmt.Scan(&many)
			sortingBronze(&C, nC)
			hasil := BRSearch(C, nC, many)
			fmt.Println("-----------------------------------------------------")
			if hasil == -1 {
				fmt.Println("Maaf, tidak ada negara yang mendapatkan", many, "medali perunggu")
			} else {
				fmt.Println("Baiklah, berikut adalah negara yang mendapatkan", many, "medali perunggu yang pertama kami dapatkan")
				fmt.Println("Negara", C[hasil].CountryName, "dengan", many, "medali perunggu")
			}
		}
	}
}

func GoldSearch(C tabCtry, nC, idx int) int {
	var hasil int
	sortingGold(&C, nC)
	hasil = -1
	L := 0
	R := nC - 1
	for R >= L && hasil == -1 {
		mid := (L + R) / 2
		if C[mid].Gmedal < idx {
			R = mid
		} else if C[mid].Gmedal > idx {
			L = mid + 1
		} else {
			hasil = mid
		}
	}
	return hasil
}

func SilvSearch(C tabCtry, nC, idx int) int {
	var hasil int
	sortingSilver(&C, nC)
	hasil = -1
	L := 0
	R := nC - 1
	for R >= L && hasil == -1 {
		mid := (L + R) / 2
		if C[mid].Smedal < idx {
			R = mid
		} else if C[mid].Smedal > idx {
			L = mid + 1
		} else {
			hasil = mid
		}
	}
	return hasil
}

func BRSearch(C tabCtry, nC, idx int) int {
	var hasil int
	sortingBronze(&C, nC)
	hasil = -1
	L := 0
	R := nC - 1
	for R >= L && hasil == -1 {
		mid := (L + R) / 2
		if C[mid].Bmedal < idx {
			R = mid
		} else if C[mid].Bmedal > idx {
			L = mid + 1
		} else {
			hasil = mid
		}
	}
	return hasil
}
