package bootstrap

import (
	"ipdb/modules/ip2locationmod"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(ip2locationmod.NewIPdb),
	)
}
