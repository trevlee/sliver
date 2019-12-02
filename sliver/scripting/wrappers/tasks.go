package wrappers

import (
	"encoding/hex"
	"github.com/bishopfox/sliver/sliver/scripting/lua"
	"github.com/bishopfox/sliver/sliver/taskrunner"
)

func luaLocalTask(l *lua.LState) int {
	data := l.CheckString(1)
	raw := make([]byte, hex.DecodedLen(len(data)))
	_, err := hex.Decode(raw, []byte(data))
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}
	err = taskrunner.LocalTask(raw, false)
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}
	l.Push(lua.LString(""))
	return 1
}

func luaRemoteTask(l *lua.LState) int {
	data := l.CheckString(1)
	pid := l.CheckInt(2)
	raw := make([]byte, hex.DecodedLen(len(data)))
	_, err := hex.Decode(raw, []byte(data))
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}
	err = taskrunner.RemoteTask(pid, raw, false)
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}
	l.Push(lua.LString(""))
	return 1

}
