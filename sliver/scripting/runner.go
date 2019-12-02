package scripting

import (
	"github.com/bishopfox/sliver/sliver/scripting/lua"
	luar "github.com/bishopfox/sliver/sliver/scripting/utils/gopher-luar"
	"github.com/bishopfox/sliver/sliver/scripting/utils/logger"
	"github.com/bishopfox/sliver/sliver/scripting/wrappers"
)

// RunScript runs a lua script passed as a string
func RunScript(script string) (string, error) {
	state := lua.NewState()
	state.PreloadModule("sliver", wrappers.Loader)
	logger := &logger.ScriptLogger{}
	state.SetGlobal("log", luar.New(state, logger))
	if err := state.DoString(script); err != nil {
		return "", err
	}
	defer state.Close()
	return logger.String(), nil
}
