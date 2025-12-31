package orm

import (
	"github.com/google/wire"
	"github.com/uptrace/bun"
)

var Set = wire.NewSet(
	NewBun,
	wire.Bind(new(bun.IDB), new(*bun.DB)),
)
