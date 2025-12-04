module github.com/ArtisanCloud/PowerWeChat/v3

go 1.25.0

replace github.com/ArtisanCloud/PowerLibs/v3 => ../PowerLibs

// replace github.com/ArtisanCloud/PowerSocialite/v3 => ../PowerSocialite

require (
	github.com/ArtisanCloud/PowerLibs/v3 v3.3.2
	github.com/ArtisanCloud/PowerSocialite/v3 v3.0.9
	github.com/go-playground/assert/v2 v2.0.1
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.17.2
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel v1.38.0 // indirect
	go.opentelemetry.io/otel/trace v1.38.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.1 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
