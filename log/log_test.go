package log

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	tl := New(os.Stdout, "TEST", LInfo, Ldefault)
	tl.Errorf("this should be print msg -> %s\n", "hahaha")
	tl.Debugf("this shouldn't print msg -> %s\n", "ememem")
}

func TestDebug(t *testing.T) {
	Debug("not print")
	Info("print")
}
