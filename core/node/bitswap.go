package node

import (
	"context"

	blockstore "github.com/ipfs/go-ipfs-blockstore"
	exchange "github.com/ipfs/go-ipfs-exchange-interface"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/uss2022sayahi/go-bitswap"
	"github.com/uss2022sayahi/go-bitswap/network"
	config "github.com/uss2022sayahi/kubo/config"
	irouting "github.com/uss2022sayahi/kubo/routing"
	"go.uber.org/fx"

	"github.com/uss2022sayahi/kubo/core/node/helpers"
)

const (
	// Docs: https://github.com/uss2022sayahi/kubo/blob/master/docs/config.md#internalbitswap
	DefaultEngineBlockstoreWorkerCount = 128
	DefaultTaskWorkerCount             = 8
	DefaultEngineTaskWorkerCount       = 8
	DefaultMaxOutstandingBytesPerPeer  = 1 << 20
)

// OnlineExchange creates new LibP2P backed block exchange (BitSwap)
func OnlineExchange(cfg *config.Config, provide bool) interface{} {
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host, rt irouting.TieredRouter, bs blockstore.GCBlockstore) exchange.Interface {
		bitswapNetwork := network.NewFromIpfsHost(host, rt)

		var internalBsCfg config.InternalBitswap
		if cfg.Internal.Bitswap != nil {
			internalBsCfg = *cfg.Internal.Bitswap
		}

		opts := []bitswap.Option{
			bitswap.ProvideEnabled(provide),
			bitswap.EngineBlockstoreWorkerCount(int(internalBsCfg.EngineBlockstoreWorkerCount.WithDefault(DefaultEngineBlockstoreWorkerCount))),
			bitswap.TaskWorkerCount(int(internalBsCfg.TaskWorkerCount.WithDefault(DefaultTaskWorkerCount))),
			bitswap.EngineTaskWorkerCount(int(internalBsCfg.EngineTaskWorkerCount.WithDefault(DefaultEngineTaskWorkerCount))),
			bitswap.MaxOutstandingBytesPerPeer(int(internalBsCfg.MaxOutstandingBytesPerPeer.WithDefault(DefaultMaxOutstandingBytesPerPeer))),
		}
		exch := bitswap.New(helpers.LifecycleCtx(mctx, lc), bitswapNetwork, bs, opts...)
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return exch.Close()
			},
		})
		return exch

	}
}
