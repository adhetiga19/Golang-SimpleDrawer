package Activity

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	TipeID string
	NoID string
}

func CariLoker (loker []string){
	var prompt int

	fmt.Printf("Cari loker berdasarkan? [1] Tipe ID, [2] No ID : ")
	fmt.Scan(&prompt)

	if prompt == 1 {
		CariTipeID(loker)
	}else if prompt == 2 {
		CariNoID(loker)
	}else{
		os.Exit(1)
	}
}

func CariTipeID (loker []string){
	var prompt string
	var dataID User
	var data int = 1

	fmt.Printf("Cari loker berdasarkan Tipe ID : ")
	fmt.Scan(&prompt)

	for i := 0; i <= len(loker)-1; i++ {
		var jsonData = []byte(loker[i])

		json.Unmarshal(jsonData, &dataID)

		if prompt == dataID.TipeID {
			fmt.Println(">>", prompt, "anda ada di dalam loker nomor", i+1)
			break
		}else {
			data = 0
		}
	}

	if data == 0{
		fmt.Println(">> Data tidak di temukan")
	}

	CekStatus(loker)
}

func CariNoID (loker []string){
	var prompt string
	var dataID User
	var data int = 1

	fmt.Printf("Cari loker berdasarkan No ID : ")
	fmt.Scan(&prompt)

	for i := 0; i <= len(loker)-1; i++ {
		var jsonData = []byte(loker[i])

		json.Unmarshal(jsonData, &dataID)

		if prompt == dataID.NoID {
			fmt.Println(">>", dataID.TipeID, "anda ada di dalam loker nomor", i+1)
			break
		}else {
			data = 0
		}
	}

	if data == 0{
		fmt.Println(">> Data tidak di temukan")
	}

	CekStatus(loker)
}