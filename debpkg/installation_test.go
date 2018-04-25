package debpkg

import (
	"path/filepath"
	"testing"
)

func TestOrderedInstallationList(t *testing.T) {
	archives, _ := filepath.Abs("../fixtures/debs/archives/")
	folder, err := NewDebianPackagesFromFolder(archives)
	if err != nil {
		t.Error("Should not have error, got: ", err)
	}

	folder.RemovePreinstalledPackages()

	list := folder.OrderedInstallationList()
	if len(list) != len(folder.Packages) {
		t.Errorf("Try harder. There are %d packages, but the ordered list has %d", len(folder.Packages), len(list))
	}
	if list[len(list)-1].PackageName != "dbus" {
		t.Errorf("Currently, dbus must be the last package installed due to dependencies")
	}
}
