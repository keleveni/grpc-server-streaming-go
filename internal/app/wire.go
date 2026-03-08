package app

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/app/notification"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/app/recipe"
	"github.com/kodaikumatani/grpc-server-streaming-go/internal/app/user"
)

var Set = wire.NewSet(
	recipe.NewCreateUseCase,
	recipe.NewReadUseCase,
	notification.NewSubscribeFeedUseCase,
	user.NewFollowUseCase,
)
