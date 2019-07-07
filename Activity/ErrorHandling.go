package Activity

import (
	"fmt"
	"os"
)

func InputLokerError()  {
	var prompt string

	fmt.Println(">> Jumlah loker yang dapat di buat adalah 1-5")
	fmt.Println("")
	fmt.Printf("Apakah anda ingin mengulangi? [Y/N] : ")
	fmt.Scan(&prompt)

	if prompt == "Y" || prompt == "y"{
		InputLoker()
	}else if prompt == "N" || prompt == "n"{
		os.Exit(1)
	}else{
		os.Exit(1)
	}
}