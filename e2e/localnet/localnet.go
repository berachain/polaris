package localnet

import (
	"fmt"
)

type Localnet interface {
	Build() error
	Start() error
	Stop() error
	Reset() error
	SetGenesis(string) error
	GetGenesis() string
	GetHTTPAddress() string
	GetWSAddress() string
}

type dockerizedNetwork struct {
	Container
	genesis     string
	httpAddress string
	wsAddress   string

	imageConfig ImageBuildConfig
}

// NewDockerizedNetwork creates a new localnet client.
func NewDockerizedNetwork(name, imageName, genesis, httpAddress, wsAddress string) (*dockerizedNetwork, error) {
	// Check for a genesis file.
	if genesis == "" {
		return nil, fmt.Errorf("genesis cannot be empty")
	}

	// Create the container config using the given input args.
	config := ContainerConfig{
		Name:        name,
		ImageName:   imageName,
		HTTPAddress: httpAddress,
		WSAddress:   wsAddress,
	}

	// Create the container client.
	imageConfig := ImageBuildConfig{
		ImageName:  localnetImageName,
		Context:    localnetContext,
		Dockerfile: localnetDockerfile,
		BuildArgs: map[string]string{
			"BASE_IMAGE": baseImageName,
		},
	}
	container, err := NewDefaultContainerClient(config, imageConfig)
	if err != nil {
		return nil, err
	}

	return &dockerizedNetwork{
		genesis:     genesis,
		httpAddress: httpAddress,
		wsAddress:   wsAddress,
		imageConfig: imageConfig,
		Container:   container,
	}, nil
}

func (c *dockerizedNetwork) Build() error {
	return c.Container.Build(c.imageConfig)
}

func (c *dockerizedNetwork) Reset() error {
	if err := c.Stop(); err != nil {
		return err
	}
	if err := c.Build(); err != nil {
		return err
	}
	return c.Start()
}

func (c *dockerizedNetwork) SetGenesis(genesis string) error {
	// override a config file/set one
	return nil
}

func (c *dockerizedNetwork) GetGenesis() string {
	return c.genesis
}

func (c *dockerizedNetwork) GetHTTPAddress() string {
	return c.httpAddress
}

func (c *dockerizedNetwork) GetWSAddress() string {
	return c.wsAddress
}
