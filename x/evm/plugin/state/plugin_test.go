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

package state_test

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ethstate "github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/crypto"
	"github.com/berachain/stargazer/testutil"
	"github.com/berachain/stargazer/x/evm/plugin/state"
	"github.com/berachain/stargazer/x/evm/plugin/state/storage"
)

var (
	alice         = testutil.Alice
	bob           = testutil.Bob
	emptyCodeHash = crypto.Keccak256Hash(nil)
)

var _ = Describe("State Plugin", func() {
	var ak state.AccountKeeper
	var bk state.BankKeeper
	var ctx sdk.Context
	var sp ethstate.StatePlugin

	BeforeEach(func() {
		ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
		sp = state.NewPlugin(ctx, ak, bk, testutil.EvmKey, "abera") // TODO: use lf
	})

	It("should have the correct registry key", func() {
		Expect(sp.RegistryKey()).To(Equal("statePlugin"))
	})

	Describe("TestCreateAccount", func() {
		It("should create account", func() {
			sp.CreateAccount(alice)
			Expect(sp.Exist(alice)).To(BeTrue())
		})
	})

	Describe("TestBalance", func() {
		It("should have start with zero balance", func() {
			Expect(sp.GetBalance(alice)).To(Equal(new(big.Int)))
		})

		It("should correctly Transfer Balance", func() {
			sp.AddBalance(alice, big.NewInt(50))
			Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(50)))
			Expect(sp.GetBalance(bob)).To(Equal(big.NewInt(0)))

			sp.TransferBalance(alice, bob, big.NewInt(25))
			Expect(sp.GetBalance(alice)).To(Equal(big.NewInt(25)))
			Expect(sp.GetBalance(bob)).To(Equal(big.NewInt(25)))

			// should panic if not enough funds
			Expect(func() { sp.TransferBalance(alice, bob, big.NewInt(50)) }).To(Panic())
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
			It("should panic if using negative value", func() {
				Expect(func() {
					sp.AddBalance(alice, big.NewInt(-100))
				}).To(Panic())
			})
		})

		Context("TestSubBalance", func() {
			It("should not set balance to negative value", func() {
				Expect(func() {
					sp.SubBalance(alice, big.NewInt(100))
				}).To(Panic())
			})
			It("should panic if using negative value", func() {
				Expect(func() {
					sp.SubBalance(alice, big.NewInt(-100))
				}).To(Panic())
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

			// Subtract some balance from alice
			Expect(func() {
				sp.SubBalance(alice, big.NewInt(200))
			}).To(Panic())
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
				// When("suicided", func() {
				// 	BeforeEach(func() {
				// 		// Only contracts can be suicided
				// 		sp.SetCode(alice, []byte("code"))
				// 		Expect(sp.Suicide(alice)).To(BeTrue())
				// 	})
				// 	It("should still exist", func() {
				// 		Expect(sp.Exist(alice)).To(BeTrue())
				// 	})
				// 	When("commit", func() {
				// 		BeforeEach(func() {
				// 			Expect(sp.Finalize()).To(BeNil())
				// 		})
				// 		It("should not exist", func() {
				// 			Expect(sp.Exist(alice)).To(BeFalse())
				// 		})
				// 	})
				// })
			})
		})

		// Describe("TestReset", func() {
		// 	BeforeEach(func() {
		// 		sp.AddRefund(1000)
		// 		sp.AddLog(&coretypes.Log{})
		// 		sp.Prepare(common.Hash{1}, 3)

		// 		sp.CreateAccount(alice)
		// 		sp.SetCode(alice, []byte("code"))
		// 		sp.Suicide(alice)
		// 	})
		// 	It("should have reset state", func() {
		// 		sp.Reset(ctx)
		// 		Expect(sp.GetNonce(alice)).To(Equal(uint64(0)))
		// 		Expect(sp.Logs()).To(BeNil())
		// 		Expect(sp.GetRefund()).To(Equal(uint64(0)))
		// 		Expect(sp.GetSavedErr()).To(BeNil())
		// 		Expect(sp.HasSuicided(alice)).To(BeFalse())
		// 		// TODO: check the txhash and blockhash stuff
		// 		Expect(sp, state.NewStateDB(ctx, ak, bk, testutil.EvmKey, "bera"))
		// 	})
		// })

		// Describe("TestEmpty", func() {
		// 	When("account does not exist", func() {
		// 		It("should return true", func() {
		// 			Expect(sp.Empty(alice)).To(BeTrue())
		// 		})
		// 	})
		// 	When("account exists", func() {
		// 		BeforeEach(func() {
		// 			sp.CreateAccount(alice)
		// 		})
		// 		It("new account", func() {
		// 			Expect(sp.Empty(alice)).To(BeTrue())
		// 		})
		// 		It("has a balance", func() {
		// 			sp.AddBalance(alice, big.NewInt(1))
		// 			Expect(sp.Empty(alice)).To(BeFalse())
		// 		})
		// 		It("has a nonce", func() {
		// 			sp.SetNonce(alice, 1)
		// 			Expect(sp.Empty(alice)).To(BeFalse())
		// 		})
		// 		It("has code", func() {
		// 			sp.SetCode(alice, []byte{0x69})
		// 			Expect(sp.Empty(alice)).To(BeFalse())
		// 		})
		// 	})
		// })

		Describe("Test ForEachStorage", func() {
			// initialAliceBal := big.NewInt(69)
			// initialBobBal := big.NewInt(420)
			// bobCode := []byte("bobcode")

			BeforeEach(func() {
				sp.CreateAccount(alice)
				sp.CreateAccount(bob)
			})

			// It("cannot suicide eoa", func() {
			// 	Expect(sp.Suicide(alice)).To(BeFalse())
			// 	Expect(sp.HasSuicided(alice)).To(BeFalse())
			// })

			It("should iterate through storage correctly", func() {
				Expect(sp.GetCode(alice)).To(BeNil())
				var aliceStorage storage.Storage
				err := sp.ForEachStorage(alice,
					func(key, value common.Hash) bool {
						aliceStorage = append(aliceStorage,
							storage.NewSlot(key, value))
						return true
					})
				Expect(err).To(BeNil())
				Expect(len(aliceStorage)).To(BeZero())

				sp.SetState(bob, common.BytesToHash([]byte{1}), common.BytesToHash([]byte{2}))
				var bobStorage storage.Storage
				err = sp.ForEachStorage(bob,
					func(key, value common.Hash) bool {
						bobStorage = append(bobStorage, storage.NewSlot(key, value))
						return true
					})
				Expect(err).To(BeNil())
				Expect(len(bobStorage)).To(Equal(1))
				Expect(bobStorage[0].Key).
					To(Equal("0x0000000000000000000000000000000000000000000000000000000000000001"))
				Expect(bobStorage[0].Value).
					To(Equal("0x0000000000000000000000000000000000000000000000000000000000000002"))

				sp.SetState(bob, common.BytesToHash([]byte{3}), common.BytesToHash([]byte{4}))
				var bobStorage2 storage.Storage
				i := 0
				err = sp.ForEachStorage(bob,
					func(key, value common.Hash) bool {
						if i > 0 {
							return false
						}

						bobStorage2 = append(bobStorage2, storage.NewSlot(key, value))
						i++
						return true
					},
				)
				Expect(err).To(BeNil())
				Expect(len(bobStorage2)).To(Equal(1))
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
				sp.DeleteSuicides([]common.Address{alice, alice})
				Expect(ak.HasAccount(ctx, alice[:])).To(BeFalse())
				Expect(sp.GetCode(alice)).To(BeNil())
				Expect(sp.GetState(alice, common.BytesToHash([]byte{1}))).To(Equal(common.Hash{}))
			})
		})

		// 	When("address has code and balance", func() {
		// 		BeforeEach(func() {
		// 			sp.SetCode(alice, aliceCode)
		// 			sp.SetCode(bob, bobCode)
		// 			sp.AddBalance(alice, initialAliceBal)
		// 			sp.AddBalance(bob, initialBobBal)
		// 			// Give Alice some state
		// 			for i := 0; i < 5; i++ {
		// 				sp.SetState(alice, common.BytesToHash([]byte(fmt.Sprintf("key%d", i))),
		// 					common.BytesToHash([]byte(fmt.Sprintf("value%d", i))))
		// 			}
		// 			// Give Bob some state
		// 			for i := 5; i < 15; i++ {
		// 				sp.SetState(bob, common.BytesToHash([]byte(fmt.Sprintf("key%d", i))),
		// 					common.BytesToHash([]byte(fmt.Sprintf("value%d", i))))
		// 			}
		// 		})
		// 		When("alice commits suicide", func() {
		// 			BeforeEach(func() {
		// 				Expect(sp.Suicide(alice)).To(BeTrue())
		// 				Expect(sp.HasSuicided(alice)).To(BeTrue())
		// 			})
		// 			It("alice should be marked as suicidal, but not bob", func() {
		// 				Expect(sp.HasSuicided(alice)).To(BeTrue())
		// 				Expect(sp.HasSuicided(bob)).To(BeFalse())
		// 			})
		// 			It("alice should have her balance set to 0", func() {
		// 				Expect(sp.GetBalance(alice)).To(Equal(new(big.Int)))
		// 				Expect(sp.GetBalance(bob)).To(Equal(initialBobBal))
		// 			})
		// 			It("both alice and bob should have their code and state untouched", func() {
		// 				Expect(sp.GetCode(alice)).To(Equal(aliceCode))
		// 				Expect(sp.GetCode(bob)).To(Equal(bobCode))
		// 				for i := 0; i < 5; i++ {
		// 					Expect(sp.GetState(alice,
		// 						common.BytesToHash([]byte(fmt.Sprintf("key%d", i))))).
		// 						To(Equal(common.BytesToHash([]byte(fmt.Sprintf("value%d", i)))))
		// 				}

		// 				for i := 5; i < 15; i++ {
		// 					Expect(sp.GetState(bob,
		// 						common.BytesToHash([]byte(fmt.Sprintf("key%d", i))))).
		// 						To(Equal(common.BytesToHash([]byte(fmt.Sprintf("value%d", i)))))
		// 				}
		// 			})
		// 		})
		//  })

		Describe("TestAccount", func() {
			It("account does not exist", func() {
				Expect(sp.Exist(alice)).To(BeFalse())
				// Expect(sp.Empty(alice)).To(BeTrue())
				Expect(sp.GetBalance(alice)).To(Equal(new(big.Int)))
				Expect(sp.GetNonce(alice)).To(BeZero())
				Expect(sp.GetCodeHash(alice)).To(Equal(common.Hash{}))
				Expect(sp.GetCode(alice)).To(BeNil())
				Expect(sp.GetCodeSize(alice)).To(BeZero())
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

		// Describe("Test Refund", func() {
		// 	It("simple refund", func() {
		// 		sp.AddRefund(10)
		// 		Expect(sp.GetRefund()).To(Equal(uint64(10)))
		// 		sp.AddRefund(200)
		// 		Expect(sp.GetRefund()).To(Equal(uint64(210)))
		// 	})

		// 	It("nested refund", func() {
		// 		sp.AddRefund(uint64(10))
		// 		sp.SubRefund(uint64(5))
		// 		Expect(sp.GetRefund()).To(Equal(uint64(5)))
		// 	})

		// 	It("negative refund", func() {
		// 		sp.AddRefund(5)
		// 		Expect(func() { sp.SubRefund(10) }).To(Panic())
		// 	})
		// })
		// Describe("Test Log", func() {
		// 	txHash := common.BytesToHash([]byte("tx"))
		// 	blockHash := common.BytesToHash([]byte("block"))
		// 	data := []byte("bing bong bananas")
		// 	blockNumber := uint64(13)

		// 	BeforeEach(func() {
		// 		sp.Prepare(txHash, 0)
		// 	})
		// 	When("We add a log to the state", func() {
		// 		BeforeEach(func() {

		// 			sp.AddLog(&coretypes.Log{
		// 				Address:     alice,
		// 				Topics:      []common.Hash{},
		// 				Data:        data,
		// 				BlockNumber: blockNumber,
		// 				TxHash:      txHash,
		// 				TxIndex:     0,
		// 				BlockHash:   blockHash,
		// 				Index:       0,
		// 				Removed:     false,
		// 			})
		// 		})
		// 		It("should have the correct log", func() {
		// 			logs := sp.GetLogs(txHash, blockHash)
		// 			Expect(logs).To(HaveLen(1))
		// 			Expect(logs[0].Address).To(Equal(alice))
		// 			Expect(logs[0].Data).To(Equal(data))
		// 			Expect(logs[0].BlockNumber).To(Equal(blockNumber))
		// 			Expect(logs[0].TxHash).To(Equal(txHash))
		// 			Expect(logs[0].TxIndex).To(Equal(uint(0)))
		// 			Expect(logs[0].BlockHash).To(Equal(blockHash))
		// 			Expect(logs[0].Index).To(Equal(uint(0)))
		// 			Expect(logs[0].Removed).To(BeFalse())
		// 		})
		// 		When("we add a second log", func() {
		// 			BeforeEach(func() {
		// 				sp.AddLog(&coretypes.Log{
		// 					Address:     alice,
		// 					Topics:      []common.Hash{},
		// 					Data:        data,
		// 					BlockNumber: blockNumber,
		// 					TxHash:      txHash,
		// 					TxIndex:     0,
		// 					BlockHash:   blockHash,
		// 					Index:       1,
		// 					Removed:     false,
		// 				})
		// 			})
		// 			It("should have the correct logs", func() {
		// 				logs := sp.GetLogs(txHash, blockHash)
		// 				Expect(logs).To(HaveLen(2))
		// 				Expect(logs[1].Address).To(Equal(alice))
		// 				Expect(logs[1].Data).To(Equal(data))
		// 				Expect(logs[1].BlockNumber).To(Equal(blockNumber))
		// 				Expect(logs[1].TxHash).To(Equal(txHash))
		// 				Expect(logs[1].TxIndex).To(Equal(uint(0)))
		// 				Expect(logs[1].BlockHash).To(Equal(blockHash))
		// 				Expect(logs[1].Index).To(Equal(uint(1)))
		// 				Expect(logs[1].Removed).To(BeFalse())
		// 			})
		// 		})
		// 	})

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
