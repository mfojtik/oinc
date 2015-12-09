oinc - OpenShift in Container
==============================

The `oinc` project allows you to easily configure, run and manage the OpenShift
server that runs in a Docker container. It does not require any configuration
steps and it is fully automated to bring you the best experience with using
OpenShift out-of-box.

[![asciicast](https://asciinema.org/a/1pe3bed84u8w6zon5a1tt8j6f.png)](https://asciinema.org/a/1pe3bed84u8w6zon5a1tt8j6f)


### Installation

Download the [binary](https://github.com/mfojtik/oinc/releases/download/v0.0.2/oinc-linux-amd64) from the Release page (currently only linux-amd64), or:

```console
$ git clone https://github.com/mfojtik/oinc
$ cd oinc && make install
```

### Usage

The `oinc` tool provides several commands:

`$ oinc setup`

* Configures the `/etc/sysconfig/docker` file (adds the internal registry)
* Disables firewalld (optional, ignore if you don't use it)
* Restarts Docker daemon
* Pulls all OpenShift Docker images needed for running OpenShift (registry, deployer, S2I builder, etc...)
* Downloads the latest release CLI tools

`$ oinc run`

* Creates all host directories for configuration, volumes, etc...
* Fixes permissions on them (SELinux)
* Starts OpenShift server
* Creates the OpenShift Docker Registry
* Creates the HAProxy router

`$ oinc clean`

* Cleans up all host directories
* Kills and removes the OpenShift container

