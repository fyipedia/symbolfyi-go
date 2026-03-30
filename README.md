# symbolfyi

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/symbolfyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/symbolfyi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Symbol encoding in 11 formats and Unicode property lookup — API client for [symbolfyi.com](https://symbolfyi.com).

> **Try the interactive tools at [symbolfyi.com](https://symbolfyi.com)**

## Install

`go get github.com/fyipedia/symbolfyi-go`

## Quick Start

```go
package main

import (
    "fmt"
    "fyipedia/symbolfyi-go"
)

func main() {
    client := symbolfyi.NewClient()
    result, err := client.Search("infinity")
    if err != nil {
        panic(err)
    }
    fmt.Println(result.Total, "results")
}
```

## Also Available

| Platform | Package | Link |
|----------|---------|------|
| **Python** | `pip install symbolfyi` | [PyPI](https://pypi.org/project/symbolfyi/) |
| **npm** | `npm install symbolfyi` | [npm](https://www.npmjs.com/package/symbolfyi) |
| **Go** | `go get github.com/fyipedia/symbolfyi-go` | [pkg.go.dev](https://pkg.go.dev/github.com/fyipedia/symbolfyi-go) |
| **Rust** | `cargo add symbolfyi` | [crates.io](https://crates.io/crates/symbolfyi) |
| **Ruby** | `gem install symbolfyi` | [rubygems](https://rubygems.org/gems/symbolfyi) |

## Embed Widget

Embed [SymbolFYI](https://symbolfyi.com) widgets on any website with [symbolfyi-embed](https://widget.symbolfyi.com):

```html
<script src="https://cdn.jsdelivr.net/npm/symbolfyi-embed@1/dist/embed.min.js"></script>
<div data-symbolfyi="entity" data-slug="infinity"></div>
```

Zero dependencies · Shadow DOM · 4 themes (light/dark/sepia/auto) · [Widget docs](https://widget.symbolfyi.com)

## Links

- [SymbolFYI](https://symbolfyi.com) — Main site
- [API Documentation](https://symbolfyi.com/developers/)
- [OpenAPI Spec](https://symbolfyi.com/api/openapi.json)
- [Glossary](https://symbolfyi.com/glossary/)

## License

MIT
