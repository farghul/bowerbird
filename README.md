# Bowerbird

Bowerbird is a WordPress plugin update install tool. It queries a Jira API to find and simplify the process of updating WordPress plugins, while still tracking them via Jira tickets. Meant for an environment where strict version control is needed. Named after a very industrious little creature who excels at building.

![Bird](bowerbird.webp)

## Prerequisites

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

Login information to download update packages. -- ***premium content only*** --

A `jsons/access.json` file containing your API URL and Basic token to enable authorized querying:

``` json
{
    "dev": "Internal Development search keyword",
    "wpac": "WordPress Packagist search keyword",
    "prem": "Premium Subscription search keyword",
    "root": "Path to folder containing all Premium repositories",
    "repo": "Path to main repository",
    "prod": "Jira Issue base URL",
    "testing": "JQL search string to query the Jira API",
    "token": "Email:Jira API Token combination with Base 64 Encoding"
}
```

## Function

Bowerbird searches the targeted Jira API for tickets marked as **"New"** (aka ToDo), and filtered with labels such as *wordpress-plugin*. It then gathers the qualifying candidates and runs a series of `composer require` commands on the ***composer.json*** file and pushes the updates to a designated test branch. Additional steps such as downloading update files and version tagging may be performed for premium or in-house content prior to the push.


## Build

From the root folder containing the `go` files, use the command that matches your environment:

### Windows & Mac:

``` console
go build -o [name] .
```

### Linux:

``` console
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

Ensure the folder containing your ***composer.json*** file is predefined as variable and run:

``` console
[program] [flag]
```

## Options

``` console
-h, --help        Help Information
-d, --developer   Install internal developer updates
-p, --premium     Install paid subscription updates
-v, --version     Display Program Version
-w, --wpackagist  Install free wpackagist updates
```

## Example

``` console
bowerbird -w
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/bowerbird/blob/main/LICENSE.md) and is part of the Public Domain.
