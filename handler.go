package log15ter

import "github.com/inconshreveable/log15"

// TernaryHandler writes all log records to the `yes` handler if the function
// is true. Otherwise it will write logs to the `no` handler.
func TernaryHandler(fn func(r *log15.Record) bool, yes log15.Handler, no log15.Handler) log15.Handler {
	return log15.FuncHandler(func(r *log15.Record) error {
		if fn(r) {
			return yes.Log(r)
		}

		return no.Log(r)
	})
}

// LvlTernaryHandler writes all logs above the log level provide to the `yes`
// handler. Otherwise it will write the logs to the `no` handler. For example
// you might want to add stack traces to logs only if the log level is error or
// above:
//
//     log15ter.LvlTernaryHandler(
//         log15.LvlError,
//         log15.CallerStackHandler("%+v", log15.StdoutHandler),
//         log15.StdoutHandler)
//
func LvlTernaryHandler(minLvl log15.Lvl, yes log15.Handler, no log15.Handler) log15.Handler {
	return TernaryHandler(func(r *log15.Record) bool {
		return r.Lvl <= minLvl
	}, yes, no)
}
