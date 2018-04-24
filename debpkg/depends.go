package debpkg

import (
	"context"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// DebianPackage describes some metadata of a .deb package via dpkg-deb
type DebianPackage struct {
	PackageName string
	FileName    string
	FilePath    string
	RawVersion  string
	RawDepends  string
	Depends     []DebianPackageDependency
}

// DebianPackageDependency describes a named dependency of a DebianPackage
type DebianPackageDependency struct {
	PackageName string
}

// NewDebianPackageFromFile extracts metadata from a .deb package via dpkg-deb
func NewDebianPackageFromFile(filePath string) (pkg *DebianPackage, err error) {
	ctx := context.Background()

	pkg = &DebianPackage{
		FilePath: filePath,
		FileName: filepath.Base(filePath),
	}
	out, err := exec.CommandContext(ctx, "dpkg-deb", "-f", filePath, "Depends").Output()
	if err != nil {
		return nil, err
	}
	pkg.RawDepends = strings.TrimSpace(string(out))

	out, err = exec.CommandContext(ctx, "dpkg-deb", "-f", filePath, "Package").Output()
	if err != nil {
		return nil, err
	}
	pkg.PackageName = strings.TrimSpace(string(out))

	out, err = exec.CommandContext(ctx, "dpkg-deb", "-f", filePath, "Version").Output()
	if err != nil {
		return nil, err
	}
	pkg.RawVersion = strings.TrimSpace(string(out))

	pkg.processDepends()
	return
}

func (pkg *DebianPackage) processDepends() {
	re, _ := regexp.Compile(`^([^\s]+)(?:\s*\((.*)\))?`)
	rawDependencies := strings.Split(pkg.RawDepends, ", ")
	pkg.Depends = make([]DebianPackageDependency, len(rawDependencies))
	for i, dependencyStr := range rawDependencies {
		matches := re.FindAllStringSubmatch(dependencyStr, -1)
		pkg.Depends[i].PackageName = matches[0][1]
	}
}
