# Jailer

## TLDR

The goal of Jailer is to solve the same problems that Docker solved on Linux, but for FreeBSD. Imagine a project where next to your already existing `Dockerfile` you have a `Jailerfile`, which can be used to start a FreeBSD jail running your installed & configurated application.

## Warning

```text
This software is in it's infancy. Use at your own risk. Have fun.
```

## Status

|                                                                  License                                                                  |                                                           Unit test                                                           |                                                               Coverage                                                               |                                      Issues                                      |
| :---------------------------------------------------------------------------------------------------------------------------------------: | :---------------------------------------------------------------------------------------------------------------------------: | :----------------------------------------------------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------: |
| [![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://github.com/tobiashienzsch/jailer/blob/master/LICENSE) | [![Build Status](https://travis-ci.org/tobiashienzsch/jailer.svg?branch=master)](https://travis-ci.org/tobiashienzsch/jailer) | [![codecov](https://codecov.io/gh/tobiashienzsch/jailer/branch/master/graph/badge.svg)](https://codecov.io/gh/tobiashienzsch/jailer) | ![GitHub issues](https://img.shields.io/github/issues/tobiashienzsch/jailer.svg) |

## Quick Start

### Jailerfile

A container can be described in a `Jailerfile`. The goal is to keep the syntax as similar to a `Dockerfile` as possible.

**Example**:

```docker
FROM freebsd
RUN echo "Hello Jailer!"
```

### Supported commands

Most commands need to run as `root`:

```sh
jailer help
jailer version
jailer config
jailer storage
jailer top $CONTAINER
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
cp config.example.yml config.yml
zfs create -o mountpoint=/jailer $ZROOT/jailer
```

### Resources

#### Internal

- [Jail from scratch](https://github.com/tobiashienzsch/jailer/blob/master/docs/jail_from_scratch.md)

#### External

- [Reddit: FreeBSD 12 + VNET + ZFS](https://www.reddit.com/r/freebsd/comments/ahdbbq/howto_jails_freebsd_12_vnet_zfs/)
- [Devin Teske - Jail Networking, MeetBSD 2016](https://www.youtube.com/watch?v=aoW7pWuhT_A)
- [FreeBSD 12 Jail with IPv6, VNET and ZFS](https://medium.com/@melanj/how-to-configure-a-globally-routable-freebsd-12-jail-with-ipv6-vnet-and-zfs-4c750ef31b80)
- [Stefan Grönke: Imprisoning software with libiocage -- BSDCan 2018](https://www.youtube.com/watch?v=CTGc3zYToh0)
