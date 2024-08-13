package userinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bufioSys() {
	// using bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Your name is %v", name)
}

func scanSys() {
	// using scan
	fmt.Print("Enter your name: ")
	var reader string
	fmt.Scan(&reader)
	fmt.Printf("Your name is %v", reader)
}

func Run() {
	scanSys()
}
