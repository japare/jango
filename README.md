# Jango - a jan lookup library

[![GoDoc](https://godoc.org/github.com/japare/jango?status.svg)](https://godoc.org/github.com/japare/jango)

jango allows you to find japanese article numbers (JANs) for products potentially matching a series of keywords you pass to its interface.

## API stability warning
The API is still unstable and may undergo changes without any warning.

## Ethics
Depending on your configuration, this package will perform a limited amount of web scraping.
It is not intended for large-scale server-side use and is exclusively for client-side use, looking up data already freely available to users in a more efficient manner.

Irresponsible use can potentially cause a denial-of-service or incur bandwidth costs to yourself or a website provider.
The authors of the package are not responsible for any kind of misuse of the package.


## Installation

### Go

#### Go modules (1.11+)
Import the package where necessary, go will update your go.mod and go.sum for you automatically the next time you build your project.

```Go
import "github.com/japare/jango"
```

### Javascript (GopherJS)

Work in progress.

### Javascript (WASM)

Work in progress.


## Examples

Usage examples will follow soon.