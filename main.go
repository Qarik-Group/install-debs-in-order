package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func checkDebian() {
	_, err := exec.LookPath("dpkg-deb")
	if err != nil {
		log.Fatal("Please run inside Debian/Ubuntu")
	}
}

func main() {
	checkDebian()
	ctx := context.Background()

	targetPath := os.Args[1]

	fmt.Println("Looking for files", targetPath)
	targetPathFiles, err := ioutil.ReadDir(targetPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range targetPathFiles {
		if filepath.Ext(file.Name()) == ".deb" {
			fmt.Println(file.Name())

			fullpath := filepath.Join(targetPath, file.Name())
			out, err := exec.CommandContext(ctx, "dpkg-deb", "-f", fullpath, "Depends").Output()
			if err != nil {
				fmt.Printf("Run failed: %s\n", err)
			} else {
				fmt.Printf("Output: %s\n", string(out))
			}
		}
	}

}
