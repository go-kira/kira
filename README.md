<p align="center">
  <h1 align="center">Kira</h1>
  <p align="center">Minimal web framework.</p>
</p>

<p align="center">
  <a href="#features">Features</a> |
  <a href="#installation">Installation</a> |
  <a href="#getting-started">Getting Started</a> |
  <a href="#examples">Examples</a> |
  <a href="#docs">Docs</a> <br/>
  <a href="https://travis-ci.com/go-kira/kira"><img src="https://api.travis-ci.com/go-kira/kira.svg?branch=master" alt="Build Status"></a>
  <a href="https://codecov.io/gh/go-kira/kira"><img src="https://codecov.io/gh/go-kira/kira/branch/master/graph/badge.svg" alt="Code Coverage"/></a>
</p>

---

<p align="center">
  <p align="center"><b>Kira</b> web framework. Simply a minimal web framework.</p>
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

MIT License

Copyright (c) 2019 Rachid Lafriakh

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.