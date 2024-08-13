package savefile

import (
	"fmt"
	"os"
)

func Run() {
	data := []byte("Dipra")

	err := os.WriteFile("name.txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Name was saved")
}
