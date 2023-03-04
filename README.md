# WeaverUtils

[![Go Reference](https://img.shields.io/badge/go-reference-%23007d9c?style=for-the-badge&logo=go)](https://pkg.go.dev/gopkg.eu.org/weaverutils)

[Service Weaver](https://serviceweaver.dev/) Utilities

## Installation

```shell
go get -u gopkg.eu.org/weaverutils
```

## Components

- [WeaverFastCache](#weaverfastcache)

### WeaverFastCache

WeaverFastCache is a [FastCache](https://github.com/VictoriaMetrics/fastcache) component for Service Weaver.

Routing Key: xxhash64(key)
