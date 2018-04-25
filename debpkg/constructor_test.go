package debpkg

import "testing"

func TestNewDebianPackageFromFile_tree(t *testing.T) {
	deb, err := NewDebianPackageFromFile("/app/fixtures/debs/archives/tree_1.7.0-5_amd64.deb")
	if err != nil {
		t.Error("Should not have error, got: ", err)
	}
	if deb.PackageName != "tree" {
		t.Error("Expected PackageName to be tree, got ", deb.PackageName)
	}
	if deb.RawVersion != "1.7.0-5" {
		t.Error("Expected Version to be 1.7.0-5, got ", deb.RawVersion)
	}
	if len(deb.Depends) != 1 {
		t.Error("Expected 1 dependency, got ", deb.Depends, "from", deb.RawDepends)
	}
	if deb.Depends[0].PackageName != "libc6" {
		t.Error("Expected dependency 'libc6', got ", deb.Depends[0].PackageName)
	}
}

func TestNewDebianPackageFromFile_dbus(t *testing.T) {
	deb, err := NewDebianPackageFromFile("/app/fixtures/debs/archives/dbus_1.10.26-0+deb9u1_amd64.deb")
	if err != nil {
		t.Error("Should not have error, got: ", err)
	}
	if deb.PackageName != "dbus" {
		t.Error("Expected PackageName to be dbus, got ", deb.PackageName)
	}
	if deb.RawVersion != "1.10.26-0+deb9u1" {
		t.Error("Expected Version to be 1.10.26-0+deb9u1, got ", deb.RawVersion)
	}
	if len(deb.Depends) != 11 {
		t.Error("Expected 11 dependencies, got ", deb.Depends, "from", deb.RawDepends)
	}
	if deb.Depends[0].PackageName != "adduser" {
		t.Error("Expected first dependency 'adduser', got ", deb.Depends[0].PackageName)
	}
}
