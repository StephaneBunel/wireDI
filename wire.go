//+build wireinject

package main

import "github.com/google/wire"

func initUserStore(uri string) (*UserStore, func(), error) {
	panic(wire.Build(ProvideUserStore, ProvideUserStoreDefaultConfig, ProvideDB))
}
