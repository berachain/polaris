// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//go:build mage

package main

import (
	//mage:import
	"fmt"

	_ "github.com/berachain/stargazer/build/mage"

	"context"
	"time"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultHttpPortTcp = "8545/tcp"
	defaultWsPortTcp   = "8546/tcp"
)

var (
	goVersion = "1.19.5"
	goAlpine  = "golang:1.19.5-alpine3.14"
)

func main() {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "./",
			Dockerfile: "jsonrpc/Dockerfile",
			BuildArgs: map[string]*string{
				"GO_VERSION":   &goVersion,
				"RUNNER_IMAGE": &goAlpine,
			},
			PrintBuildLog: true,
		},
		// ExposedPorts: []string{defaultHttpPortTcp, defaultWsPortTcp},
		// WaitingFor:   wait.ForHTTP("/").WithPort(defaultHttpPortTcp),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 15)
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			panic(fmt.Sprintf("failed to terminate container: %s", err.Error()))
		}
	}()
}

// func TestWithRedis(t *testing.T) {

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	time.Sleep(time.Second * 15)
// 	defer func() {
// 		if err := redisC.Terminate(ctx); err != nil {
// 			t.Fatalf("failed to terminate container: %s", err.Error())
// 		}
// 	}()
// }
