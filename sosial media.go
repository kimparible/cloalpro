package main

import (
	"bufio"
	"fmt"
	"os"
)

type profile struct {
	name       string
	username   string
	password   string
	status     string
	following  [1000]string
	nFollowing int
	followers  [1000]string
	nFollowers int
}

type Profiles [1000]profile

type post struct {
	usernamePengirim string
	posting          string
	nKomentar        int
	komentar         Comment
}

type Post [1000]post

type comment struct {
	usernameKomen string
	commenting    string
}

type Comment [1000]comment

func teks() string {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()
}

func register(arrProfile *Profiles, nProfile *int) {
	var nama, username, password string

	fmt.Println("---Melakukan registrasi akun---")
	fmt.Printf("Nama     : ")
	nama = teks()
	fmt.Printf("Username : ")
	username = teks()
	fmt.Printf("Password : ")
	password = teks()

	arrProfile[*nProfile].name = nama
	arrProfile[*nProfile].username = username
	arrProfile[*nProfile].password = password
	*nProfile++
}

func login(arrProfile Profiles, nProfile int) int {
	var username, password string

	var found bool = false
	var idx, i int

	fmt.Println("---Melakukan login---")
	fmt.Printf("Username : ")
	username = teks()
	fmt.Printf("Password : ")
	password = teks()

	for i = 0; i < nProfile; i++ {
		if arrProfile[i].username == username {
			found = true
			idx = i
		}
	}

	if !found {
		fmt.Println("Akun tidak ditemukan!")
		return -1
	} else {
		if arrProfile[idx].password == password {
			fmt.Println("Berhasil melakukan login!")
			return idx
		} else {
			fmt.Println("Password salah!")
			return -1
		}
	}
}

func editProfile(arrProfile *Profiles, idx int, arrPost *Post, nPost int) {

	var input string
	var i int

	fmt.Println("---Melakukan edit profile---")

	fmt.Printf("Ingin ubah nama? (Y/N) : ")
	fmt.Scanln(&input)
	if input == "Y" {
		fmt.Printf("Masukkan nama baru : ")
		input = teks()

		arrProfile[idx].name = input
	}

	fmt.Printf("Ingin ubah username? (Y/N) : ")
	fmt.Scanln(&input)
	if input == "Y" {
		fmt.Printf("Masukkan username baru : ")
		input = teks()

		for i = 0; i < nPost; i++ {
			if arrPost[i].usernamePengirim == arrProfile[idx].name {
				arrPost[i].usernamePengirim = input
			}
		}

		arrProfile[idx].username = input
	}

	fmt.Printf("Ingin ubah password? (Y/N) : ")
	fmt.Scanln(&input)
	if input == "Y" {
		fmt.Printf("Masukkan password baru : ")
		input = teks()

		arrProfile[idx].password = input
	}

	fmt.Printf("Ingin ubah status? (Y/N) : ")
	fmt.Scanln(&input)
	if input == "Y" {
		fmt.Printf("Masukkan status baru : ")
		input = teks()

		arrProfile[idx].status = input
	}

	fmt.Println("Profile berhasil diedit")

}

func tampilPost(arrPost Post, nPost int) {
	var i int

	fmt.Println("---Menampilkan post---")
	for i = 0; i < nPost; i++ {
		fmt.Printf("%v. User : ", i+1)
		fmt.Println(arrPost[i].usernamePengirim)
		fmt.Printf("\t\t")
		fmt.Println(arrPost[i].posting)
		fmt.Println()
	}
}

func tambahPost(arrPost *Post, nPost *int, arrProfile Profiles, idxProfile int) {
	var masukan string

	fmt.Println("---Menambahkan post---")
	fmt.Printf("POST : ")
	masukan = teks()

	arrPost[*nPost].usernamePengirim = arrProfile[idxProfile].username
	arrPost[*nPost].posting = masukan
	*nPost++
}

