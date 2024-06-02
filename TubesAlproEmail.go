package main

import (
	"fmt"
)

const NMAX = 100000

type databaseEmail struct {
	nama, gender, username, password string
	tglLahir, blnLahir, thnLahir     int
	sender, receiver                 string
	pesanMasuk, balasan              string
}

type beforePersetujuan struct {
	nama, gender, username, password string
	tglLahir, blnLahir, thnLahir     int
	sender, receiver                 string
	pesanMasuk, balasan              string
}

type tabEmail [NMAX]databaseEmail
type tabAgree [NMAX]beforePersetujuan

func registerEmail(A *tabEmail, B *tabAgree, n int, usernameDiCari *string) {
	var tempFindB, tempFind, i int

	fmt.Println()

	fmt.Println("Registrasi email")

	fmt.Print("Masukkan nama anda: ")
	fmt.Scan(&B[n].nama)
	fmt.Println()
	fmt.Print("Masukkan gender anda: ")
	fmt.Scan(&B[n].gender)
	fmt.Println()
	fmt.Print("Masukkan tanggal lahir anda: ")
	fmt.Scan(&B[n].tglLahir)
	fmt.Println()
	fmt.Print("Masukkan bulan lahir anda (angka 1-12): ")
	fmt.Scan(&B[n].blnLahir)
	fmt.Println()
	fmt.Print("Masukkan tahun lahir anda: ")
	fmt.Scan(&B[n].thnLahir)
	fmt.Println()
	fmt.Print("Silahkan buat username yang anda inginkan: ")
	fmt.Scan(usernameDiCari)
	B[n].username = *usernameDiCari

	tempFindB = searchUsernameB(A, B, &n, i, *usernameDiCari)
	tempFind = searchUsername(A, &n, i, *usernameDiCari)
	fmt.Println()
	if (tempFind > -1 && tempFind != n) || (tempFindB > -1 && tempFindB != n) {
		for {
			fmt.Println("Username yang anda inginkan telah digunakan")
			fmt.Print("Silahkan masukkan kembali username yang anda inginkan: ")
			fmt.Scan(usernameDiCari)
			B[n].username = *usernameDiCari
			tempFindB = searchUsernameB(A, B, &n, i, *usernameDiCari)
			tempFind = searchUsername(A, &n, i, *usernameDiCari)
			if tempFind == -1 && tempFindB == -1 {
				break
			}
		}
	}

	fmt.Print("Masukkan password anda: ")
	fmt.Scan(&B[n].password)
	fmt.Println()
	if len(B[n].password) < 8 {
		for {
			fmt.Println("Password yang anda masukkan tidak valid")
			fmt.Println("Password yang valid adalah terdiri dari 8 karakter")
			fmt.Print("Silahkan masukkan kembali password yang anda inginkan: ")
			fmt.Scan(&B[n].password)
			if len(B[n].password) >= 8 {
				break
			}
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println()

	var jawabanMakeSure, jawabsetujugak string

	fmt.Println("Berikut adalah data yang sudah anda inputkan")
	fmt.Println("Periksa kembali apakah data berikut sudah benar")
	fmt.Println()

	fmt.Println("Data email ke-", (n + 1))

	fmt.Print("Nama lengkap: ")
	fmt.Println(B[n].nama)
	fmt.Println()
	fmt.Print("Gender: ")
	fmt.Println(B[n].gender)
	fmt.Println()
	fmt.Print("tanggal/bulan/tahun: ")
	fmt.Println(B[n].tglLahir, "/", B[n].blnLahir, "/", B[n].thnLahir)
	fmt.Println()
	fmt.Print("username: ")
	fmt.Println(B[n].username)
	fmt.Println()

	fmt.Println("Apakah data tersebut sudah benar?")
	fmt.Println("A. Sudah")
	fmt.Println("B. Registrasi ulang")
	fmt.Scan(&jawabanMakeSure)
	fmt.Println()
	if jawabanMakeSure == "B" {
		fmt.Println()
		registerEmail(A, B, n, usernameDiCari)
	} else if jawabanMakeSure == "A" {
		fmt.Println()
		fmt.Println("Dengan ini anda setuju dengan segala persyaratan dan peraturan yang ada")
		fmt.Println("A. Setuju")
		fmt.Println("B. Tidak setuju")
		fmt.Scan(&jawabsetujugak)
		if jawabsetujugak == "A" {
			fmt.Println("|__  PENTING !!!! ____________________________________________________________________________________|")
			fmt.Println("| Admin akan segera memberikan persetujuan akun yang anda buat.                                       |")
			fmt.Println("| Jika saat anda tetap tidak bisa login, maka data yang anda isi tidak valid.                         |")
			fmt.Println("| Dan jika anda mencoba log in sebelum akun ada disetujui oleh admin, maka registrasi anda akan gagal |")
			fmt.Println("| Silahkan registrasi ulang jika hal tersebut terjadi pada anda                                       |")
			fmt.Println("|_____________________________________________________________________________________________________|")
			adaAdmin(A, B, &n)

		} else if jawabsetujugak == "B" {
			fmt.Println()
			fmt.Println("Registrasi gagal")
			menuAwal(A, n)
		}
	}
}

func adaAdmin(A *tabEmail, B *tabAgree, n *int) {
	var jawab, usernameDiCari, passwordDiCari string
	var posisiLogin int
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("------------------------   |Aplikasi Email|   ------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("A. Register akun")
	fmt.Println("B. Log in akun")
	fmt.Println("C. Log in sebagai admin")
	fmt.Scan(&jawab)
	if jawab == "A" {
		registerEmail(A, B, *n, &usernameDiCari)
	} else if jawab == "B" {
		loginEmail(A, *n, &posisiLogin, &usernameDiCari, &passwordDiCari)
	} else if jawab == "C" {
		persetujuanAdmin(A, B, *n)
	}
}

func persetujuanAdmin(A *tabEmail, B *tabAgree, n int) {
	var jawabAdmin string
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("-----------------------  Persetujuan by Admin  -----------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")

	fmt.Println()

	fmt.Println("Akun yang meminta persetujuan: ")
	fmt.Println()
	fmt.Print("Nama lengkap: ")
	fmt.Println(B[n].nama)
	fmt.Println()
	fmt.Print("Gender: ")
	fmt.Println(B[n].gender)
	fmt.Println()
	fmt.Print("tanggal/bulan/tahun: ")
	fmt.Println(B[n].tglLahir, "/", B[n].blnLahir, "/", B[n].thnLahir)
	fmt.Println()
	fmt.Print("username: ")
	fmt.Println(B[n].username)
	fmt.Println()
	fmt.Println("Apakah anda menyetujui data yang pengguna berikan?")
	fmt.Println("A. Setuju")
	fmt.Println("B. Tidak Setuju")
	fmt.Scan(&jawabAdmin)

	if jawabAdmin == "A" {

		A[n].nama = B[n].nama
		A[n].gender = B[n].gender
		A[n].username = B[n].username
		A[n].password = B[n].password
		A[n].tglLahir = B[n].tglLahir
		A[n].blnLahir = B[n].blnLahir
		A[n].thnLahir = B[n].thnLahir
		A[n].sender = B[n].sender
		A[n].receiver = B[n].receiver
		A[n].pesanMasuk = B[n].pesanMasuk
		A[n].balasan = B[n].balasan
		fmt.Println()
		fmt.Println("Registrasi anda diterima")
		fmt.Println()
		n++

		menuAwal(A, n)

	} else if jawabAdmin == "B" {

		fmt.Println()
		fmt.Println("Registrasi anda ditolak")
		fmt.Println()
		menuAwal(A, n)
	}
}

func searchPassword(A *tabEmail, n *int, i int, passwordDiCari string) int {
	var temp int
	temp = -1
	i = 0
	for temp == -1 && i < *n {
		if A[i].password == passwordDiCari {
			temp = i

		}
		i++

	}

	return temp
}

func searchUsername(A *tabEmail, n *int, i int, usernameDiCari string) int {
	var temp int
	temp = -1
	i = 0
	for temp == -1 && i < *n {
		if A[i].username == usernameDiCari {
			temp = i

		}
		i++

	}

	return temp
}

func searchUsernameB(A *tabEmail, B *tabAgree, n *int, i int, usernameDiCari string) int {
	var temp int
	temp = -1
	i = 0
	for temp == -1 && i < *n {
		if B[i].username == usernameDiCari {
			temp = i

		}
		i++

	}

	return temp
}

func loginEmail(A *tabEmail, n int, posisiLogin *int, usernameDiCari, passwordDiCari *string) {

	var tempCariUname, tempCariPass, i int
	var jawabanSub string
	tempCariUname = -1
	tempCariPass = -1

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------    LOG IN    ----------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")

	fmt.Println()
	fmt.Print("Silahkan masukkan Username email anda: ")
	fmt.Scan(usernameDiCari)
	fmt.Println()
	fmt.Print("Silahkan masukkan Password email anda: ")
	fmt.Scan(passwordDiCari)
	fmt.Println()

	tempCariUname = searchUsername(A, &n, i, *usernameDiCari)
	tempCariPass = searchPassword(A, &n, i, *passwordDiCari)

	if tempCariUname > -1 && tempCariPass > -1 {
		if A[tempCariUname].username == *usernameDiCari && A[tempCariPass].password == *passwordDiCari {
			fmt.Println("Log in berhasil !")

			*posisiLogin = tempCariPass
			setelahLogin(A, n, posisiLogin)
		}
	} else {
		if tempCariUname == -1 && tempCariPass == -1 {
			fmt.Println("Log in gagal !")
			fmt.Println("Username atau password anda tidak ditemukan")
			fmt.Println()
			fmt.Println("Apakah anda ingin mencoba login kembali?")
			fmt.Println("A. Iya")
			fmt.Println("B. Tidak")
			fmt.Scan(&jawabanSub)
			if jawabanSub == "A" {
				fmt.Println("Silahkan log in ulang")
				loginEmail(A, n, posisiLogin, usernameDiCari, passwordDiCari)
			} else if jawabanSub == "B" {
				menuAwal(A, n)
			}

		} else {
			fmt.Println(tempCariUname)
			fmt.Println(tempCariPass)
			fmt.Println("Log in gagal !")
			fmt.Println("Username atau password anda tidak ditemukan")
			fmt.Println()
			fmt.Println("Apakah anda ingin mencoba login kembali?")
			fmt.Println("A. Iya")
			fmt.Println("B. Tidak")
			fmt.Scan(&jawabanSub)
			if jawabanSub == "A" {
				fmt.Println("Silahkan log in ulang")
				loginEmail(A, n, posisiLogin, usernameDiCari, passwordDiCari)
			} else {
				menuAwal(A, n)
			}
		}
	}
}

func setelahLogin(A *tabEmail, n int, posisiLogin *int) {
	var jawabanMenuSetelahLogin string
	var usernameDiCari string
	fmt.Println()
	fmt.Println("Hallo,", A[*posisiLogin].nama)
	fmt.Println("Apa yang ingin anda lakukan? ")
	fmt.Println("A. Kirim Pesan")
	fmt.Println("B. Baca Pesan")
	fmt.Println("C. Hapus Pesan")
	fmt.Println("D. Log Out")
	fmt.Scan(&jawabanMenuSetelahLogin)
	if jawabanMenuSetelahLogin == "A" {
		kirimPesan(A, n, posisiLogin, &usernameDiCari)
	} else if jawabanMenuSetelahLogin == "B" {
		bacaInbox(A, n, posisiLogin, &usernameDiCari)
	} else if jawabanMenuSetelahLogin == "C" {
		hapus1Pesan(A, n, posisiLogin, &usernameDiCari)
	} else if jawabanMenuSetelahLogin == "D" {
		fmt.Println("Berhasil Log Out")
		abisLogOut(A, &n)
	}
}

func kirimPesan(A *tabEmail, n int, posisiLogin *int, usernameDiCari *string) {
	var tempCariUname, i int
	fmt.Println()
	fmt.Print("Silahkan masukkan username yang ingin anda kirim pesan: ")
	fmt.Scan(usernameDiCari)
	fmt.Println("Mencari username:", *usernameDiCari)
	fmt.Println()
	tempCariUname = searchUsername(A, &n, i, *usernameDiCari)
	if tempCariUname == -1 {
		fmt.Println("Username yang ingin anda kirim pesan tidak ditemukan")
		setelahLogin(A, n, posisiLogin)
	} else {
		A[tempCariUname].sender = A[*posisiLogin].username
		fmt.Println("Silahkan Masukkan pesan yang ingin anda kirim: ")
		fmt.Scan(&A[tempCariUname].pesanMasuk)
		fmt.Println("Pesan berhasil dikirim ke:", A[tempCariUname].username)
		setelahLogin(A, n, posisiLogin)
	}
}

func hapus1Pesan(A *tabEmail, n int, posisiLogin *int, usernameDiCari *string) {
	var tempCariUname, i int
	var jawaban string
	fmt.Println("|||||||||||||||||||||||||||||||||||||||||||||||||")
	fmt.Println("|||||||||||||||||| HAPUS PESAN ||||||||||||||||||")
	fmt.Println("|||||||||||||||||||||||||||||||||||||||||||||||||")
	fmt.Println()
	fmt.Println("Dari", A[*posisiLogin].sender, ": ")
	fmt.Println(A[*posisiLogin].pesanMasuk)
	fmt.Println()

	fmt.Println("Apakah anda yakin ingin menghapus pesan ini? ")
	fmt.Println("A. Iya")
	fmt.Println("B. Cancel")
	fmt.Scan(&jawaban)
	if jawaban == "A" {
		*usernameDiCari = A[*posisiLogin].sender
		tempCariUname = searchUsername(A, &n, i, *usernameDiCari)
		A[*posisiLogin].sender = ""

		A[*posisiLogin].pesanMasuk = ""
		fmt.Println()
		fmt.Println("Pesan dari:", A[tempCariUname].username, "berhasil dihapus :")
		setelahLogin(A, n, posisiLogin)
	} else if jawaban == "B" {
		setelahLogin(A, n, posisiLogin)
	}

}

func bacaInbox(A *tabEmail, n int, posisiLogin *int, usernameDiCari *string) {
	var jawabBacaInbox string
	var tempCariUname, i int
	fmt.Println()
	fmt.Println("Berikut adalah pesan yang ada pada email anda")
	fmt.Println()
	if A[*posisiLogin].sender == "" && A[*posisiLogin].pesanMasuk == "" {
		fmt.Println()
		fmt.Println("--------------- Tidak ada pesan yang masuk ---------------")
		fmt.Println()
		fmt.Println("A. kembali")
		fmt.Scan(&jawabBacaInbox)
		if jawabBacaInbox == "A" {
			setelahLogin(A, n, posisiLogin)
		}
	} else {
		fmt.Println("Dari", A[*posisiLogin].sender, ": ")
		fmt.Println(A[*posisiLogin].pesanMasuk)
		fmt.Println()
		fmt.Println("Apakah anda ingin membalas pesan tersebut?")
		fmt.Println("A. Iya")
		fmt.Println("B. Tidak")
		fmt.Scan(&jawabBacaInbox)
		if jawabBacaInbox == "A" {
			*usernameDiCari = A[*posisiLogin].sender
			tempCariUname = searchUsername(A, &n, i, *usernameDiCari)
			A[tempCariUname].sender = A[*posisiLogin].username
			fmt.Println("Silahkan Masukkan pesan yang ingin anda kirim: ")
			fmt.Scan(&A[tempCariUname].pesanMasuk)
			fmt.Println("Pesan berhasil dibalas ke:", A[tempCariUname].username)
			setelahLogin(A, n, posisiLogin)
		} else if jawabBacaInbox == "B" {
			setelahLogin(A, n, posisiLogin)
		}
	}
}

func abisLogOut(A *tabEmail, n *int) {
	var jawabAbisLogOut, usernameDiCari, passwordDiCari string
	var posisiLogin int
	var B tabAgree
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("------------------------   |Aplikasi Email|   ------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("A. Register akun")
	fmt.Println("B. Log in akun")
	fmt.Scan(&jawabAbisLogOut)
	if jawabAbisLogOut == "A" {
		registerEmail(A, &B, *n, &usernameDiCari)
	} else if jawabAbisLogOut == "B" {
		loginEmail(A, *n, &posisiLogin, &usernameDiCari, &passwordDiCari)
	}
}

func menuAwal(A *tabEmail, n int) {
	var B tabAgree
	var jawabanMenuAwal, usernameDiCari, passwordDiCari string
	var posisiLogin int
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("------------------------    Aplikasi Email    ------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("A. Register akun")
	fmt.Println("B. Log in akun")
	fmt.Scan(&jawabanMenuAwal)
	if jawabanMenuAwal == "A" {
		registerEmail(A, &B, n, &usernameDiCari)
	} else if jawabanMenuAwal == "B" {
		loginEmail(A, n, &posisiLogin, &usernameDiCari, &passwordDiCari)
	}
}

func main() {
	var dataEmail tabEmail
	var nEmail int
	menuAwal(&dataEmail, nEmail)
}
