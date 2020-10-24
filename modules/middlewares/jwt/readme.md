<p align="center">
  <h1 align="center">Kira</h1>
  <p align="center">JWT for kira framework</p>
</p>

<p align="center">
  <a href="#getting-started">Getting Started</a> |
  <a href="#configuration">Configuration</a>
</p>

---

## Getting Started

    package main

    import (
        "github.com/go-kira/kira"
	    "github.com/go-kira/kira/middlewares/jwt"
    )

    func main() {
        app := kira.New()
        app.Use(csrf.NewCSRF()) // We use the middleware globally

        app.Get("/jwt", func(ctx *kira.Context) {
            token, err := jwt.CreateToken(ctx, map[string]interface{}{
                "foo": "bar",
            })
            if err != nil {
                ctx.Error(err)
            }

            ctx.Stringf("JWT token: %s", token)
        })

        app.Post("/jwt", func(ctx *kira.Context) {
            ctx.String("Protected.")
        })

        app.Run()
    }

## Configuration

You can configure this middleware from your `config.toml` file.

- **[app.key]:** It's used to encrypt the JWT token.
- **[jwt.lookup]:** You can shouls which method to use in the lookup: `<method>:<name>`. By default it's `header:Authorization`

You can use in the lookup only: `header, cookie`. Maybe in the future we support more methods like `query`...
