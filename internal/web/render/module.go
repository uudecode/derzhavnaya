package render

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewEngine))
