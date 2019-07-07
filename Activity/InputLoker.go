package Activity

import (
	"fmt"
	"os"
	"strconv"
)

func InputLoker() {
	var input string

	fmt.Println("")
	fmt.Printf("Input jumlah loker : ")
	fmt.Scan(&input)

	if _, err := strconv.Atoi(input); err == nil {
		input, err := strconv.Atoi(input)
		if err == nil {
			if (0 < input) && (input <= 5)  {
				fmt.Println(">> Loker berhasil di buat")

				loker := make([]string, input)
				CekStatus(loker)
			}else {
				InputLokerError()
			}
		}
	}else {
		InputLokerError()
	}
}

func CekStatus(loker []string)  {
	var jmlLoker = 0
	var jmlLokerTersedia = 0
	var prompt string

	fmt.Println("")
	fmt.Println("=====================")
	for i := 0; i <= len(loker)-1; i++ {
		if loker[i] == "" {
			jmlLoker += 1
			jmlLokerTersedia +=1
			fmt.Println("Loker", jmlLoker, "||", "Kosong")
		}else {
			jmlLoker += 1
			fmt.Println("Loker", jmlLoker, "||", "Ada isinya")
		}
	}
	fmt.Println("=====================")

	fmt.Println(">> Jumlah loker tersedia : ", jmlLokerTersedia)

	if  jmlLokerTersedia < len(loker){
		fmt.Println("")
		fmt.Printf("Cek Loker? [Y/N] : ")
		fmt.Scan(&prompt)

		if prompt == "Y" || prompt == "y"{
			CariLoker(loker)
		}else if prompt == "N" || prompt == "n"{

		}else{

		}
	}else{

	}

	fmt.Println("")
	fmt.Printf("Simpan ID ke Loker? [Y/N] : ")
	fmt.Scan(&prompt)

	if prompt == "Y" || prompt == "y"{
		InputID(loker, jmlLokerTersedia)
	}else if prompt == "N" || prompt == "n"{
		os.Exit(1)
	}else{
		os.Exit(1)
	}
}

func InputID(loker []string, jmlLokerTersedia int){

	var IDtipe string
	var IDno string
	var lokerID int
	var prompt string

	if jmlLokerTersedia == 0 {
		fmt.Println(">> Loker sudah penuh" )
		fmt.Println("" )
		fmt.Printf("Kosongkan loker? [Y/N] : ")
		fmt.Scan(&prompt)

		if prompt == "Y" || prompt == "y"{
			ClearLoker(loker)
		}else if prompt == "N" || prompt == "n"{
			os.Exit(1)
		}else{
			os.Exit(1)
		}
	}else{
		fmt.Println("")
		fmt.Printf("Masukan tipe ID : " )
		fmt.Scan(&IDtipe)
		fmt.Printf("Masukan no ID : " )
		fmt.Scan(&IDno)

		for i := 0; i <= len(loker)-1; i++ {
			if loker[i] == "" {
				lokerID = i
				break
			}
		}

		loker[lokerID] = `{"tipeID": "` + IDtipe + `", "noID": "`+ IDno +`"}`
		CekStatus(loker)
	}
}

func ClearLoker(loker []string)  {
	var prompt string

	fmt.Printf("Kosongkan loker nomor : " )
	fmt.Scan(&prompt)

	if _, err := strconv.Atoi(prompt); err == nil {
		prompt, err := strconv.Atoi(prompt)

		if err == nil {
			if (prompt > 0) && (prompt <= len(loker)){
				loker[prompt-1] = ""

				prompt := strconv.Itoa(prompt)
				fmt.Println(">> Loker Nomor " + prompt + " telah di kosongkan" )
				CekStatus(loker)
			}else {
				fmt.Println(">> Nomor Loker tidak tersedia" )
				CekStatus(loker)
			}
		}
	} else {
		fmt.Println(">> Nomor Loker tidak tersedia" )
		CekStatus(loker)
	}
}