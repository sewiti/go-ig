# go-ig

[![Go Reference](https://pkg.go.dev/badge/github.com/sewiti/go-ig.svg)](https://pkg.go.dev/github.com/sewiti/go-ig)
[![Build](https://github.com/sewiti/go-ig/actions/workflows/build.yml/badge.svg)](https://github.com/sewiti/go-ig/actions/workflows/build.yml)

Instagram client for extracting basic data.

It works by download HTML page of the profile and extracting data from `script`
tag.

**NOTE: this provides only a small amount of data and is not reliable since it
depends on what initial HTML does Instagram serve.**

## Install

```sh
go get -u github.com/sewiti/go-ig
```

## Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/sewiti/go-ig"
)

func main() {
	const username = "instagram"
	profile, posts, err := ig.Get(context.Background(), username)
	if err != nil {
		panic(err)
	}
	fmt.Printf("extracted %d %s's posts", len(posts), profile.Identifier.Value)
	// Output:
	// extracted 9 instagram's posts
}
```