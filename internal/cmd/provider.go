package cmd

import "github.com/google/wire"

var ProvideCmdSet = wire.NewSet(
	newMigrateCmd,
	newUpCmd,
	newDownCmd,
	newInitCmd,
	newConfigCmd,
)
