<p align="center">
  <h1 align="center">Kira</h1>
  <p align="center">Minimal web framework.</p>
</p>

<p align="center">
  <a href="#features">Features</a> |
  <a href="#installation">Installation</a> |
  <a href="#getting-started">Getting Started</a> |
  <a href="#examples">Examples</a> |
  <a href="#docs">Docs</a>
  <a href="https://travis-ci.com/go-kira/kira"><img src="https://api.travis-ci.com/go-kira/kira.svg?branch=master" alt="Build Status"></a>
  <a href="https://codecov.io/gh/go-kira/kira"><img src="https://codecov.io/gh/go-kira/kira/branch/master/graph/badge.svg" alt="Code Coverage"/></a>
</p>

---

<p align="center">
  <p align="center">**Kira** web framework. Simply a minimal web framework.</p>
</p>

## Features

- **Simplicity** kira is simple. You will find yourself familiar with it quickly.
- **Fast** simplicity comes with speed. Kira is super fast.

## Installation

    go get -u github.com/go-kira/kira

## Getting Started

```go
package main

import "github.com/go-kira/kira"

func main() {
    app := kira.New()

    app.Get("/", func (c *kira.Context) {
        c.String("Hello, Kira :)")
    })

    app.Run()
}
```

## License

Copyright 2019 Lafriakh Rachid <lafriakh.rachid@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