func komentar(arrPost *Post, arrProfile Profiles, idxPost, idxProfile int) {
	var komentar string
	var idxKomen int
	var i int

	idxKomen = arrPost[idxPost].nKomentar
	for i = 0; i < idxKomen; i++ {
		fmt.Printf("%v. %s\n", i+1, arrPost[idxPost].komentar[i].usernameKomen)
		fmt.Printf("\t\t")
		fmt.Println(arrPost[idxPost].komentar[i].commenting)
	}

	fmt.Println("---Menambahkan komentar---")
	fmt.Printf("Masukkan komentar : ")
	komentar = teks()

	idxKomen = arrPost[idxPost].nKomentar
	arrPost[idxPost].komentar[idxKomen].usernameKomen = arrProfile[idxProfile].username
	arrPost[idxPost].komentar[idxKomen].commenting = komentar
	arrPost[idxPost].nKomentar++
}

func cariProfile(arrProfile *Profiles, nProfile int, idxProfile int) {
	var uName, input string
	var idx, i int

	idx = -1

	fmt.Println("---Cari profile---")
	fmt.Printf("Masukkan username : ")
	uName = teks()

	for i = 0; i < nProfile; i++ {
		if arrProfile[i].username == uName {
			idx = i
		}
	}

	if idx == -1 {
		fmt.Println("Username tidak ditemukan!")
	} else {
		fmt.Println("---Menampilkan profile---")
		fmt.Println("Username : ", arrProfile[idx].username)
		fmt.Println("Nama     : ", arrProfile[idx].name)
		fmt.Println("Status   : ", arrProfile[idx].status)

		fmt.Printf("Ingin follow? (Y/N) : ")
		fmt.Scanln(&input)

		if input == "Y" {
			tambahFollow(&*arrProfile, idxProfile, idx)
			fmt.Println("---Berhasil melakukan follow---")
		}
	}
}

func tambahFollow(arrProfile *Profiles, idxProfile int, idxTeman int) {
	var i int

	i = arrProfile[idxTeman].nFollowers
	arrProfile[idxTeman].followers[i] = arrProfile[idxProfile].username
	arrProfile[idxTeman].nFollowers++

	i = arrProfile[idxProfile].nFollowing
	arrProfile[idxProfile].following[i] = arrProfile[idxTeman].username
	arrProfile[idxProfile].nFollowing++
}

func cariProfileTeman(arrProfile Profiles, nProfile int, idxProfile int) {
	var uName, input string
	var idx, i int
	var teman bool = false

	fmt.Println("---Cari profile---")
	fmt.Printf("Masukkan username : ")
	uName = teks()

	for i = 0; i < arrProfile[idxProfile].nFollowing; i++ {
		if arrProfile[idxProfile].following[i] == uName {
			teman = true
		}
	}

	if teman {
		for i = 0; i < nProfile; i++ {
			if arrProfile[i].username == uName {
				idx = i
			}
		}

		fmt.Println("---Menampilkan profile---")
		fmt.Println("Username : ", arrProfile[idx].username)
		fmt.Println("Nama     : ", arrProfile[idx].name)
		fmt.Println("Status   : ", arrProfile[idx].status)

		fmt.Printf("Ingin unfollow? (Y/N) : ")
		fmt.Scanln(&input)

		if input == "Y" {
			hapusFollow(&arrProfile, idxProfile, idx)
			fmt.Println("---Berhasil melakukan unfollow---")
		}
	} else {
		fmt.Println("Tidak ada teman dengan username", uName)
	}

}

func hapusFollow(arrProfile *Profiles, idxProfile int, idxTeman int) {
	var i, j, dummy int //var dummy hanya untuk menyimpan data sementara

	dummy = arrProfile[idxProfile].nFollowing
	for i = 0; i < dummy; i++ {
		if arrProfile[idxProfile].following[i] == arrProfile[idxTeman].username {
			for j = i; j < dummy; j++ {
				arrProfile[idxProfile].following[j] = arrProfile[idxProfile].following[j+1]
			}
			arrProfile[idxProfile].nFollowing--
		}
	}

	dummy = arrProfile[idxTeman].nFollowers
	for i = 0; i < dummy; i++ {
		if arrProfile[idxTeman].followers[i] == arrProfile[idxProfile].username {
			for j = i; j < dummy; j++ {
				arrProfile[idxTeman].followers[j] = arrProfile[idxTeman].followers[j+1]
			}
			arrProfile[idxTeman].nFollowers--
		}
	}
}

