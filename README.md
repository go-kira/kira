<p align="center"><strong>Kira</strong></p>

<p align="center">
<a href="https://travis-ci.com/go-kira/kira"><img src="https://api.travis-ci.com/go-kira/kira.svg?branch=master" alt="Build Status"></a>
<a href="https://codecov.io/gh/go-kira/kira"><img src="https://codecov.io/gh/go-kira/kira/branch/master/graph/badge.svg" alt="Code Coverage"/></a>
</p>

Kira micro framework

# Example

    func main() {
        app := kira.New()

        app.Get("/", func (c *kira.Context) {
            c.String("Hello, Kira :)")
        })

        app.Run()
    }
