package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/starkandwayne/install-debs-in-order/debpkg"
)

var Version = ""

func checkDebian() {
	_, err := exec.LookPath("dpkg-deb")
	if err != nil {
		log.Fatal("Please run inside Debian/Ubuntu")
	}
}

func showVersion() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			if Version == "" {
				fmt.Printf("install-debs-in-order (development)\n")
			} else {
				fmt.Printf("install-debs-in-order v%s\n", Version)
			}
			os.Exit(0)
		}
	}
}

func main() {
	showVersion()
	checkDebian()

	targetPath := os.Args[1]

	folder, err := debpkg.NewDebianPackagesFromFolder(targetPath)
	if err != nil {
		log.Fatal(err)
	}
	folder.RemovePreinstalledPackages()
	installList := folder.OrderedInstallationList()
	for _, pkg := range installList {
		fmt.Printf("dpkg -i %s\n", filepath.Join(targetPath, pkg.FileName))
	}
}
