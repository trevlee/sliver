package wrappers

import (
	"github.com/bishopfox/sliver/sliver/scripting/lua"
	"io/ioutil"
	"net/http"
)

func luaHTTPGet(l *lua.LState) int {
	url := l.CheckString(1)
	resp, err := http.Get(url)
	if err != nil {
		l.Push(lua.LString(""))
		l.Push(lua.LString(err.Error()))
		return 2
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Push(lua.LString(""))
		l.Push(lua.LString(err.Error()))
		return 2
	}
	l.Push(lua.LString(string(body)))
	l.Push(lua.LString(""))
	return 2
}
