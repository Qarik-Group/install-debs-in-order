package debpkg

import (
	"io/ioutil"
	"path/filepath"
)

// DebianPackagesFolder describes a list of *.deb packages in a folder
type DebianPackagesFolder struct {
	FolderPath             string
	Packages               []*DebianPackage
	FileNamesToPackages    map[string]*DebianPackage
	PackageNameToFileNames map[string]string
}

// NewDebianPackagesFromFolder discovers *.deb files in a folder
func NewDebianPackagesFromFolder(folderpath string) (folder *DebianPackagesFolder, err error) {
	folder = &DebianPackagesFolder{
		FolderPath:             filepath.Clean(folderpath),
		Packages:               []*DebianPackage{},
		FileNamesToPackages:    map[string]*DebianPackage{},
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
			folder.FileNamesToPackages[debpkg.FileName] = debpkg
			folder.PackageNameToFileNames[debpkg.PackageName] = debpkg.FileName
		}
	}
	return
}

// RemovePreinstalledPackages cleans up each folder.Packages[].Depends of any packages
// that are not in PackageNameToFileNames map.
// We do this because we only want to track each Package relative to other Packages
// we need to explicitly install. We can ignore dependencies we think are already installed.
func (folder *DebianPackagesFolder) RemovePreinstalledPackages() {
	for _, deb := range folder.Packages {
		for _, dependency := range deb.Depends {
			if _, ok := folder.PackageNameToFileNames[dependency.PackageName]; ok {
				pkgDependency := DebianPackageDependency{PackageName: dependency.PackageName}
				deb.UninstalledDependencies = append(deb.UninstalledDependencies, pkgDependency)
			}
		}
	}
}
