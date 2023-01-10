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
package mage

import (
	"strings"

	"github.com/magefile/mage/sh"
)

var allPkgs, _ = sh.Output("go", "list", "./...")

// RunCmd is a helper function that returns a function that runs the given
// command with the given arguments.
func RunCmdV(cmd string, args ...string) func(args ...string) error {
	return func(args2 ...string) error {
		return sh.RunV(cmd, append(args, args2...)...)
	}
}

// RunOutput is a helper function that returns a function that runs the given
// command with the given arguments and returns the output.
func RunOutput(cmd string, args ...string) func(args ...string) (string, error) {
	return func(args2 ...string) (string, error) {
		return sh.Output(cmd, append(args, args2...)...)
	}
}

// GoListFilter returns a list of packages that match the given filter.
func GoListFilter(include bool, contains ...string) []string {
	return filter(strings.Split(allPkgs, "\n"), func(s string) bool {
		for _, c := range contains {
			if strings.Contains(s, c) {
				return include
			}
		}
		return !include
	})
}

// filter returns a new slice containing only the elements of ss that
// satisfy the predicate test.
func filter[T any](ss []T, test func(T) bool) []T {
	ret := make([]T, 0, len(ss))
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}
