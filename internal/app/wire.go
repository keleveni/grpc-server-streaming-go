package app

import (
	"github.com/google/wire"
	"github.com/kodaikumatani/grpc-golang/internal/app/notification"
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	"github.com/kodaikumatani/grpc-golang/internal/app/user"
)

var Set = wire.NewSet(
	recipe.NewCreateUseCase,
	recipe.NewReadUseCase,
	notification.NewSubscribeFeedUseCase,
	user.NewFollowUseCase,
)
