package localnet

import (
	"context"
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
	dt "github.com/ory/dockertest"
)

type Localnet interface {
	Start(context.Context) error
	Stop() error
	Reset(context.Context) error
	SetGenesis(string) error
	GetGenesis() string
	GetHTTPAddress() string
	GetWSAddress() string
}

type LocalnetClient struct {
	genesis     string
	httpAddress string
	wsAddress   string

	container dt.Pool
}

func NewLocalnetClient(ctx context.Context, imageName, genesis, httpAddress, wsAddress string) (*LocalnetClient, error) {
	if genesis == "" {
		return nil, fmt.Errorf("genesis cannot be empty")
	}

	pool, err := dt.NewPool("")

	err = pool.Client.Ping()
	if err != nil {
		return nil, err
	}

	container = pool.Client.CreateContainer(docker.CreateContainerOptions{
		Name: "localnet",
		Config: &docker.Config{
			Image: imageName,
			ExposedPorts: map[docker.Port]struct{}{
				"8545/tcp": {},
				"8546/tcp": {},
			},
		},
	})

	// req := tc.GenericContainerRequest{
	// 	ContainerRequest: tc.ContainerRequest{
	// 		Image:        imageName,
	// 		ExposedPorts: []string{"8545/tcp", "8546/tcp"},
	// 	},
	// }

	// container, err := tc.GenericContainer(ctx, req)
	// if err != nil {
	// 	return nil, err
	// }

	return &LocalnetClient{
		genesis:     genesis,
		container:   container,
		httpAddress: httpAddress,
		wsAddress:   wsAddress,
	}, nil
}

func (c *LocalnetClient) Start(ctx context.Context) error {
	return c.container.Start(ctx)
}

func (c *LocalnetClient) Stop() error {
	return c.container.Terminate(context.Background())
}

func (c *LocalnetClient) Reset(ctx context.Context) error {
	return nil
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
