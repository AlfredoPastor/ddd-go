//go:build wireinject
// +build wireinject

package internal

import "github.com/google/wire"

func InitializeServer() (Server, error) {
	wire.Build(superSet)
	return Server{}, nil
}
