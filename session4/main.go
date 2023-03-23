package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman
type Friend struct {
	Name    string
	Address string
	Job     string
	Alasan  string
}

// Function untuk mengambil data teman berdasarkan nomor absen
func getFriendData(absen int) Friend {
	friendList := []Friend{
		{
			Name:    "Hera Herwati",
			Address: "Semarang",
			Job:     "Akuntan",
			Alasan:  "Mau mencoba ngoding dengan go lang huhuhu",
		},
		{
			Name:    "Ahmad Hanif",
			Address: "Bogor",
			Job:     "Fullstack Engineer",
			Alasan:  "Bosen dengan javascript dan php hehehe",
		},
		{
			Name:    "Hilal Faiq",
			Address: "Magelang",
			Job:     "Web Developer",
			Alasan:  "go lang is may way HAHAHA",
		},
	}

	// Validasi nomor absen
	if absen < 1 || absen > len(friendList) {
		fmt.Println("Id teman tidak valid!")
		os.Exit(1)
	}

	return friendList[absen-1]
}

func main() {
	// Mengambil argument nomor absen dari command line
	absen := 0
	if len(os.Args) > 1 {
		fmt.Sscanf(os.Args[1], "%d", &absen)
	}

	// Menampilkan data teman berdasarkan nomor absen
	friendData := getFriendData(absen)
	fmt.Println("Data teman dengan nomor aidi", absen)
	fmt.Println("Nama:", friendData.Name)
	fmt.Println("Alamat:", friendData.Address)
	fmt.Println("Pekerjaan:", friendData.Job)
	fmt.Println("Moto:", friendData.Alasan)
}
