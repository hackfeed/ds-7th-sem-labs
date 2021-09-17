package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hackfeed/ds-7th-sem-labs/lab_01/pkg/sys"
	"github.com/logrusorgru/aurora/v3"
)

func main() {
	fmt.Println(aurora.BgBlue("Performing installation of cat utility..."))

	s := spinner.New(spinner.CharSets[32], 100*time.Millisecond)
	s.Color("bgYellow", "bold", "fgBlack")
	s.Start()
	time.Sleep(5 * time.Second)
	s.Stop()

	key, err := sys.GetKey()
	if err != nil {
		fmt.Println(aurora.BgRed("Failed to obtain host key. Aborting"))
		os.Exit(1)
	}

	cmd := exec.Command("go", "build", "-ldflags", fmt.Sprintf("-X main.LicenseKey=%s", key), "-o", "build/cat_arm64.exe", "cmd/cat/main.go")
	cmd.Env = append(os.Environ(), "GOARCH=arm64")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%s %s\n", aurora.BgRed("Error while building ARM executable:"), err)
		os.Exit(1)
	}

	cmd = exec.Command("go", "build", "-ldflags", fmt.Sprintf("-X main.LicenseKey=%s", key), "-o", "build/cat_amd64.exe", "cmd/cat/main.go")
	cmd.Env = append(os.Environ(), "GOARCH=amd64")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%s %s\n", aurora.BgRed("Error while building x86 executable:"), err)
		os.Exit(1)
	}

	fmt.Printf("\n%s%s\n", aurora.BgGreen("Performed installation for host key: "), aurora.BgGreen(aurora.Blue(key)))
}
