package middlewares

import (
	"github.com/go-kira/kira"
	"github.com/go-kira/kira/middlewares/csrf"
	"github.com/go-kira/kira/middlewares/limitbody"
	"github.com/go-kira/kira/middlewares/logger"
	"github.com/go-kira/kira/middlewares/recover"
	"github.com/go-kira/kira/middlewares/requestid"
)

// Default - Register the default middlewares, ex: recover, logger, requestid, limitbody.
func Default(app *kira.App) {
	// app.Use(jwt.New())
	app.Use(limitbody.New())
	app.Use(csrf.New())
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(logger.New())
}
