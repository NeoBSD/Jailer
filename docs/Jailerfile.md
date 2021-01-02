# Jailerfile

- [COPY](#copy)
- [FROM](#from)
- [LABEL](#label)
- [RUN](#run)
- [SHELL](#shell)
- [WORKDIR](#workdir)

## COPY

```docker
COPY <src> <dest>
```

The `COPY` instruction copies new files or directories from `<src>` and adds them to the filesystem of the jail at the path `<dest>`.

## FROM

```docker
FROM <image>
```

The `FROM` instruction selects the `<image>` to base the current build on. It searches in the `JAILER_PATH` env variable.

## LABEL

```docker
LABEL <key>=<value>

# EXAMPLE
LABEL maintainer="example@example.com"
LABEL version="1.0"
```

The `LABEL` instruction adds metadata to an image. A `LABEL` is a key-value pair. Multi instances are allowed.

## RUN

```docker
RUN <command>
```

The `RUN` instruction will execute any commands in the default shell. See `SHELL`

## SHELL

```docker
SHELL ["executable", "parameters"]
```

The `SHELL` instruction allows the default shell used for the shell form of commands to be overridden. The default shell is `["/bin/sh", "-c"]`.

The `SHELL` instruction can appear multiple times. Each `SHELL` instruction overrides all previous `SHELL` instructions, and affects all subsequent instructions. For example:

## WORKDIR

```docker
WORKDIR /path/to/workdir
```

The `WORKDIR` instruction sets the working directory for any `RUN`, `CMD` and `COPY` instructions that follow it in the `Jailerfile`. If the `WORKDIR` doesn’t exist, it will be created even if it’s not used in any subsequent `Jailerfile` instruction.
