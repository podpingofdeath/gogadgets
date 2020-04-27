package resources

import "testing"

type T struct {
	count int
	buf   chan string
}

func (t *T) Init() {
	t.count = 5
	t.buf = make(chan string, t.count)
}

func (t *T) Free() {
	t.count = 0
	close(t.buf)
}

func TestRes(t *testing.T) {
	var res T
	Enroll(&res)
	if len(resMgrInst.reses) != 1 {
		t.Error("Enroll error")
	}

	Init()
	if res.count != 5 {
		t.Error("Init error")
	}

	Free()
	if res.count != 0 {
		t.Error("Free error")
	}
}
