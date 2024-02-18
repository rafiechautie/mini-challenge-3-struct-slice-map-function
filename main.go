package main

import (
	"fmt"
	"os"
)

func main (){
	if len(os.Args) < 2 {
		fmt.Println("Harap masukkan nama peserta")
		return
	}

	namaPeserta := os.Args[1]
	showDataParticipants(namaPeserta)
}