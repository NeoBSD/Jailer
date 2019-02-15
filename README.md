# Jailer

The goal of Jailer is to solve the same problems that Docker solved on Linux, but for FreeBSD. Imagine a project where next to your already existing `Dockerfile` you have a `Jailerfile`, which can be used to start a FreeBSD jail running your installed & configurated application.

## Warning

```
This software is in it's infancy. Use at your own risk. Have fun.
```

## Quick Start

### Jailerfile

A container can be described in a `Jailerfile`. The goal is to keep the syntax as similar to a `Dockerfile` as possible.

**Example**:

```docker
FROM freebsd
RUN echo
```

### Supported commands

```sh
jailer version          # Prints jailer version info
jailer top $CONTAINER   # Prints a one time output of top in a container
```

## Development

### Dependencies

- Golang with modules
- gmake

### Setup

```sh
zfs create $ZROOT/jailer
```
