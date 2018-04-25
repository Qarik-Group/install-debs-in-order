Initial release.

Basic usage:

```plain
install-debs-in-order path/to/many/debs
```

But primary use cases:

```plain
. <(install-debs-in-order path/to/many/debs)
```

Or

```plain
install-debs-in-order path/to/many/debs > install.sh
./install.sh
```