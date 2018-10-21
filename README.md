# Kira
Kira micro framework

# Default example
```
package main

import (
	"github.com/Lafriakh/kira"
)

func main() {
	app := kira.New()

	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello kira :)"))
	})
	app.Run()
}

```
# Response
json
```
app.JSON(w, data, 200)
```
render template
```
app.Render(w, data,"template/path")
```
# Middleware
```
// Log - log middleware
type Log struct{}

func NewLogger() *Log {
	return &Log{}
}

// Handler - middleware handler
func (l *Log) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var now = time.Now()

		// logger message
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(now),
		)

		next.ServeHTTP(w, r)
	})
}

// Pattern - middleware
func (l *Log) Pattern() []string {
	return []string{"*"}
}

// Name - middleware
func (l *Log) Name() string {
	return "logger"
}

```
## TODO

- [ ] Command-line interface (CLI)
- [ ] Live Reload
