package engine

import (
	"errors"
	"sync"
	"time"
)

// Conveyor 传送带上可以滚动任何零件
type Conveyor interface {
	// GetSelf 返回传送带本身
	GetPipeline() chan interface{}

	// AddWorker 分配一个工人在这个传送带上工作，一个传送带上可以分配多个工人
	AddWorker(w Worker, n int) error

	// 向流水线上放零件，如果满了就一直等到有空间在放
	PutPart(p Part) error

	// Run 工人开始在这个流水线上工作
	Run() error

	// Stop 传送带不在接收新的零件，等待工人把传送带上的工作做完
	Stop()
}

type emptyConveyor struct {
	pipeline chan interface{}
	workers  []Worker
	running  bool
	group    sync.WaitGroup
	mutex    sync.Mutex
}

// ErrLineIsFull 流水线已满错误，需要业务端判断cpu负载，如果可以就增加工人
var ErrLineIsFull = errors.New("任务太多了，再加一个机器吧")

func (e *emptyConveyor) GetPipeline() chan interface{} {
	return e.pipeline
}

func (e *emptyConveyor) Run() error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for _, w := range e.workers {
		w.Working(e.pipeline, &e.group)
		e.group.Add(1)
	}

	e.running = true
	return nil
}

func (e *emptyConveyor) AddWorker(w Worker, n int) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for i := 0; i < n; i++ {
		e.workers = append(e.workers, w)
	}
	if e.running {
		w.Working(e.pipeline, &e.group)
	}
	return nil
}

func (e *emptyConveyor) PutPart(p Part) error {
	tm := time.NewTimer(5 * time.Second)
	select {
	case e.pipeline <- p:
	case <-tm.C:
		return ErrLineIsFull
	}
	return nil
}

func (e *emptyConveyor) Stop() {
	close(e.pipeline)
	e.group.Wait()
}

func NewConveyor(cap int) Conveyor {
	c := new(emptyConveyor)
	c.pipeline = make(chan interface{}, cap)
	return c
}
