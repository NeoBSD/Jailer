# jailer-compose

## TOC

- [Example](#example)
- [Spec Top-Level](#spec-top-level)
  - [services](#services)
  - [version](#version)
- [Spec Service](#spec-service)
  - [command](#command)
  - [cpu_percent](#cpu-percent)
  - [cpu_threads](#cpu-threads)
  - [depends_on](#depends-on)
  - [hostname](#hostname)
  - [working_dir](#working-dir)

## Example

```yml
---
version: 0.1
services:
  - label: web
    image: nginx
  - label: db
    image: postgres
```

## Spec Top-Level

### services

TYPE: **string**

`services` property is defined by the specification for backward compatibility but is only informative.

### version

TYPE: **string**

`version` property is defined by the specification for backward compatibility but is only informative.

## Spec Service

### command

TYPE: **string**

`command` overrides the the default command declared by the jail image (i.e. by Jailerfile's CMD).

```yml
command: "bundle exec thin -p 3000"
```

### cpu_percent

TYPE: **int**

cpu_percent defines the usable percentage of the available CPUs.

```sh
# %CPU, in percents of a single CPU core
rctl -a jail:<jailname>:pcpu:deny=25
```

### cpu_threads

TYPE: **int**

`cpu_threads` defines the number of usable threads for a service jail.

```sh
rctl -a jail:<jailname>:nthr:deny=50
```

### depends_on

TYPE: **array of string**

`depends_on` expresses startup and shutdown dependencies between services.
Short syntax

The short syntax variant only specifies service names of the dependencies. Service dependencies cause the following behaviors:

```text
Compose implementations MUST create services in dependency order. In the following example, db and redis are created before web.

Compose implementations MUST remove services in dependency order. In the following example, web is removed before db and redis.
```

Simple example:

```yml
services:
  - label: web
    image: nginx
    depends_on:
      - db
  - label: db
    image: postgres
```

Compose implementations **MUST** guarantee dependency services have been started before starting a dependent service. Compose implementations MAY wait for dependency services to be "ready" before starting a dependent service.

### hostname

TYPE: **string**

`hostname` declares a custom host name to use for the service jail.

### image

TYPE: **string**

`image` specifies the image to start the jail from.

### label

TYPE: **string**

`label` declares a name to use for the service jail.

### working_dir

TYPE: **string**

`working_dir` overrides the container's working directory from that specified by image (i.e. Jailerfile `WORKDIR`).
