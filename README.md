# Install Debian packages in order

Let's say you've downloaded many Debian packages into a folder, and then later you want to install them. You'll need to do this in the reverse order of their dependency tree. Packages without dependencies get installed first, etc.

This CLI creates a shell script to be used later to correctly install the folder of `.deb` files in a correct order.

It is assumed that this CLI is being run in a Debian/Ubuntu environment with access to the `dpkg-deb` command to inspect `.deb` files for metadata.

## Usage

You can fetch Debian packages and store the `*.deb` files in a specific folder:

```plain
apt-get install -y -d -o debug::nolocking=true \
    -o dir::cache=path/to/many/debs \
    package1 package2
```

The two packages `package1` and `package2` might have nested dependency Debian packages. Later, when you want to install the packages, you need to install them in reverse order so `package1` and `package2` are installed last.

You can either:

```plain
. <(install-debs-in-order path/to/many/debs)
```

Or, after caching the files with `apt-get install -o dir::cache`, you can create a script for the next person to use to install the packages:

```plain
cd path/to/many/debs
install-debs-in-order . > install-packages.sh
```

Later, someone can install the packages with that script:

```plain
cd path/to/whereever/the/debs/are/now
. install-packages.sh
```

## Installation

The `install-debs-in-order` CLI can be installed via Debian:

```plain
wget -q -O - https://raw.githubusercontent.com/starkandwayne/homebrew-cf/master/public.key | apt-key add -
echo "deb http://apt.starkandwayne.com stable main" | tee /etc/apt/sources.list.d/starkandwayne.list
apt-get update

apt-get install install-debs-in-order
```

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
docker run -ti -v $PWD/fixtures/debs:/debs golang:1 bash
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
