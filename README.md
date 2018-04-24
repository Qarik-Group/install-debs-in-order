# Install Debian packages in order

Let's say you've downloaded many Debian packages into a folder, and then later you want to install them. You'll need to do this in the reverse order of their dependency tree. Packages without dependencies get installed first, etc.

This CLI creates a shell script to be used later to correctly install the folder of `.deb` files in a correct order.

It is assumed that this CLI is being run in a Debian/Ubuntu environment with access to the `dpkg-deb` command to inspect `.deb` files for metadata.

Either run the CLI and cache the output to a script to be run later:

```plain
install-debs-in-order path/to/many/debs > install-debs.sh
```

Later, run `./install-debs.sh` to install the packages.

## Install from source

```plain
go get github.com/starkandwayne/install-debs-in-order
```