# Install Debian packages in order

Let's say you've downloaded many Debian packages into a folder, and then later you want to install them. You'll need to do this in the reverse order of their dependency tree. Packages without dependencies get installed first, etc.

This CLI creates a shell script to be used later to correctly install the folder of `.deb` files in a correct order.

It is assumed that this CLI is being run in a Debian/Ubuntu environment with access to the `dpkg-deb` command to inspect `.deb` files for metadata.

Either run the CLI and cache the output to a script to be run later:

```plain
install-debs-in-order path/to/many/debs > install-debs.sh
```

Later, run `./install-debs.sh` to install the packages.


## Run tests inside Ubuntu Docker container

```plain
bin/test
```

## Run inside Ubuntu Docker container

```plain
docker run -ti \
    -v $PWD:/go/src/github.com/starkandwayne/install-debs-in-order \
    -v $PWD:/app golang:1 \
    /app/bin/install-and-run /app/fixtures/debs/archives
```

Inside the container:

```plain
go test github.com/starkandwayne/install-debs-in-order
go install github.com/starkandwayne/install-debs-in-order
install-debs-in-order
```

## Install from source

```plain
go get github.com/starkandwayne/install-debs-in-order
```

## Fetch a fixture package

```plain
docker run -ti -v $PWD/fixtures/debs:/debs golang:1
```

Inside the container:

```plain
apt-get update
apt-get install -o debug::nolocking=true -o dir::cache=/debs \
    tree
```

The `tree_1.7.0-5_amd64.deb` file will be placed in `/debs/archives/tree_1.7.0-5_amd64.deb`, which maps to `fixtures/debs/archives/tree_1.7.0-5_amd64.deb` in this project.

To see the list of dependencies:

```plain
# dpkg-deb -f archives/tree*.deb Depends
libc6 (>= 2.14)
```
