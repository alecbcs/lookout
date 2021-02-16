package cli

import (
	"github.com/DataDrake/cli-ng/v2/cmd"
)

//GlobalFlags contains the flags for commands.
type GlobalFlags struct{}

// Root is the main command.
var Root *cmd.Root

// init creates the command interface and registers the possible commands.
func init() {
	Root = &cmd.Root{
		Name:  "lookout",
		Short: "Lookout is an Upstream Update Watcher",
		Flags: &GlobalFlags{},
	}
	cmd.Register(&cmd.Help)
	cmd.Register(&Add)
	cmd.Register(&AddDep)
	cmd.Register(&Import)
	cmd.Register(&Info)
	cmd.Register(&List)
	cmd.Register(&Remove)
	cmd.Register(&RemoveDep)
	cmd.Register(&Run)
	cmd.Register(&Upgrade)
}
