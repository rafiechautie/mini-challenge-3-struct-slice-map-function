package main

import (
	"fmt"
	"os"
	"strconv"
)

type Users struct {
	id int
	name    string
	address string
	job     string
	reason  string
}

var participantBootcampGo = []Users{
	{1, "Fitri", "Jakarta", "Software Engineer", "Ingin mempelajari Golang"},
	{2, "Budi", "Bandung", "Data Analyst", "Ingin mengembangkan skill backend"},
	{3, "Vanessa", "Medan", "Data Analyst", "Ingin berkarir di backend developer"},
	{4, "Syauqi", "Jambi", "Data Analyst", "Penasaran dengan bahasa golang yang katanya cepat"},
	{5, "Rafie", "Depok", "Fullstack Developer", "Ingin menambah pengetahuan karna golang banyak digunakan oleh perusahaan di Indonesia"},
}


func ShowDataParticipants(input string) {
	id, err := strconv.Atoi(input) 

	if err == nil {
		for _, participant := range participantBootcampGo {
			if participant.id == id {
				fmt.Printf("ID: %d\n", participant.id)
				fmt.Printf("Nama: %s\n", participant.name)
				fmt.Printf("Alamat: %s\n", participant.address)
				fmt.Printf("Pekerjaan: %s\n", participant.job)
				fmt.Printf("Alasan memilih kelas Golang: %s\n", participant.reason)
				return
			}
		}
	} else { 
		for _, participant := range participantBootcampGo {
			if participant.name == input {
				fmt.Printf("ID: %d\n", participant.id)
				fmt.Printf("Nama: %s\n", participant.name)
				fmt.Printf("Alamat: %s\n", participant.address)
				fmt.Printf("Pekerjaan: %s\n", participant.job)
				fmt.Printf("Alasan memilih kelas Golang: %s\n", participant.reason)
				return
			}
		}
	}

	fmt.Println("Data dengan ID/nama tersebut tidak tersedia")
}

func main (){
	if len(os.Args) < 2 {
		fmt.Println("Harap masukkan nama peserta")
		return
	}

	input := os.Args[1]
	ShowDataParticipants(input)
}