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



