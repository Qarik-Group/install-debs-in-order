package debpkg

// OrderedInstallationList returns a list of DebianPackages that can be installed before
// those packages that depend upon them.
func (folder *DebianPackagesFolder) OrderedInstallationList() (packages []*DebianPackage) {
	return folder.addPackagesWithNoUninstalledDependencies([]*DebianPackage{}, folder.Packages)
}

func (folder *DebianPackagesFolder) addPackagesWithNoUninstalledDependencies(
	orderedPackages, remainingPackages []*DebianPackage) []*DebianPackage {

	if len(remainingPackages) == 0 {
		return orderedPackages
	}
	newOrderedPackages := orderedPackages
	newRemainingPackages := []*DebianPackage{}
	newlyInstalledPackages := []*DebianPackage{}

	// find uninstalled packages with no uninstalled dependencies,
	// and install them now
	for _, pkg := range remainingPackages {
		if len(pkg.UninstalledDependencies) == 0 {
			newlyInstalledPackages = append(newlyInstalledPackages, pkg)
			newOrderedPackages = append(newOrderedPackages, pkg)
		} else {
			newRemainingPackages = append(newRemainingPackages, pkg)
		}
	}

	// remove the newly installed packages from the Depends list
	// of packages that are not yet installed
	for _, installedPkg := range newlyInstalledPackages {
		for _, remainingPkg := range newRemainingPackages {
			newDependsList := []DebianPackageDependency{}
			for _, dependency := range remainingPkg.UninstalledDependencies {
				if dependency.PackageName != installedPkg.PackageName {
					newDependsList = append(newDependsList, dependency)
				}
			}
			remainingPkg.UninstalledDependencies = newDependsList
		}
	}
	if len(newRemainingPackages) == len(remainingPackages) {
		for _, pkg := range newRemainingPackages {
			pkg.IgnoreDependencies = pkg.UninstalledDependencies
			newOrderedPackages = append(newOrderedPackages, pkg)
		}
		return newOrderedPackages
	}

	return folder.addPackagesWithNoUninstalledDependencies(newOrderedPackages, newRemainingPackages)
}
