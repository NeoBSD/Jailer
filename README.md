# Jailer

## TLDR

The goal of Jailer is to solve the same problems that Docker solved on Linux, but for FreeBSD. Imagine a project where next to your already existing `Dockerfile` you have a `Jailerfile`, which can be used to start a FreeBSD jail running your installed & configurated application.

## Warning

```text
This software is in it's infancy. Use at your own risk. Have fun.
```

## Status

|                                                               License                                                               |                                        Unit test                                         |                                                       Coverage                                                       |                                  Issues                                  |
| :---------------------------------------------------------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------------------------------------------: | :----------------------------------------------------------------------: |
| [![License](https://img.shields.io/badge/License-BSD%202--Clause-orange.svg)](https://github.com/neobsd/jailer/blob/master/LICENSE) | ![Unittest](https://github.com/NeoBSD/Jailer/workflows/Unittest/badge.svg?branch=master) | [![codecov](https://codecov.io/gh/neobsd/jailer/branch/master/graph/badge.svg)](https://codecov.io/gh/neobsd/jailer) | ![GitHub issues](https://img.shields.io/github/issues/neobsd/jailer.svg) |

## Quick Start

### Jailerfile

A jail can be described in a `Jailerfile`. The goal is to keep the syntax as similar to a `Dockerfile` as possible.

**Example**:

```docker
FROM freebsd
RUN echo "Hello Jailer!"
```

### Supported commands

Not all commands are implemented yet. Most commands need to run as `root`:

```text
Usage:
  jailer [flags]
  jailer [command]

Available Commands:
  build       Build an image from a Jailerfile
  config      Print current config
  dev         Subcommand for development only
  exec        Execute a command inside an existing jail
  fetch       Fetch base from FreeBSD mirror
  help        Help about any command
  info        Display system-wide information
  init        Init for jailer. Creates zfs datasets
  ps          List active jails
  restart     Restart one or more jails
  rm          Remove one or more jails
  run         Run a command in a new jail
  start       Start one or more stopped jails
  stop        Stop one or more running jails
  storage     Manages jail & image storage
  top         Run top inside a jail
  version     Print current version

Flags:
  -c, --config string   Config file (default is $PWD/jailer.yaml)
  -h, --help            help for jailer
  -v, --verbose         Verbose output

Use "jailer [command] --help" for more information about a command.
```

## Development

### Dependencies

- FreeBSD 12.1
  - ZFS
- Go (tested with 1.13)
  - Cobra
  - Viper
  - Logrus
- gmake

### Setup

```sh
cp jailer.example.yml jailer.yml
zfs create -o mountpoint=/jailer $ZROOT/jailer
```

### Resources

#### Internal

- [Jail from scratch](https://github.com/NeoBSD/jailer/blob/master/docs/jail_from_scratch.md)

#### External

- [Reddit: FreeBSD 12 + VNET + ZFS](https://www.reddit.com/r/freebsd/comments/ahdbbq/howto_jails_freebsd_12_vnet_zfs/)
- [Devin Teske - Jail Networking, MeetBSD 2016](https://www.youtube.com/watch?v=aoW7pWuhT_A)
- [FreeBSD 12 Jail with IPv6, VNET and ZFS](https://medium.com/@melanj/how-to-configure-a-globally-routable-freebsd-12-jail-with-ipv6-vnet-and-zfs-4c750ef31b80)
- [Stefan Grönke: Imprisoning software with libiocage -- BSDCan 2018](https://www.youtube.com/watch?v=CTGc3zYToh0)
