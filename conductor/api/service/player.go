package service

import "conductor/common/model"

var (
	PlayerService = Service[model.PlayerResponse]{
		Key:    "player",
		Getter: nil,
		Lister: nil,
	}
)
