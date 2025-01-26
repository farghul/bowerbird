# Bowerbird

Bowerbird is a WordPress plugin update install tool. It queries a Jira API to find and simplify the process of updating WordPress plugins, while still tracking them via Jira tickets. Meant for an environment where strict version control is needed. Named after an industrious creature who excels at building.

![Bird](bowerbird.webp)

## Prerequisites

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

An selection of `json` files to enable authorized Jira API querying, and everything needed to push plugin update files to a repository (see `jsons` folder for reference).

## Function

Bowerbird searches the targeted Jira API for tickets marked as **"New"** (aka ToDo), with a summary containing the `wordpress-plugin` vendor tag. It then gathers the qualifying candidates and runs a series of `composer require` commands on the ***composer.json*** file and pushes the updates to a designated branch.

## Build

Before building the application, change the value of the `assets` constant to reflect your environment:

``` go
assets string = "/data/scripts/automation/assets/"
```

Then, from the root folder containing `main.go`, use the command that matches your environment:

### Windows & Mac:

``` zsh
go build -o [name] .
```

### Linux:

``` zsh
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

``` zsh
[program] [flag]
```

## Optional Flags

``` zsh
-h, --help        Help information
-v, --version     Display program version
```

## Example

``` zsh
bowerbird -h
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/bowerbird/blob/main/LICENSE.md) and is part of the Public Domain.
