# Bowerbird

Bowerbird is a WordPress plugin update install tool. It queries the target Jira API to find 

It simplifies the process of pulling free updates from [WordPress Packagist](https://wpackagist.org) or premium updates from a subscription, then pushing them to Git and adding version tags to a private repository, if needed. Named after a very industrious creature who excels at building.

![Bird](bowerbird.webp)

## Prerequisites

Login information to download the update package. -- ***premium content only*** --

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

A `secrets/jira.json` file containing your API URL and Basic token to enable ticket creation:

``` go
{
    "dev": "Internal Development search filter",
    "free": "WordPress Packagist search filter",
    "paid": "Subscription search filter",
    "base": "Jira Issue base URL",
    "search": "JQL search string to query the Jira API",
    "token": "Email:Jira API Token combination with Base 64 Encoding"
}
```

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
[program] [flag] [vendor/plugin]:[version] [ticket#]
```

## Examples

Currently there are three supported scenarios available, using flags to specify each.

### Subscription (Purchased third party plugins):

``` console
bowerbird -s bcgov-plugin/events-virtual:1.13.4 795
```

### Packaged (Plugins currently available via WPackagist or Satis):

``` console
bowerbird -p wpackagist-plugin/mailpoet:4.6.1 821
```

### Release (In-house Production ready content):

``` console
bowerbird -r bcgov-plugin/bcgov-inline-comments:1.9.0 820
```

Flags `-r` and `-p` can accept multiple updates, chain together as many as you like!

## License

Code is distributed under [The Unlicense](https://github.com/farghul/bowerbird/blob/main/LICENSE.md) and is part of the Public Domain.
