// 工厂 / 车间
package engine

import (
	"errors"
	"sync"
)

type Factory interface {
	// AddLine 给工厂添加流水线
	AddLine(c Conveyor) error

	// Run 工厂开工
	Run() error

	// Stop 工厂停工
	Stop()
}

// emptyFactory 一个空的工厂类型，实现Factory接口
type emptyFactory struct {
	lines   []Conveyor
	running bool
	mutex   sync.Mutex
}

// AddLine 给工厂添加流水线，线程安全
func (e *emptyFactory) AddLine(c Conveyor) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.lines = append(e.lines, c)

	if e.running {
		return c.Run()
	}

	return nil
}

// Run 工厂开工，运行工厂中所有流水线
func (e *emptyFactory) Run() error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	var err error

	if e.running {
		return ErrFactoryRunning
	}

	for _, line := range e.lines {
		err = line.Run()
		if err != nil {
			return err
		}
	}

	e.running = true
	return nil
}

// Stop 工厂停工，但要求各个流水线上的工作要做完
func (e *emptyFactory) Stop() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for _, line := range e.lines {
		line.Stop()
	}

	e.running = false
}

var ErrFactoryRunning = errors.New("factory is already running")

// defaultFactory 默认工厂
var defaultFactory = new(emptyFactory)

// AddLine 给默认工厂添加流水线
func AddLine(c Conveyor) error {
	return defaultFactory.AddLine(c)
}

// 默认工厂开工
func Run() error {
	return defaultFactory.Run()
}

// 默认工厂停工
func Stop() {
	defaultFactory.Stop()
}
