// wire.go

package internal

import "github.com/google/wire"

func InitializeServer() (Server, error) {
	wire.Build(superSet)
	return Server{}, nil
}
