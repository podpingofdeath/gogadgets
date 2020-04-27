package main

// resources 包的用例
import (
	"fmt"
	"github.com/podpingofdeath/gogadgets/resources"
	"time"
)

type T struct {
	count int
	buf   chan string
}

func (t *T) Init() {
	t.count = 5
	t.buf = make(chan string, t.count)
}

func (t *T) Free() {
	close(t.buf)
}

var res T

func init() {
	resources.Enroll(&res)
}

func main() {
	resources.Init()

	go func() {
		time.Sleep(time.Second * 5) //这里正常应该是接收外部信号
		resources.Free()
	}()

	for s := range res.buf {
		fmt.Println(s)
	}
	fmt.Println("buf is closed!")
}
