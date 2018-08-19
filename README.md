# log15ter [![godoc reference](https://godoc.org/github.com/calebcase/log15ter?status.png)](https://godoc.org/github.com/calebcase/log15ter) [![Build Status](https://travis-ci.org/calebcase/log15ter.svg?branch=master)](https://travis-ci.org/calebcase/log15ter)

Package log15ter extends [log15](https://github.com/inconshreveable/log15) with
ternary handlers. A ternary handler allows for splitting the logic flow of the
handler pipeline based on arbitrary decision points. This can be useful for
conditionally adding more information to logs when it makes sense.

For example, the following will only call the stack trace handler if the log
level is `ERROR` or above:

```go
log15ter.LvlTernaryHandler(
	log15.LvlError,
	log15.CallerStackHandler("%+v", log15.StdoutHandler),
	log15.StdoutHandler)
```
