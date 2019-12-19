package wrappers

import (
	"github.com/bishopfox/sliver/sliver/scripting/lua"
)

var exports = map[string]lua.LGFunction{
	// Process utils
	// sliver.ListProcesses() []Process, err
	"ListProcesses": listProcesses,
	// sliver.FindProcess(pid) Process, err
	"FindProcess": findProcess,

	// Tasks
	// sliver.LocalTask(shellcode) err
	"LocalTask": luaLocalTask,
	// sliver.LocalTask(shellcode, pid) err
	"RemoteTask": luaRemoteTask,

	// StartProcess(path, args) err
	// "StartProcess": luaStartProcess,
}

// Loader sets up the lua state
func Loader(l *lua.LState) int {
	// register functions to the table
	mod := l.SetFuncs(l.NewTable(), exports)
	// returns the module
	l.Push(mod)
	return 1
}