func tampilDataTeman(arrProfile Profiles, idxProfile int) {
	var n int = arrProfile[idxProfile].nFollowing
	var temp [1000]string = arrProfile[idxProfile].following
	var i, j int

	if n == 0 {
		fmt.Println("---Tidak ada teman---")
	}
	for i = 1; i < n; i++ {
		j = i
		for j > 0 {
			if len(temp[j-1]) > len(temp[j]) {
				temp[j-1], temp[j] = temp[j], temp[j-1]
			}
			j = j - 1
		}
	}

	for i = 0; i < n; i++ {
		fmt.Printf("%v. %s\n", i+1, temp[i])
	}
}

func main() {
	var arrProfile Profiles
	var arrPost Post
	var nProfile, nPost int
	var idx, idxTemp, input int

	idx = -1
	for true {
		fmt.Println("______________________________________")
		fmt.Println("-------------Sosial Media-------------")
		fmt.Println("______________________________________")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Printf("Input : ")
		fmt.Scanln(&input)

		if input == 1 {
			register(&arrProfile, &nProfile)
		} else if input == 2 {
			idx = login(arrProfile, nProfile)

			for idx > -1 {
				fmt.Println("--------------------------------------")
				fmt.Println("1. Homepage")
				fmt.Println("2. Search")
				fmt.Println("3. Profile")
				fmt.Println("0. Log out")
				fmt.Printf("Input : ")
				fmt.Scanln(&input)

				if input == 1 {
					if nPost > 0 {
						tampilPost(arrPost, nPost)
					} else {
						fmt.Println("---Tidak ada post---")
					}

					fmt.Println("--------------------")
					fmt.Println("1. Tambah post")
					if nPost > 0 {
						fmt.Println("2. Beri komentar")
					}
					fmt.Println("0. Kembali")

					fmt.Printf("Input : ")
					fmt.Scanln(&input)

					if input == 1 {
						tambahPost(&arrPost, &nPost, arrProfile, idx)
					} else {
						if nPost > 0 {
							if input == 2 {
								fmt.Println("---Pilih post---")
								fmt.Printf("Input : ")
								fmt.Scanln(&idxTemp)

								komentar(&arrPost, arrProfile, idxTemp-1, idx)
							}
						}
					}
				} else if input == 2 {
					fmt.Println("1. Cari pengguna")
					fmt.Println("2. Cari teman")
					fmt.Println("0. Kembali")

					fmt.Printf("Input : ")
					fmt.Scanln(&input)

					if input == 1 {
						cariProfile(&arrProfile, nProfile, idx)
					} else if input == 2 {
						cariProfileTeman(arrProfile, nProfile, idx)
					}

				} else if input == 3 {

					fmt.Println("---Menampilkan profile---")
					fmt.Printf("Username : ")
					fmt.Println(arrProfile[idx].username)

					fmt.Printf("Name : ")
					fmt.Println(arrProfile[idx].name)

					fmt.Printf("Status : ")
					fmt.Println(arrProfile[idx].status)

					fmt.Printf("Following : ")
					fmt.Println(arrProfile[idx].nFollowing)

					fmt.Printf("Followers : ")
					fmt.Println(arrProfile[idx].nFollowers)

					fmt.Println("1. Edit profile")
					fmt.Println("2. Tampil teman")
					fmt.Println("0. Kembali")

					fmt.Printf("Input : ")
					fmt.Scanln(&input)

					if input == 1 {
						editProfile(&arrProfile, idx, &arrPost, nPost)
					} else if input == 2 {
						tampilDataTeman(arrProfile, idx)
					}

				} else if input == 0 {
					idx = -1
				}
			}
		}

	}
}
