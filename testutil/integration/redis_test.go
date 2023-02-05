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

package integration

// const (
// 	defaultHttpPortTcp = "8545/tcp"
// 	defaultWsPortTcp   = "8546/tcp"
// )

// func TestWithRedis(t *testing.T) {
// 	ctx := context.Background()
// 	os.Setenv("DOCKER_BUILDKIT", "1")
// 	req := testcontainers.ContainerRequest{
// 		FromDockerfile: testcontainers.FromDockerfile{
// 			Context:    ".",
// 			Dockerfile: "Dockerfile",
// 		},
// 		// ExposedPorts: []string{defaultHttpPortTcp, defaultWsPortTcp},
// 		// WaitingFor:   wait.ForHTTP("/").WithPort(defaultHttpPortTcp),
// 	}
// 	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
// 		ContainerRequest: req,
// 		Started:          true,
// 	})
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
