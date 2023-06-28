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

type LocalnetClient struct {
	genesis     string
	httpAddress string
	wsAddress   string

	imageConfig ImageBuildConfig
	container   Container
}

// NewLocalnetClient creates a new localnet client.
func NewLocalnetClient(name, imageName, genesis, httpAddress, wsAddress string) (*LocalnetClient, error) {
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

	return &LocalnetClient{
		genesis:     genesis,
		httpAddress: httpAddress,
		wsAddress:   wsAddress,
		imageConfig: imageConfig,
		container:   container,
	}, nil
}

func (c *LocalnetClient) Build() error {
	return c.container.Build(c.imageConfig)
}

func (c *LocalnetClient) Reset() error {
	if err := c.container.Stop(); err != nil {
		return err
	}
	if err := c.container.Build(c.imageConfig); err != nil {
		return err
	}
	return c.container.Start()
}

func (c *LocalnetClient) SetGenesis(genesis string) error {
	// override a config file/set one
	return nil
}

func (c *LocalnetClient) GetGenesis() string {
	return c.genesis
}

func (c *LocalnetClient) GetHTTPAddress() string {
	return c.httpAddress
}

func (c *LocalnetClient) GetWSAddress() string {
	return c.wsAddress
}
