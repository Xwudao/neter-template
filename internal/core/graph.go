package core

import (
	"github.com/Xwudao/loom"

	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
)

//go:generate go tool loom generate .

var testAppGraph = loom.Graph[*Test](
	loom.Name("TestApp"),
	loom.Provide(NewTestApp),
	loom.Provide(system.NewTestAppContext),
	loom.Provide(config.NewTestConfig),
	loom.Provide(logger.NewTestLogger),
)
