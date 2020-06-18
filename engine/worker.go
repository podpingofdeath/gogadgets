// 工人，有不同工种的工人，不同工种的工人处理不同流水线上的零件，工人下班要打卡
package engine

import "sync"

// Worker 工人接口，工人不断地从指定流水线上拿零件来处理
type Worker interface {
	Working(c <-chan interface{}, group *sync.WaitGroup)
}

type FuncWorker func(c <-chan interface{}, group *sync.WaitGroup)

func (f FuncWorker) Working(c <-chan interface{}, group *sync.WaitGroup) {
	f(c, group)
}
