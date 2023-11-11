# Introduction

This is my third attempt.

I just want made an easy to use, easy to maintenance template.

So far, I am very satisfied with this project.

# Features

- wire - injects dependencies
- ent - database orm
- gin - router
- koanf - config manager
- zap - the logger
- and so on...
    - mail
    - ...

# Usage

I also made a cli to init project from template and generate new route file.

So, first install the cli tool:

```bash
go install github.com/Xwudao/neter/cmd/nr@latest
```

Then, you can use the cli to init project from template:

```bash
nr init <project-name> [-m <mod-name>]
```

Then, run `neter run` to start the project.

# `nr` cli

```shell
PS C:\Users\tim> nr -h
A helper for nr

Usage:
   [command]

Available Commands:
  build       build the final binary
  completion  Generate the autocompletion script for the specified shell
  gen         gen some for nr
  help        Help about any command
  init        init new project from template
  run         run the application by command
  wire        wire the dependency

Flags:
  -h, --help     help for this command
  -t, --toggle   Help message for toggle

Use " [command] --help" for more information about a command.
```

## `nr run`

```sh
# run with `wire`
nr run -w
# run with delete bin file when stopped.
nr run -d

## you can comine all flag
nr run -dw
```

## `nr build`

```sh
nr build --trim -wlm # will build windows/linux/mac binary with trim path
# like: go build -trimpath -ldflags=-s -w -extldflags '-static' -o app-linux ./cmd/app/
```

## `nr gen`

you can use `nr gen` to generate something.


biz: `internal/biz/gen.go`

```sh
//go:generate nr gen -t biz -n diskHelper --no-repo
# will generate disk_helper_biz.go and update provider.go
```

---

cmd: `internal/cmd`

```sh
nr gen cmd -n CmdName
```

---


types: typescrit / golang api client in `build` folder which on project root.

```sh
nr gen tp --go --ts
```

---

ent:

```sh
nr gen ent
# will run `go generate ./ent -run entgo.io/ent/cmd/ent` in `internal/data` directory which is the `ent` folder.
```

## `nr wire`


will find all wire.go files in project, and try to run `wire` command

```sh
nr wire
```

# Database

In this template, I use `ent` to connect to database.

And, in newer version of ent, it supported versioned-migration.

So, in `neter-cli` I also added a command to migrate database.

```bash
# also you need config the database info in config.yaml that locate in the project root directory
# generate migration file
nr run migrate name=<migration-filename>

# up
nr run migrate up [all=true]

# down
nr run migrate down [all=true]
```



