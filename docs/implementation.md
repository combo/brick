# Clean room implementation

## 1. Spawning the clean room

First, a rootfs for the container is constructed, by creating an [overlayfs](https://www.kernel.org/doc/Documentation/filesystems/overlayfs.txt) merging all the dependencies (including build dependencies) as lower layers, and an empty `output directory` as upper layer.

All changes done in the rootfs will be recorded in the output directory, which we will then use to create the package. The lower layers (dependencies) will be unaffected by changes.

Then, a container is spawned with this rootfs mounted read-write, and the package directory (source code) mounted read-only at /input

Example:

```
    # Creating the rootfs
    sudo mount -t overlay overlay -olowerdir=/dep1-tree:/dep2-tree:/dep3-tree,upperdir=output,workdir=work rootfs/

    # Spawning the container
    sudo systemd-nspawn --bind-ro=$(pwd):/input -bD rootfs/
```

## 2. Creating the package

`brick build` will execute the build script in the container, which will take care to install the package itself into the merged rootfs. As described above, all changes on the rootfs will be recorded in the `output directory`.

TODO: hashing

A tar archive, called `<package name>-<version>-<hash>.brick` will be created as follows:
- `brick.package.toml`, copied verbatim
- `tree/`, with the contents of the output directory
