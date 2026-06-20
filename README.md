# Bowerbird

Bowerbird is a WordPress plugin update install tool. It queries a Jira API to find and simplify the process of updating WordPress plugins, while still tracking them via Jira tickets. Meant for an environment where strict version control is needed. Named after an industrious creature who excels at building.

## 📚 Prerequisites

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

An selection of `json` files to enable authorized Jira API querying, and everything needed to push plugin update files to a repository (see `jsons` folder for reference).

## 🔩 Function

Bowerbird searches the targeted Jira API for tickets marked as **"New"** (aka ToDo), with a summary containing the `wordpress-plugin` vendor tag. It then gathers the qualifying candidates and runs a series of `composer require` commands on the ***composer.json*** file and pushes the updates to a designated branch.

## 📂 Project Structure

Inside of your Bowerbird project, you'll see the following folders and files:

``` zsh
.
├── cmd/
│   ├── main.go
├── internal/
│   ├── tools.go
│   ├── vars.go
├── jsons/
│   ├── definitions.json
│   ├── jira.json
├── workers/
│   ├── packagist.go
│   ├── tasks.go
├── .gitignore
├── go.mod
├── LICENSE.md
├── README.md
```

## 🚧 Build

Before building the application, change the value of the `meta` and `temp` constants to reflect your environment:

``` go
meta     string = "/data/automation/jsons/"
temp     string = "/data/automation/temp/"
```

Then, from the root folder containing `main.go`, run:

``` zsh
go build -o bowerbird .
```

## 🏃 Run

``` zsh
bowerbird -u
```

## 🎏 Available Flags

| Command               | Action                      |
|:----------------------|:----------------------------|
|    `-h, -help`        |   Help information          |
|    `-u, -update`      |   Run main program          |
|    `-v, -version`     |   Display program version   |

## 🎫 License

Code is distributed under [The Unlicense](https://github.com/farghul/bowerbird/blob/main/LICENSE.md) and is part of the Public Domain.