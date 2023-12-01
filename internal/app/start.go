package app

import "github.com/smartfood-capstone/backend/internal/server"

type StartCmd struct {
	Server server.Server
}

func NewStartCmd() *StartCmd {
	svr := server.New()

	return &StartCmd{
		Server: svr,
	}
}
