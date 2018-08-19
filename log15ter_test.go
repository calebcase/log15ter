package log15ter

import (
	"testing"

	"github.com/inconshreveable/log15"
)

func testHandler() (log15.Handler, *log15.Record) {
	rec := new(log15.Record)
	return log15.FuncHandler(func(r *log15.Record) error {
		*rec = *r
		return nil
	}), rec
}

func TestTernaryHandler(t *testing.T) {
	t.Parallel()

	l := log15.New()
	yes, yes_r := testHandler()
	no, no_r := testHandler()

	alwaysTrue := func(r *log15.Record) bool {
		return true
	}

	alwaysFalse := func(r *log15.Record) bool {
		return false
	}

	l.SetHandler(TernaryHandler(alwaysTrue, yes, no))

	l.Debug("test ok")
	if no_r.Msg != "" {
		t.Fatalf("expected only yes handler to run")
	}
	if yes_r.Msg == "" {
		t.Fatalf("expected yes handler to run")
	}

	yes, yes_r = testHandler()
	no, no_r = testHandler()
	l.SetHandler(TernaryHandler(alwaysFalse, yes, no))

	l.Debug("test ok")
	if yes_r.Msg != "" {
		t.Fatalf("expected only no handler to run")
	}
	if no_r.Msg == "" {
		t.Fatalf("expected no handler to run")
	}
}

func TestLvlTernaryHandler(t *testing.T) {
	t.Parallel()

	l := log15.New()
	yes, yes_r := testHandler()
	no, no_r := testHandler()

	l.SetHandler(LvlTernaryHandler(log15.LvlError, yes, no))

	l.Debug("test ok")
	if no_r.Msg == "" || yes_r.Msg != "" {
		t.Fatalf("expected no handler to run for debug level")
	}

	yes, yes_r = testHandler()
	no, no_r = testHandler()
	l.SetHandler(LvlTernaryHandler(log15.LvlDebug, yes, no))

	l.Debug("test ok")
	if no_r.Msg != "" || yes_r.Msg == "" {
		t.Fatalf("expected yes handler to run for debug level")
	}
}
