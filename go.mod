module github.com/berachain/polaris

go 1.20

replace pkg.berachain.dev/polaris/build => ./build

require pkg.berachain.dev/polaris/build v0.0.0-00010101000000-000000000000

require (
	github.com/TwiN/go-color v1.4.0 // indirect
	github.com/magefile/mage v1.14.0 // indirect
)
