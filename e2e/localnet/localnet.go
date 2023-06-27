package localnet

import (
	"context"
	"fmt"
	"os"

	dt "github.com/ory/dockertest"
	dc "github.com/ory/dockertest/docker"
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

	pool      *dt.Pool
	container *dc.Container
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

	container, err := pool.Client.CreateContainer(dc.CreateContainerOptions{
		Name: "localnet",
		Config: &dc.Config{
			Image: imageName,
			ExposedPorts: map[dc.Port]struct{}{
				dc.Port(httpAddress): {},
				dc.Port(wsAddress):   {},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &LocalnetClient{
		genesis:     genesis,
		pool:        pool,
		container:   container,
		httpAddress: httpAddress,
		wsAddress:   wsAddress,
	}, nil
}

func (c *LocalnetClient) Start(ctx context.Context) error {
	return c.pool.Client.StartContainer(c.container.ID, nil)
}

func (c *LocalnetClient) Stop() error {
	return c.pool.Client.StopContainer(c.container.ID, 0)
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

//

const (
	baseImageName  = "polard/base:v0.0.0"
	baseContext    = "../../"
	baseDockerfile = "./cosmos/docker/base.Dockerfile"

	localnetImageName  = "polard/localnet:v0.0.0"
	localnetContext    = "./"
	localnetDockerfile = "Dockerfile"
)

func buildImage(pool *dt.Pool, buildArgs []dc.BuildArg, image, context, dockerfile string) error {
	baseBuildOpts := dc.BuildImageOptions{
		Name:         image,
		ContextDir:   context,
		Dockerfile:   dockerfile,
		BuildArgs:    buildArgs,
		OutputStream: os.Stdout,
	}

	return pool.Client.BuildImage(baseBuildOpts)
}

func buildDefault(pool *dt.Pool) error {
	baseBuildArgs := []dc.BuildArg{
		{
			Name:  "GO_VERSION",
			Value: "1.20.4",
		},
		{
			Name:  "PRECOMPILE_CONTRACTS_DIR",
			Value: "./contracts",
		},
		{
			Name:  "GOOS",
			Value: "linux",
		},
		{
			Name:  "GOARCH",
			Value: "arm64",
		},
	}
	if err := buildImage(pool, baseBuildArgs, baseImageName, baseContext, baseDockerfile); err != nil {
		return err
	}

	localnetBuildArgs := []dc.BuildArg{
		{
			Name:  "BASE_IMAGE",
			Value: baseImageName,
		},
	}
	return buildImage(pool, localnetBuildArgs, localnetImageName, localnetContext, localnetDockerfile)
}
