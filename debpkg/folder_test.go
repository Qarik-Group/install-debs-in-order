package debpkg

import (
	"testing"
)

func TestNewDebianPackagesFromFolder(t *testing.T) {
	folder, err := NewDebianPackagesFromFolder("/app/fixtures/debs/archives/")
	if err != nil {
		t.Error("Should not have error, got: ", err)
	}
	if folder.FolderPath != "/app/fixtures/debs/archives" {
		t.Error("Expected FolderPath to be trimmed, got ", folder.FolderPath)
	}
	if folder.PackageNameToFileNames["tree"] != "tree_1.7.0-5_amd64.deb" {
		t.Error("Expected tree to map to tree_1.7.0-5_amd64.deb, got ", folder.PackageNameToFileNames["tree"])
	}
	if folder.FileNamesToPackageName["tree_1.7.0-5_amd64.deb"] != "tree" {
		t.Error("Expected tree_1.7.0-5_amd64.deb to map to tree, got ", folder.FileNamesToPackageName["tree_1.7.0-5_amd64.deb"])
	}
}
