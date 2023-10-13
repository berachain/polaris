package runtime

import (
	"cosmossdk.io/log"
	ethlog "pkg.berachain.dev/polaris/eth/log"
)

// LoggerFuncHandler injects the cosmos-sdk logger into geth.
func LoggerFuncHandler(logger log.Logger) ethlog.Handler {
	return ethlog.FuncHandler(func(r *ethlog.Record) error {
		polarisGethLogger := logger.With("module", "polaris-geth")
		switch r.Lvl { //nolint:nolintlint,exhaustive // linter is bugged.
		case ethlog.LvlTrace:
		case ethlog.LvlDebug:
			polarisGethLogger.Debug(r.Msg, r.Ctx...)
		case ethlog.LvlInfo:
			polarisGethLogger.Info(r.Msg, r.Ctx...)
		case ethlog.LvlWarn:
		case ethlog.LvlCrit:
		case ethlog.LvlError:
			polarisGethLogger.Error(r.Msg, r.Ctx...)
		}
		return nil
	})
}
