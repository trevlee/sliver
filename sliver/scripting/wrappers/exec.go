package wrappers

import (
	"os/exec"

	"github.com/bishopfox/sliver/sliver/scripting/lua"
)

func luaStartProcess(l *lua.LState) int {
	path := l.CheckString(1)
	args := l.CheckTable(2)
	cmd := exec.Command(path, args.String())
	err := cmd.Run()
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}
	l.Push(lua.LString(""))
	return 1
}
