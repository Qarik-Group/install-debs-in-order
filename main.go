package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/starkandwayne/install-debs-in-order/debpkg"
)

func checkDebian() {
	_, err := exec.LookPath("dpkg-deb")
	if err != nil {
		log.Fatal("Please run inside Debian/Ubuntu")
	}
}

func main() {
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
