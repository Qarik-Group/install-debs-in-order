package main

import (
	"fmt"
	"io/ioutil"
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

	fmt.Println("Looking for files", targetPath)
	targetPathFiles, err := ioutil.ReadDir(targetPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range targetPathFiles {
		if filepath.Ext(file.Name()) == ".deb" {
			fullpath := filepath.Join(targetPath, file.Name())
			debpkg, err := debpkg.NewDebianPackageFromFile(fullpath)
			if err != nil {
				fmt.Printf("Failed to lookup package %s: %s\n", file.Name(), err)
			} else {
				fmt.Printf("%#v\n", debpkg)
			}

		}
	}

}
