package middlewares

import (
	"github.com/go-kira/kira"
	"github.com/go-kira/kira/middlewares/limitbody"
	"github.com/go-kira/kira/middlewares/logger"
	"github.com/go-kira/kira/middlewares/recover"
	"github.com/go-kira/kira/middlewares/requestid"
)

// Default - Register the default middlewares, ex: recover, logger, requestid, limitbody.
func Default(app *kira.App) {
	middlewares := []kira.Middleware{
		recover.NewRecover(app),
		logger.NewLogger(app),
		limitbody.Newlimitbody(app),         // body limit.
		requestid.NewRequestID(app.Configs), // keep this the last one.
	}

	app.UseMiddlewares(middlewares)
}
