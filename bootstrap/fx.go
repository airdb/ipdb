package bootstrap

import (
	"ipdb/modules/ip2locationmod"
	"ipdb/modules/ipipmod"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(ip2locationmod.NewIPdb),
		fx.Provide(ipipmod.NewDB),
		fx.Provide(NewServe),
	)
}
