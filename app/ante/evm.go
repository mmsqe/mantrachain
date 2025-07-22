package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmante "github.com/cosmos/evm/ante/evm"
	chainante "github.com/cosmos/evm/evmd/ante"
)

func newMonoEVMAnteHandler(options EVMHandlerOptions) sdk.AnteHandler {
	decorators := []sdk.AnteDecorator{
		evmante.NewEVMMonoDecorator(
			options.AccountKeeper,
			options.FeeMarketKeeper,
			options.EvmKeeper,
			options.MaxTxGasWanted,
		),
	}
	if options.PendingTxListener != nil {
		decorators = append(decorators, chainante.NewTxListenerDecorator(options.PendingTxListener))
	}
	return sdk.ChainAnteDecorators(decorators...)
}
