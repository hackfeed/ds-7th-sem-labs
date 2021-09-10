package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hackfeed/ds-7th-sem-labs/lab_01/internal/sys"
	"github.com/logrusorgru/aurora/v3"
)

var LicenseKey string

func main() {
	isLicensed, err := sys.CheckKey(LicenseKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if !isLicensed {
		fmt.Println(aurora.BgRed("Program is not registered for this PC. Aborting"))
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println(aurora.Green("USAGE: ./cat <file_to_read>"))
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
