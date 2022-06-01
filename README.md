# API CMR
> Use API CMR (Code + Message + Result) Specification, like [jSend](https://github.com/omniti-labs/jsend).

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/api-cmr)](https://pkg.go.dev/github.com/go-zoox/api-cmr)
[![Build Status](https://github.com/go-zoox/api-cmr/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/api-cmr/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/api-cmr)](https://goreportcard.com/report/github.com/go-zoox/api-cmr)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/api-cmr/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/api-cmr?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/api-cmr.svg)](https://github.com/go-zoox/api-cmr/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/api-cmr.svg?label=Release)](https://github.com/go-zoox/api-cmr/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/api-cmr
```

## Getting Started

```go
package main

import (
  "github.com/go-zoox/api-cmr"
)

func main(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(cmr.Success("Hello World!"))
		// or
		cmr.WriteSuccess(w, "Hello World!")
	})

	http.ListenAndServe(":8080", nil)
}
```

## Inspired By
* [omniti-labs/jsend](https://github.com/omniti-labs/jsend) - JSend is a specification for a simple, no-frills, JSON based format for application-level communication

## License
GoZoox is released under the [MIT License](./LICENSE).
