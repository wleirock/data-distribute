package api

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
)

type lStatePool struct {
	m     sync.Mutex
	saved []*lua.LState
}

// Get 从池中取出一个LState，若池中没有，则新增一个
func (pl *lStatePool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.saved)
	if n == 0 {
		return pl.New()
	}
	x := pl.saved[n-1]
	pl.saved = pl.saved[0 : n-1]
	return x
}

func (pl *lStatePool) New() *lua.LState {
	L := lua.NewState()
	return L
}

// Put 用完LState后放回池中
func (pl *lStatePool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.saved = append(pl.saved, L)
}

func (pl *lStatePool) ShutDown() {
	for _, L := range pl.saved {
		L.Close()
	}
}

//初始化连接池
var luaPool = &lStatePool{
	saved: make([]*lua.LState, 0, 5),
}
