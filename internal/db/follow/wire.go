package follow

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal/app/user"
)

var Set = wire.NewSet(
	NewStorage,

	wire.Bind(new(user.FollowStorage), new(*Storage)),
)
