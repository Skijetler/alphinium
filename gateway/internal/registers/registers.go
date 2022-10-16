package registers

import "context"

var ProviderSet = wire.NewSet(RegisterAll)

func RegisterAll() (*runtime.ServeMux, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())

	cleanup := func() {
		cancel()
	}

	mux := runtime.NewServeMux()

	// register services

	return mux, cleanup, nil
}
