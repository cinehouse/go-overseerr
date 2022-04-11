# go-overseerr #

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cinehouse/go-overseerr)](https://github.com/cinehouse/go-overseerr/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/cinehouse/go-overseerr/v1/overseerr)
[![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/cinehouse/go-overseerr/tests/main?label=tests)](https://github.com/cinehouse/go-overseerr/actions?query=workflow%3Atests)
[![Go Report Card](https://goreportcard.com/badge/github.com/cinehouse/go-overseerr)](https://goreportcard.com/report/github.com/cinehouse/go-overseerr)
[![Codecov branch](https://img.shields.io/codecov/c/github/cinehouse/go-overseerr/main?token=p78MbVUq1e)](https://codecov.io/gh/cinehouse/go-overseerr)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=cinehouse_go-overseerr&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=cinehouse_go-overseerr)
[![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/cinehouse/go-overseerr/main?label=Go)](https://golang.org/doc/install)
[![Go Reference](https://pkg.go.dev/badge/github.com/cinehouse/go-overseerr.svg)](https://pkg.go.dev/github.com/cinehouse/go-overseerr)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/cinehouse/go-overseerr/blob/main/LICENSE)

go-overseerr is a Go client library for accessing the [Overseerr API][].

Currently, **go-overseerr requires Go version 1.17 or greater**.  go-github tracks
[Go's version support policy][support-policy].  We do our best not to break
older versions of Go if we don't have to, but due to tooling constraints, we
don't always test older versions.

[support-policy]: https://golang.org/doc/devel/release.html#policy

## Installation ##

go-overseerr is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/cinehouse/go-overseerr/v1
```
will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/cinehouse/go-overseerr/v1/overseerr"
```

and run `go get` without parameters.

Finally, to use the top-of-trunk version of this repo, use the following command:

```bash
go get github.com/cinehouse/go-overseerr/v1@main
```

[Overseerr API]: https://api-docs.overseerr.dev/
