oinc - OpenShift in Container
==============================

The `oinc` project allows you to easely configure, run and manage OpenShift
server that runs in the Docker container. It does not require any configuration
steps and it is fully automated to bring you the best experience with using
OpenShift out-of-box.

### Installation

Download the [binary](https://github.com/mfojtik/oinc/releases/download/v0.0.1/oinc-linux-amd64) from the Release page (currently only linux-amd64), or:

```console
$ git clone https://github.com/mfojtik/oinc
$ cd oinc && make install
```

### Usage

The `oinc` provides this commands:

`$ oinc setup`

This command does:

* Configure the `/etc/sysconfig/docker` (add internal registry)
* Disable firewalld (optional, ignore if you don't use it)
* Restart Docker daemon
* Pull all OpenShift Docker images needed for running OpenShift (registry, deployer, S2I builder, etc...)
* Download latest release CLI tools

`$ oinc run`

This command does:

* Create all host directories for configuration, volumes, etc...
* Fix permissions on them (SELinux)
* Start OpenShift server
* Create OpenShift Docker Registry
* Create HAProxy router

`$ oinc clean`

This command does

* Cleanup all host directories
* Kill and remove OpenShift container

