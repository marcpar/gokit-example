module github.com/marcpar/HextoRGB/hex_to_rgb

require (
	github.com/marcpar/HextoRGB/hex_to_rgb/cmd v0.0.0 // indirect
	github.com/marcpar/HextoRGB/hex_to_rgb/cmd/service v0.0.0 // indirect
	github.com/marcpar/HextoRGB/hex_to_rgb/pkg/endpoint v0.0.0 // indirect
	github.com/marcpar/HextoRGB/hex_to_rgb/pkg/http v0.0.0 // indirect
	github.com/marcpar/HextoRGB/hex_to_rgb/pkg/service v0.0.0 // indirect
)

replace (
	github.com/marcpar/HextoRGB/hex_to_rgb/cmd => ./cmd
	github.com/marcpar/HextoRGB/hex_to_rgb/cmd/service => ./cmd/service
	github.com/marcpar/HextoRGB/hex_to_rgb/pkg/endpoint => ./pkg/endpoint
	github.com/marcpar/HextoRGB/hex_to_rgb/pkg/http => ./pkg/http
	github.com/marcpar/HextoRGB/hex_to_rgb/pkg/service => ./pkg/service
)
