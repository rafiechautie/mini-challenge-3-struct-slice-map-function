package main

import "fmt"

func showDataParticipants(name string) {
	for _, participant := range participantBootcampGo {
		if participant.name == name {
			fmt.Printf("Nama: %s\n", participant.name)
			fmt.Printf("Alamat: %s\n", participant.address)
			fmt.Printf("Pekerjaan: %s\n", participant.job)
			fmt.Printf("Alasan memilih kelas Golang: %s\n", participant.reason)
			return
		}
	}

	fmt.Println("Data dengan nama tersebut tidak tersedia")

}