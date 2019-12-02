package wrappers

import (
	"github.com/bishopfox/sliver/sliver/ps"
	"github.com/bishopfox/sliver/sliver/scripting/lua"
	luar "github.com/bishopfox/sliver/sliver/scripting/utils/gopher-luar"
)

func listProcesses(l *lua.LState) int {
	table := &lua.LTable{}
	procs, err := ps.Processes()
	if err != nil {
		l.Push(table)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	for _, p := range procs {
		table.Append(luar.New(l, p))
	}
	l.Push(table)
	l.Push(lua.LString(""))
	return 2
}

func findProcess(l *lua.LState) int {
	pid := l.CheckInt(1)
	proc, err := ps.FindProcess(pid)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	l.Push(luar.New(l, proc))
	l.Push(lua.LString(""))
	return 2
}
