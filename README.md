# Bowerbird

Bowerbird is a WordPress plugin update install tool. It queries a Jira API to find and simplify the process of updating WordPress plugins, while still tracking them via Jira tickets. Meant for an environment where strict version control is needed. Named after an industrious creature who excels at building.

![Bird](bowerbird.webp)

## Prerequisites

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

An selection of `json` files to enable authorized Jira API querying, and everything needed to push plugin update files to a repository (see `jsons` folder for reference).

## Function

Bowerbird searches the targeted Jira API for tickets marked as **"New"** (aka ToDo), and filtered with labels such as *wordpress-plugin*. It then gathers the qualifying candidates and runs a series of `composer require` commands on the ***composer.json*** file and pushes the updates to a designated branch.

## Build

From the root folder containing `main.go`, use the command that matches your environment:

### Windows & Mac:

``` zsh
go build -o [name] .
```

### Linux:

``` zsh
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

Ensure the folder containing your ***composer.json*** file is predefined as variable and run:

``` zsh
[program] [flag]
```

## Options

``` zsh
-c, --core        Install WordPress core updates
-h, --help        Help information
-v, --version     Display program version
-w, --wpackagist  Install free WPackagist updates
```

## Example

``` zsh
bowerbird -w
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/bowerbird/blob/main/LICENSE.md) and is part of the Public Domain.
