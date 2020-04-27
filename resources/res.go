//这个包是为了管理全局资源的，比如数据库，kafka等从程序启动开始一直到程序退出的全局资源。
package resources

// Res 接口定义了，资源需要实现的接口
type Res interface {
	Free()
	Init()
}

// resMgr 统一存储资源的
type resMgr struct {
	//非线程安全
	reses []Res
}

var resMgrInst resMgr

// Enroll 资源注册
func Enroll(r Res) {
	resMgrInst.reses = append(resMgrInst.reses, r)
}

// Init 初始化已注册的资源
func Init() {
	for _, v := range resMgrInst.reses {
		v.Init()
	}
}

// Free 释放所有注册的资源，注册顺序的倒序
func Free() {
	if len(resMgrInst.reses) < 1 {
		return
	}

	for i := len(resMgrInst.reses) - 1; i >= 0; i-- {
		resMgrInst.reses[i].Free()
	}
}
