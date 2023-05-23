package e2e

import (
	"context"

	tc "github.com/testcontainers/testcontainers-go"
	"pkg.berachain.dev/polaris/testing"
)

var _ testing.EthereumLikeChain = &PolarisChain{}

type PolarisChain struct {
	container *testing.Container
}

func (pc *PolarisChain) Start(ctx context.Context) error {
	// Implement your method here
	return pc.container.Start(ctx)
}

func (pc *PolarisChain) Stop() error {
	// Implement your method here
	return pc.container.Stop()
}

func (pc *PolarisChain) Reset(ctx context.Context) error {
	// Implement your method here
	return nil
}

func (pc *PolarisChain) SetGenesis(genesisPath string) {
	// Implement your method here

}

func (pc *PolarisChain) GetGenesis() string {
	return ""
}

func NewPolarisChainWithGenesis(genesisPath string) (*PolarisChain, error) {
	return NewPolarisChainFromDockerfile("../../..", "./cosmos/docker/local/Dockerfile", nil)
}

// NewEthChainFromDockerfile creates a new container from the provided configuration.
func NewPolarisChainFromDockerfile(dockerContext, dockerfilePath string, buildArgs map[string]*string) (*PolarisChain, error) {
	if buildArgs == nil {
		buildArgs = make(map[string]*string)
	}
	dockerReq := tc.FromDockerfile{
		Context:       dockerContext,
		Dockerfile:    dockerfilePath,
		BuildArgs:     buildArgs,
		PrintBuildLog: true,
	}

	container, err := testing.NewContainer(context.Background(), dockerReq)
	if err != nil {
		return nil, err
	}

	return &PolarisChain{container: container}, nil
}
