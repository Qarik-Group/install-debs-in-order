package debpkg

import (
	"io/ioutil"
	"path/filepath"
)

// DebianPackagesFolder describes a list of *.deb packages in a folder
type DebianPackagesFolder struct {
	FolderPath             string
	Packages               []*DebianPackage
	FileNamesToPackageName map[string]string
	PackageNameToFileNames map[string]string
}

// NewDebianPackagesFromFolder discovers *.deb files in a folder
func NewDebianPackagesFromFolder(folderpath string) (folder *DebianPackagesFolder, err error) {
	folder = &DebianPackagesFolder{
		FolderPath:             filepath.Clean(folderpath),
		FileNamesToPackageName: map[string]string{},
		PackageNameToFileNames: map[string]string{},
	}
	folder.loadPackagesFromFiles()
	return
}

func (folder *DebianPackagesFolder) loadPackagesFromFiles() (err error) {
	targetPathFiles, err := ioutil.ReadDir(folder.FolderPath)
	if err != nil {
		return
	}
	for _, file := range targetPathFiles {
		if filepath.Ext(file.Name()) == ".deb" {
			fullpath := filepath.Join(folder.FolderPath, file.Name())
			debpkg, err := NewDebianPackageFromFile(fullpath)
			if err != nil {
				return err
			}
			folder.Packages = append(folder.Packages, debpkg)
			folder.FileNamesToPackageName[debpkg.FileName] = debpkg.PackageName
			folder.PackageNameToFileNames[debpkg.PackageName] = debpkg.FileName
		}
	}
	return
}
