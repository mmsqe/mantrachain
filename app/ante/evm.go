package ante

import (
	sanctionkeeper "github.com/MANTRA-Chain/mantrachain/v5/x/sanction/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmante "github.com/cosmos/evm/ante"
	cosmosevmante "github.com/cosmos/evm/ante/evm"
)

func newEVMAnteHandler(ctx sdk.Context, options HandlerOptions) sdk.AnteHandler {
	evmParams := options.EvmOptions.EvmKeeper.GetParams(ctx)
	feemarketParams := options.EvmOptions.FeeMarketKeeper.GetParams(ctx)
	decorators := []sdk.AnteDecorator{
		sanctionkeeper.NewEVMBlacklistCheckDecorator(*options.SanctionKeeper),
		cosmosevmante.NewEVMMonoDecorator(
			options.EvmOptions.AccountKeeper,
			options.EvmOptions.FeeMarketKeeper,
			options.EvmOptions.EvmKeeper,
			options.EvmOptions.MaxTxGasWanted,
			&evmParams,
			&feemarketParams,
		),
	}
	if options.EvmOptions.PendingTxListener != nil {
		decorators = append(decorators, evmante.NewTxListenerDecorator(options.EvmOptions.PendingTxListener))
	}
	return sdk.ChainAnteDecorators(decorators...)
}
