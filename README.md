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
import (
  "testing"
  "github.com/go-zoox/api-cmr"
)

func main(t *testing.T) {
	loadConfig := func() {
		panic("load config panic")
	}

	if err := api-cmr.Do(loadConfig); err != nil {
		log.Error(err)
	}
}
```

## Inspired By
* [kenkyu392/go-api-cmr](https://github.com/kenkyu392/go-api-cmr) - Provides a sandbox for the api-cmr execution of panic-inducing programs
* [go-zoox/retry](https://github.com/andskur/argon2-hashing) - Catch Panic In Retries

## License
GoZoox is released under the [MIT License](./LICENSE).
