// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package state_test

import (
	"math/big"

	"cosmossdk.io/log"

	testutil "github.com/berachain/polaris/cosmos/testutil"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	alice         = testutil.Alice
	bob           = testutil.Bob
	emptyCodeHash = crypto.Keccak256Hash(nil)
)

var _ = Describe("State Plugin", func() {
	var ak state.AccountKeeper
	var ctx sdk.Context
	var sp state.Plugin

	BeforeEach(func() {
		ctx, ak, _, _ = testutil.SetupMinimalKeepers(log.NewTestLogger(GinkgoT()))
		sp = state.NewPlugin(ak, testutil.EvmKey, nil, &mockPLF{})
		sp.Reset(ctx)
	})

	It("should have the correct registry key", func() {
		Expect(sp.RegistryKey()).To(Equal("statePlugin"))
	})

	Describe("TestevmReset", func() {
		It("should reset", func() {
			sp.CreateAccount(alice)
			sp.AddBalance(alice, big.NewInt(50))
			sp.SetCode(alice, []byte{1, 2, 3})
			sp.SetState(alice, common.BytesToHash([]byte{1}), common.BytesToHash([]byte{2}))

			sp.Reset(testutil.NewContext(log.NewTestLogger(GinkgoT())))

			Expect(sp.Exist(alice)).To(BeFalse())
			Expect(sp.GetBalance(alice)).To(Equal(new(big.Int)))
			Expect(sp.GetCode(alice)).To(BeEmpty())
			Expect(sp.GetState(alice, common.BytesToHash([]byte{1}))).To(Equal(common.Hash{}))
		})
	})

	Describe("TestCreateAccount", func() {
		It("should create account", func() {
			sp.CreateAccount(alice)
			Expect(sp.Exist(alice)).To(BeTrue())
		})

		It("should handle empty", func() {
			sp.CreateAccount(alice)
			Expect(sp.Empty(alice)).To(BeTrue())

			sp.SetCode(alice, []byte{1, 2, 3})
			Expect(sp.Empty(alice)).To(BeFalse())
		})
	})

	Describe("TestBalance", func() {
		It("should have start with zero balance", func() {
			Expect(sp.GetBalance(alice)).To(Equal(new(big.Int)))
		})

		Context("TestAddBalance", func() {
			It("should be able to add zero", func() {
				Expect(sp.GetBalance(alice)).To(Equal(new(big.Int)))
				sp.AddBalance(alice, new(big.Int))
				Expect(sp.GetBalance(alice)).To(Equal(new(big.Int)))
			})
			It("should have 100 balance", func() {
				sp.AddBalance(alice, big.NewInt(100))
				Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(100)))
			})
		})

		It("should handle complex balance updates", func() {
			// Initial balance for alice should be 0
			Expect(sp.GetBalance(alice)).To(Equal(new(big.Int)))

			// Add some balance to alice
			sp.AddBalance(alice, big.NewInt(100))
			Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(100)))

			// Subtract some balance from alice
			sp.SubBalance(alice, big.NewInt(50))
			Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(50)))

			// Add some balance to alice
			sp.AddBalance(alice, big.NewInt(100))
			Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(150)))
		})
	})

	Describe("TestNonce", func() {
		When("account exists", func() {
			BeforeEach(func() {
				sp.CreateAccount(alice)
			})
			It("should have start with zero nonce", func() {
				Expect(sp.GetNonce(alice)).To(Equal(uint64(0)))
			})
			It("should have 100 nonce", func() {
				sp.SetNonce(alice, 100)
				Expect(sp.GetNonce(alice)).To(Equal(uint64(100)))
			})
		})
		When("account does not exist", func() {
			It("should have start with zero nonce", func() {
				Expect(sp.GetNonce(alice)).To(Equal(uint64(0)))
			})

			It("should have 100 nonce", func() {
				sp.SetNonce(alice, 100)
				Expect(sp.GetNonce(alice)).To(Equal(uint64(100)))
			})
		})
	})

	Describe("TestCode & CodeHash", func() {
		When("account does not exist", func() {
			It("should have empty code hash", func() {
				Expect(sp.GetCodeHash(alice)).To(Equal(common.Hash{}))
			})
			It("should not have code", func() { // ensure account exists
				Expect(sp.GetCode(alice)).To(BeNil())
				Expect(sp.GetCodeHash(alice)).To(Equal(common.Hash{}))
			})
			It("cannot set code", func() { // ensure account exists
				sp.SetCode(alice, []byte("code"))
				Expect(sp.GetCode(alice)).To(BeNil())
				Expect(sp.GetCodeHash(alice)).To(Equal(common.Hash{}))
			})
		})
		When("account exists", func() {
			BeforeEach(func() {
				sp.CreateAccount(alice)
			})

			It("should have empty code hash", func() {
				Expect(sp.GetCodeHash(alice)).To(Equal(emptyCodeHash))
			})

			It("should return empty code hash when account exists but no codehash", func() {
				addr := ak.NewAccountWithAddress(ctx, bob[:])
				ak.SetAccount(ctx, addr)

				Expect(sp.GetCodeHash(bob)).To(Equal(emptyCodeHash))
			})

			When("account has code", func() {
				BeforeEach(func() {
					sp.SetCode(alice, []byte("code"))
				})
				It("should have code", func() {
					Expect(sp.GetCode(alice)).To(Equal([]byte("code")))
					Expect(sp.GetCodeHash(alice)).To(Equal(crypto.Keccak256Hash([]byte("code"))))
				})
				It("should have empty code hash", func() {
					sp.SetCode(alice, nil)
					Expect(sp.GetCode(alice)).To(BeNil())
					Expect(sp.GetCodeHash(alice)).To(Equal(emptyCodeHash))
				})
			})
		})
	})

	Describe("TestState", func() {
		It("should have empty state", func() {
			Expect(sp.GetState(alice, common.Hash{3})).To(Equal(common.Hash{}))
		})
		When("set basic state", func() {
			BeforeEach(func() {
				sp.SetState(alice, common.Hash{3}, common.Hash{1})
			})

			It("should have state", func() {
				Expect(sp.GetState(alice, common.Hash{3})).To(Equal(common.Hash{1}))
			})
			It("should have state changed", func() {
				sp.SetState(alice, common.Hash{3}, common.Hash{2})
				Expect(sp.GetState(alice, common.Hash{3})).To(Equal(common.Hash{2}))
				Expect(sp.GetCommittedState(alice, common.Hash{3})).To(Equal(common.Hash{}))
			})

			When("state is committed", func() {
				BeforeEach(func() {
					sp.Finalize()
					It("should have committed state", func() {
						Expect(sp.GetCommittedState(alice, common.Hash{3})).To(Equal(common.Hash{1}))
					})
					It("should maintain committed state", func() {
						sp.SetState(alice, common.Hash{3}, common.Hash{4})
						Expect(sp.GetCommittedState(alice, common.Hash{3})).
							To(Equal(common.Hash{1}))
						Expect(sp.GetState(alice, common.Hash{3})).To(Equal(common.Hash{4}))
					})
				})
			})
		})

		Describe("TestExist", func() {
			It("should not exist", func() {
				Expect(sp.Exist(alice)).To(BeFalse())
			})
			When("account is created", func() {
				BeforeEach(func() {
					sp.CreateAccount(alice)
				})
				It("should exist", func() {
					Expect(sp.Exist(alice)).To(BeTrue())
				})
			})
		})

		Describe("Test ForEachStorage", func() {
			BeforeEach(func() {
				sp.CreateAccount(alice)
				sp.CreateAccount(bob)
			})

			type Slot struct {
				Key   common.Hash
				Value common.Hash
			}

			type Storage []Slot

			It("should iterate through storage correctly", func() {
				Expect(sp.GetCode(alice)).To(BeNil())
				var aliceStorage Storage
				err := sp.ForEachStorage(alice,
					func(key, value common.Hash) bool {
						aliceStorage = append(aliceStorage,
							Slot{key, value})
						return true
					})
				Expect(err).ToNot(HaveOccurred())
				Expect(aliceStorage).To(BeEmpty())

				sp.SetState(bob, common.BytesToHash([]byte{1}), common.BytesToHash([]byte{2}))
				var bobStorage Storage
				err = sp.ForEachStorage(bob,
					func(key, value common.Hash) bool {
						bobStorage = append(bobStorage, Slot{key, value})
						return true
					})
				Expect(err).ToNot(HaveOccurred())
				Expect(bobStorage).To(HaveLen(1))
				Expect(bobStorage[0].Key).
					To(Equal(common.HexToHash(
						"0x0000000000000000000000000000000000000000000000000000000000000001")))
				Expect(bobStorage[0].Value).
					To(Equal(common.HexToHash(
						"0x0000000000000000000000000000000000000000000000000000000000000002")))

				sp.SetState(bob, common.BytesToHash([]byte{3}), common.BytesToHash([]byte{4}))
				var bobStorage2 Storage
				i := 0
				err = sp.ForEachStorage(bob,
					func(key, value common.Hash) bool {
						if i > 0 {
							return false
						}

						bobStorage2 = append(bobStorage2, Slot{key, value})
						i++
						return true
					},
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(bobStorage2).To(HaveLen(1))
			})
		})

		Describe("Test Delete Suicides", func() {
			aliceCode := []byte("alicecode")

			BeforeEach(func() {
				sp.CreateAccount(alice)
				sp.SetCode(alice, aliceCode)
				sp.SetState(alice, common.BytesToHash([]byte{1}), common.BytesToHash([]byte{2}))
			})

			It("should remove storage/codehash/acct", func() {
				sp.DeleteAccounts([]common.Address{alice, alice})
				Expect(ak.HasAccount(ctx, alice[:])).To(BeFalse())
				Expect(sp.GetCode(alice)).To(BeNil())
				Expect(sp.GetState(alice, common.BytesToHash([]byte{1}))).To(Equal(common.Hash{}))
			})
		})

		Describe("TestAccount", func() {
			It("account does not exist", func() {
				Expect(sp.Exist(alice)).To(BeFalse())
				// Expect(sp.Empty(alice)).To(BeTrue())
				Expect(sp.GetBalance(alice)).To(Equal(new(big.Int)))
				Expect(sp.GetNonce(alice)).To(BeZero())
				Expect(sp.GetCodeHash(alice)).To(Equal(common.Hash{}))
				Expect(sp.GetCode(alice)).To(BeNil())
				Expect(sp.GetState(alice, common.Hash{})).To(Equal(common.Hash{}))
				// Expect(sp.GetRefund()).To(BeZero())
				Expect(sp.GetCommittedState(alice, common.Hash{})).To(Equal(common.Hash{}))
			})
			When("account exists", func() {
				BeforeEach(func() {
					sp.AddBalance(alice, big.NewInt(56))
					sp.SetNonce(alice, 59)
				})
				It("accidental override account", func() {
					// override
					sp.CreateAccount(alice)

					// check balance is not reset
					Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(56)))
				})
			})
		})

		Describe("TestSnapshot", func() {
			key := common.BytesToHash([]byte("key"))
			value1 := common.BytesToHash([]byte("value1"))
			value2 := common.BytesToHash([]byte("value2"))
			It("simple revert", func() {
				revision := sp.Snapshot()
				Expect(revision).To(Equal(0))
				sp.SetState(alice, key, value1)
				Expect(sp.GetState(alice, key)).To(Equal(value1))
				sp.RevertToSnapshot(revision)
				Expect(sp.GetState(alice, key)).To(Equal(common.Hash{}))
			})
			It("nested snapshot & revert", func() {
				revision1 := sp.Snapshot()
				Expect(revision1).To(Equal(0))
				sp.SetState(alice, key, value1)
				revision2 := sp.Snapshot()
				Expect(revision2).To(Equal(1))
				sp.SetState(alice, key, value2)
				Expect(sp.GetState(alice, key)).To(Equal(value2))

				sp.RevertToSnapshot(revision2)
				Expect(sp.GetState(alice, key)).To(Equal(value1))

				sp.RevertToSnapshot(revision1)
				Expect(sp.GetState(alice, key)).To(Equal(common.Hash{}))
			})
			It("jump revert", func() {
				revision1 := sp.Snapshot()
				Expect(revision1).To(Equal(0))
				sp.SetState(alice, key, value1)
				sp.Snapshot()
				sp.SetState(alice, key, value2)
				Expect(sp.GetState(alice, key)).To(Equal(value2))
				sp.RevertToSnapshot(revision1)
				Expect(sp.GetState(alice, key)).To(Equal(common.Hash{}))
			})
		})

		Describe("TestRevertSnapshot", func() {
			key := common.BytesToHash([]byte("key"))
			value := common.BytesToHash([]byte("value"))

			When("We make a bunch of arbitrary changes", func() {
				BeforeEach(func() {
					sp.SetNonce(alice, 1)
					sp.AddBalance(alice, big.NewInt(100))
					sp.SetCode(alice, []byte("hello world"))
					sp.SetState(alice, key, value)
					sp.SetNonce(bob, 1)
				})
				When("we take a snapshot", func() {
					var revision int
					BeforeEach(func() {
						revision = sp.Snapshot()
					})
					When("we do more changes", func() {
						AfterEach(func() {
							sp.RevertToSnapshot(revision)
							Expect(sp.GetNonce(alice)).To(Equal(uint64(1)))
							Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(100)))
							Expect(sp.GetCode(alice)).To(Equal([]byte("hello world")))
							Expect(sp.GetState(alice, key)).To(Equal(value))
							Expect(sp.GetNonce(bob)).To(Equal(uint64(1)))
						})

						It("if we change balance", func() {
							sp.AddBalance(alice, big.NewInt(100))
						})

						It("if we change nonce", func() {
							sp.SetNonce(alice, 2)
						})

						It("if we change code", func() {
							sp.SetCode(alice, []byte("goodbye world"))
						})

						It("if we change state", func() {
							sp.SetState(alice, key, common.Hash{})
						})

						It("if we change nonce of another account", func() {
							sp.SetNonce(bob, 2)
						})
					})

					When("we make a nested snapshot", func() {
						var revision2 int
						BeforeEach(func() {
							sp.SetState(alice, key, common.Hash{2})
							revision2 = sp.Snapshot()
						})
						When("we revert to snapshot ", (func() {
							It("revision 2", func() {
								sp.RevertToSnapshot(revision2)
								Expect(sp.GetState(alice, key)).To(Equal(common.Hash{2}))
							})
							It("revision 1", func() {
								sp.RevertToSnapshot(revision)
								Expect(sp.GetState(alice, key)).To(Equal(value))
							})
						}))
					})
				})
				When("we revert to an invalid snapshot", func() {
					It("should panic", func() {
						Expect(func() {
							sp.RevertToSnapshot(100)
						}).To(Panic())
					})
				})
			})
		})
	})
})

// MOCKS BELOW.

type mockPLF struct{}

func (mplf *mockPLF) Build(event *sdk.Event) (*ethtypes.Log, error) {
	return &ethtypes.Log{
		Address: common.BytesToAddress([]byte(event.Type)),
	}, nil
}
