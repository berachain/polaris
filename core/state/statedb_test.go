// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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
	"errors"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/state"
	"github.com/berachain/stargazer/core/state/types"
	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/lib/crypto"
	"github.com/berachain/stargazer/testutil"
)

var alice = testutil.Alice
var bob = testutil.Bob

var _ = Describe("StateDB", func() {
	var ak types.AccountKeeper
	var bk types.BankKeeper
	var ctx sdk.Context
	var sdb *state.StateDB

	BeforeEach(func() {
		ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
		sdb = state.NewStateDB(ctx, ak, bk, testutil.EvmKey, "abera") // todo use lf
	})

	Describe("TestCreateAccount", func() {
		AfterEach(func() {
			Expect(sdb.GetSavedErr()).To(BeNil())
		})
		It("should create account", func() {
			sdb.CreateAccount(alice)
			Expect(sdb.Exist(alice)).To(BeTrue())
		})
	})

	Describe("TestBalance", func() {
		It("should have start with zero balance", func() {
			Expect(sdb.GetBalance(alice)).To(Equal(new(big.Int)))
		})
		Context("TestAddBalance", func() {
			It("should be able to add zero", func() {
				Expect(sdb.GetBalance(alice)).To(Equal(new(big.Int)))
				sdb.AddBalance(alice, new(big.Int))
				Expect(sdb.GetBalance(alice)).To(Equal(new(big.Int)))
			})
			It("should have 100 balance", func() {
				sdb.AddBalance(alice, big.NewInt(100))
				Expect(sdb.GetBalance(alice)).To(Equal(big.NewInt(100)))
			})
			It("should panic if using negative value", func() {
				Expect(func() {
					sdb.AddBalance(alice, big.NewInt(-100))
				}).To(Panic())
			})
		})

		Context("TestSubBalance", func() {
			It("should not set balance to negative value", func() {
				sdb.SubBalance(alice, big.NewInt(100))
				Expect(sdb.GetSavedErr()).To(HaveOccurred())
				Expect(errors.Unwrap(errors.Unwrap(sdb.GetSavedErr()))).To(
					Equal(sdkerrors.ErrInsufficientFunds))
				Expect(sdb.GetBalance(alice)).To(Equal(new(big.Int)))
			})
			It("should panic if using negative value", func() {
				Expect(func() {
					sdb.SubBalance(alice, big.NewInt(-100))
				}).To(Panic())
			})
		})

		It("should handle complex balance updates", func() {
			// Initial balance for alice should be 0
			Expect(sdb.GetBalance(alice)).To(Equal(new(big.Int)))

			// Add some balance to alice
			sdb.AddBalance(alice, big.NewInt(100))
			Expect(sdb.GetBalance(alice)).To(Equal(big.NewInt(100)))

			// Subtract some balance from alice
			sdb.SubBalance(alice, big.NewInt(50))
			Expect(sdb.GetBalance(alice)).To(Equal(big.NewInt(50)))

			// Add some balance to alice
			sdb.AddBalance(alice, big.NewInt(100))
			Expect(sdb.GetBalance(alice)).To(Equal(big.NewInt(150)))

			// Subtract some balance from alice
			sdb.SubBalance(alice, big.NewInt(200))
			Expect(sdb.GetSavedErr()).To(HaveOccurred())
			Expect(errors.Unwrap(errors.Unwrap(sdb.GetSavedErr()))).To(
				Equal(sdkerrors.ErrInsufficientFunds))
			Expect(sdb.GetBalance(alice)).To(Equal(big.NewInt(150)))

		})
	})

	Describe("TestNonce", func() {
		When("account exists", func() {
			BeforeEach(func() {
				sdb.CreateAccount(alice)
			})
			It("should have start with zero nonce", func() {
				Expect(sdb.GetNonce(alice)).To(Equal(uint64(0)))
			})
			It("should have 100 nonce", func() {
				sdb.SetNonce(alice, 100)
				Expect(sdb.GetNonce(alice)).To(Equal(uint64(100)))
			})
		})
		When("account does not exist", func() {
			It("should have start with zero nonce", func() {
				Expect(sdb.GetNonce(alice)).To(Equal(uint64(0)))
			})

			It("should have 100 nonce", func() {
				sdb.SetNonce(alice, 100)
				Expect(sdb.GetNonce(alice)).To(Equal(uint64(100)))
			})
		})
	})

	Describe("TestCode & CodeHash", func() {
		When("account does not exist", func() {
			It("should have empty code hash", func() {
				Expect(sdb.GetCodeHash(alice)).To(Equal(common.Hash{}))
			})
			It("should not have code", func() { // ensure account exists
				Expect(sdb.GetCode(alice)).To(BeNil())
				Expect(sdb.GetCodeHash(alice)).To(Equal(common.Hash{}))
			})
			It("cannot set code", func() { // ensure account exists
				sdb.SetCode(alice, []byte("code"))
				Expect(sdb.GetCode(alice)).To(BeNil())
				Expect(sdb.GetCodeHash(alice)).To(Equal(common.Hash{}))
			})
		})
		When("account exists", func() {
			BeforeEach(func() {
				sdb.CreateAccount(alice)
			})
			It("should have zero code hash", func() {
				Expect(sdb.GetCodeHash(alice)).To(Equal(crypto.Keccak256Hash(nil)))
			})
			When("account has code", func() {
				BeforeEach(func() {
					sdb.SetCode(alice, []byte("code"))
				})
				It("should have code", func() {
					Expect(sdb.GetCode(alice)).To(Equal([]byte("code")))
					Expect(sdb.GetCodeHash(alice)).To(Equal(crypto.Keccak256Hash([]byte("code"))))
				})
				It("should have empty code hash", func() {
					sdb.SetCode(alice, nil)
					Expect(sdb.GetCode(alice)).To(BeNil())
					Expect(sdb.GetCodeHash(alice)).To(Equal(crypto.Keccak256Hash(nil)))
				})
			})
		})
	})

	Describe("TestRefund", func() {
		It("should have 0 refund", func() {
			Expect(sdb.GetRefund()).To(Equal(uint64(0)))
		})
		It("should have 100 refund", func() {
			sdb.AddRefund(100)
			Expect(sdb.GetRefund()).To(Equal(uint64(100)))
		})
		It("should have 0 refund", func() {
			sdb.AddRefund(100)
			sdb.SubRefund(100)
			Expect(sdb.GetRefund()).To(Equal(uint64(0)))
		})
		It("should panic and over refund", func() {
			Expect(func() {
				sdb.SubRefund(200)
			}).To(Panic())
		})
	})

	Describe("TestState", func() {
		It("should have empty state", func() {
			Expect(sdb.GetState(alice, common.Hash{3})).To(Equal(common.Hash{}))
		})
		When("set basic state", func() {
			BeforeEach(func() {
				sdb.SetState(alice, common.Hash{3}, common.Hash{1})
			})

			It("should have state", func() {
				Expect(sdb.GetState(alice, common.Hash{3})).To(Equal(common.Hash{1}))
			})

			It("should have state changed", func() {
				sdb.SetState(alice, common.Hash{3}, common.Hash{2})
				Expect(sdb.GetState(alice, common.Hash{3})).To(Equal(common.Hash{2}))
				Expect(sdb.GetCommittedState(alice, common.Hash{3})).To(Equal(common.Hash{}))
			})

			When("state is committed", func() {
				BeforeEach(func() {
					Expect(sdb.Commit()).Should(BeNil())
					It("should have committed state", func() {
						Expect(sdb.GetCommittedState(alice, common.Hash{3})).To(Equal(common.Hash{1}))
					})
					It("should maintain committed state", func() {
						sdb.SetState(alice, common.Hash{3}, common.Hash{4})
						Expect(sdb.GetCommittedState(alice, common.Hash{3})).
							To(Equal(common.Hash{1}))
						Expect(sdb.GetState(alice, common.Hash{3})).To(Equal(common.Hash{4}))
					})
				})
			})
		})

		Describe("TestExist", func() {
			It("should not exist", func() {
				Expect(sdb.Exist(alice)).To(BeFalse())
			})
			When("account is created", func() {
				BeforeEach(func() {
					sdb.CreateAccount(alice)
				})
				It("should exist", func() {
					Expect(sdb.Exist(alice)).To(BeTrue())
				})
				When("suicided", func() {
					BeforeEach(func() {
						// Only contracts can be suicided
						sdb.SetCode(alice, []byte("code"))
						Expect(sdb.Suicide(alice)).To(BeTrue())
					})
					It("should still exist", func() {
						Expect(sdb.Exist(alice)).To(BeTrue())
					})
					When("commit", func() {
						BeforeEach(func() {
							Expect(sdb.Commit()).To(BeNil())
						})
						It("should not exist", func() {
							Expect(sdb.Exist(alice)).To(BeFalse())
						})
					})
				})
			})
		})

		Describe("TestReset", func() {
			BeforeEach(func() {
				sdb.AddRefund(1000)
				sdb.AddLog(&coretypes.Log{})
				sdb.PrepareForTransition(common.Hash{1}, common.Hash{2}, 3, 4)

				sdb.CreateAccount(alice)
				sdb.SetCode(alice, []byte("code"))
				sdb.Suicide(alice)
			})
			It("should have reset state", func() {
				sdb.Reset(ctx)
				Expect(sdb.GetNonce(alice)).To(Equal(uint64(0)))
				Expect(sdb.Logs()).To(BeNil())
				Expect(sdb.GetRefund()).To(Equal(uint64(0)))
				Expect(sdb.GetSavedErr()).To(BeNil())
				Expect(sdb.HasSuicided(alice)).To(BeFalse())
				// todo check the txhash and blockhash stuff
				Expect(sdb, state.NewStateDB(ctx, ak, bk, testutil.EvmKey, "bera"))
			})
		})

		Describe("TestEmpty", func() {
			When("account does not exist", func() {
				It("should return true", func() {
					Expect(sdb.Empty(alice)).To(BeTrue())
				})
			})
			When("account exists", func() {
				BeforeEach(func() {
					sdb.CreateAccount(alice)
				})
				It("new account", func() {
					Expect(sdb.Empty(alice)).To(BeTrue())
				})
				It("has a balance", func() {
					sdb.AddBalance(alice, big.NewInt(1))
					Expect(sdb.Empty(alice)).To(BeFalse())
				})
				It("has a nonce", func() {
					sdb.SetNonce(alice, 1)
					Expect(sdb.Empty(alice)).To(BeFalse())
				})
				It("has code", func() {
					sdb.SetCode(alice, []byte{0x69})
					Expect(sdb.Empty(alice)).To(BeFalse())
				})
			})
		})

		Describe("TestSuicide", func() {
			BeforeEach(func() {
				sdb.CreateAccount(alice)
			})
			It("cannot suicide eoa", func() {
				Expect(sdb.Suicide(alice)).To(BeFalse())
				Expect(sdb.HasSuicided(alice)).To(BeFalse())
			})

			initialAliceBal := big.NewInt(69)
			initialBobBal := big.NewInt(420)
			aliceCode := []byte("alicecode")
			bobCode := []byte("bobcode")

			When("address has code and balance", func() {
				BeforeEach(func() {
					sdb.SetCode(alice, aliceCode)
					sdb.SetCode(bob, bobCode)
					sdb.AddBalance(alice, initialAliceBal)
					sdb.AddBalance(bob, initialBobBal)
					// Give Alice some state
					for i := 0; i < 5; i++ {
						sdb.SetState(alice, common.BytesToHash([]byte(fmt.Sprintf("key%d", i))),
							common.BytesToHash([]byte(fmt.Sprintf("value%d", i))))
					}
					// Give Bob some state
					for i := 5; i < 15; i++ {
						sdb.SetState(bob, common.BytesToHash([]byte(fmt.Sprintf("key%d", i))),
							common.BytesToHash([]byte(fmt.Sprintf("value%d", i))))
					}
				})
				When("alice commits suicide", func() {
					BeforeEach(func() {
						Expect(sdb.Suicide(alice)).To(BeTrue())
						Expect(sdb.HasSuicided(alice)).To(BeTrue())
					})
					It("alice should be marked as suicidal, but not bob", func() {
						Expect(sdb.HasSuicided(alice)).To(BeTrue())
						Expect(sdb.HasSuicided(bob)).To(BeFalse())
					})
					It("alice should have her balance set to 0", func() {
						Expect(sdb.GetBalance(alice)).To(Equal(new(big.Int)))
						Expect(sdb.GetBalance(bob)).To(Equal(initialBobBal))
					})
					It("both alice and bob should have their code and state untouched", func() {
						Expect(sdb.GetCode(alice)).To(Equal(aliceCode))
						Expect(sdb.GetCode(bob)).To(Equal(bobCode))
						for i := 0; i < 5; i++ {
							Expect(sdb.GetState(alice,
								common.BytesToHash([]byte(fmt.Sprintf("key%d", i))))).
								To(Equal(common.BytesToHash([]byte(fmt.Sprintf("value%d", i)))))
						}

						for i := 5; i < 15; i++ {
							Expect(sdb.GetState(bob,
								common.BytesToHash([]byte(fmt.Sprintf("key%d", i))))).
								To(Equal(common.BytesToHash([]byte(fmt.Sprintf("value%d", i)))))
						}
					})
					When("commit is called", func() {
						BeforeEach(func() {
							_ = sdb.Commit()
						})
						It("alice should have her code and state wiped, but not bob", func() {
							Expect(sdb.GetCode(alice)).To(BeNil())
							Expect(sdb.GetCode(bob)).To(Equal(bobCode))
							var aliceStorage types.Storage
							err := sdb.ForEachStorage(alice,
								func(key, value common.Hash) bool {
									aliceStorage = append(aliceStorage,
										types.NewState(key, value))
									return true
								})
							Expect(err).To(BeNil())
							Expect(len(aliceStorage)).To(BeZero())

							var bobStorage types.Storage
							err = sdb.ForEachStorage(bob,
								func(key, value common.Hash) bool {
									bobStorage = append(bobStorage, types.NewState(key, value))
									return true
								})
							Expect(err).To(BeNil())
							Expect(len(bobStorage)).To(Equal(10))
						})
					})

				})
			})
		})
		Describe("TestAccount", func() {
			It("account does not exist", func() {
				Expect(sdb.Exist(alice)).To(BeFalse())
				Expect(sdb.Empty(alice)).To(BeTrue())
				Expect(sdb.GetBalance(alice)).To(Equal(new(big.Int)))
				Expect(sdb.GetNonce(alice)).To(BeZero())
				Expect(sdb.GetCodeHash(alice)).To(Equal(common.Hash{}))
				Expect(sdb.GetCode(alice)).To(BeNil())
				Expect(sdb.GetCodeSize(alice)).To(BeZero())
				Expect(sdb.GetState(alice, common.Hash{})).To(Equal(common.Hash{}))
				Expect(sdb.GetRefund()).To(BeZero())
				Expect(sdb.GetCommittedState(alice, common.Hash{})).To(Equal(common.Hash{}))
			})
			When("account exists", func() {
				BeforeEach(func() {
					sdb.AddBalance(alice, big.NewInt(56))
					sdb.SetNonce(alice, 59)
				})
				It("accidental override account", func() {
					// override
					sdb.CreateAccount(alice)

					// check balance is not reset
					Expect(sdb.GetBalance(alice)).To(Equal(big.NewInt(56)))
				})
			})

		})

		Describe("TestSnapshot", func() {
			key := common.BytesToHash([]byte("key"))
			value1 := common.BytesToHash([]byte("value1"))
			value2 := common.BytesToHash([]byte("value2"))
			It("simple revert", func() {
				revision := sdb.Snapshot()
				Expect(revision).To(Equal(0))
				sdb.SetState(alice, key, value1)
				Expect(sdb.GetState(alice, key)).To(Equal(value1))
				sdb.RevertToSnapshot(revision)
				Expect(sdb.GetState(alice, key)).To(Equal(common.Hash{}))
			})
			It("nested snapshot & revert", func() {
				revision1 := sdb.Snapshot()
				Expect(revision1).To(Equal(0))
				sdb.SetState(alice, key, value1)
				revision2 := sdb.Snapshot()
				Expect(revision2).To(Equal(1))
				sdb.SetState(alice, key, value2)
				Expect(sdb.GetState(alice, key)).To(Equal(value2))

				sdb.RevertToSnapshot(revision2)
				Expect(sdb.GetState(alice, key)).To(Equal(value1))

				sdb.RevertToSnapshot(revision1)
				Expect(sdb.GetState(alice, key)).To(Equal(common.Hash{}))
			})
			It("jump revert", func() {
				revision1 := sdb.Snapshot()
				Expect(revision1).To(Equal(0))
				sdb.SetState(alice, key, value1)
				sdb.Snapshot()
				sdb.SetState(alice, key, value2)
				Expect(sdb.GetState(alice, key)).To(Equal(value2))
				sdb.RevertToSnapshot(revision1)
				Expect(sdb.GetState(alice, key)).To(Equal(common.Hash{}))
			})
		})

		Describe("Test Refund", func() {
			It("simple refund", func() {
				sdb.AddRefund(10)
				Expect(sdb.GetRefund()).To(Equal(uint64(10)))
				sdb.AddRefund(200)
				Expect(sdb.GetRefund()).To(Equal(uint64(210)))
			})

			It("nested refund", func() {
				sdb.AddRefund(uint64(10))
				sdb.SubRefund(uint64(5))
				Expect(sdb.GetRefund()).To(Equal(uint64(5)))
			})

			It("negative refund", func() {
				sdb.AddRefund(5)
				Expect(func() { sdb.SubRefund(10) }).To(Panic())
			})
		})
		Describe("Test Log", func() {
			txHash := common.BytesToHash([]byte("tx"))
			blockHash := common.BytesToHash([]byte("block"))
			data := []byte("bing bong bananas")
			blockNumber := uint64(13)

			BeforeEach(func() {
				sdb.PrepareForTransition(blockHash, txHash, 0, 0)
			})
			When("We add a log to the state", func() {
				BeforeEach(func() {

					sdb.AddLog(&coretypes.Log{
						Address:     alice,
						Topics:      []common.Hash{},
						Data:        data,
						BlockNumber: blockNumber,
						TxHash:      txHash,
						TxIndex:     0,
						BlockHash:   blockHash,
						Index:       0,
						Removed:     false,
					})
				})
				It("should have the correct log", func() {
					logs := sdb.Logs()
					Expect(logs).To(HaveLen(1))
					Expect(logs[0].Address).To(Equal(alice))
					Expect(logs[0].Data).To(Equal(data))
					Expect(logs[0].BlockNumber).To(Equal(blockNumber))
					Expect(logs[0].TxHash).To(Equal(txHash))
					Expect(logs[0].TxIndex).To(Equal(uint(0)))
					Expect(logs[0].BlockHash).To(Equal(blockHash))
					Expect(logs[0].Index).To(Equal(uint(0)))
					Expect(logs[0].Removed).To(BeFalse())
				})
				When("we add a second log", func() {
					BeforeEach(func() {
						sdb.AddLog(&coretypes.Log{
							Address:     alice,
							Topics:      []common.Hash{},
							Data:        data,
							BlockNumber: blockNumber,
							TxHash:      txHash,
							TxIndex:     0,
							BlockHash:   blockHash,
							Index:       1,
							Removed:     false,
						})
					})
					It("should have the correct logs", func() {
						logs := sdb.Logs()
						Expect(logs).To(HaveLen(2))
						Expect(logs[1].Address).To(Equal(alice))
						Expect(logs[1].Data).To(Equal(data))
						Expect(logs[1].BlockNumber).To(Equal(blockNumber))
						Expect(logs[1].TxHash).To(Equal(txHash))
						Expect(logs[1].TxIndex).To(Equal(uint(0)))
						Expect(logs[1].BlockHash).To(Equal(blockHash))
						Expect(logs[1].Index).To(Equal(uint(1)))
						Expect(logs[1].Removed).To(BeFalse())
					})
				})
			})
		})

		Describe("TestSavedErr", func() {
			When("if we see an error", func() {
				It("should have an error", func() {
					sdb.TransferBalance(alice, bob, big.NewInt(100))
					Expect(sdb.GetSavedErr()).To(HaveOccurred())
				})
			})

		})

		Describe("TestRevertSnapshot", func() {
			key := common.BytesToHash([]byte("key"))
			value := common.BytesToHash([]byte("value"))

			When("We make a bunch of arbitrary changes", func() {
				BeforeEach(func() {
					sdb.SetNonce(alice, 1)
					sdb.AddBalance(alice, big.NewInt(100))
					sdb.SetCode(alice, []byte("hello world"))
					sdb.SetState(alice, key, value)
					sdb.SetNonce(bob, 1)
				})
				When("we take a snapshot", func() {
					var revision int
					BeforeEach(func() {
						revision = sdb.Snapshot()
					})
					When("we do more changes", func() {
						AfterEach(func() {
							sdb.RevertToSnapshot(revision)
							Expect(sdb.GetNonce(alice)).To(Equal(uint64(1)))
							Expect(sdb.GetBalance(alice)).To(Equal(big.NewInt(100)))
							Expect(sdb.GetCode(alice)).To(Equal([]byte("hello world")))
							Expect(sdb.GetState(alice, key)).To(Equal(value))
							Expect(sdb.GetNonce(bob)).To(Equal(uint64(1)))
						})

						It("if we change balance", func() {
							sdb.AddBalance(alice, big.NewInt(100))
						})

						It("if we change nonce", func() {
							sdb.SetNonce(alice, 2)
						})

						It("if we change code", func() {
							sdb.SetCode(alice, []byte("goodbye world"))
						})

						It("if we change state", func() {
							sdb.SetState(alice, key, common.Hash{})
						})

						It("if we change nonce of another account", func() {
							sdb.SetNonce(bob, 2)
						})
					})

					When("we make a nested snapshot", func() {
						var revision2 int
						BeforeEach(func() {
							sdb.SetState(alice, key, common.Hash{2})
							revision2 = sdb.Snapshot()
						})
						When("we revert to snapshot ", (func() {
							It("revision 2", func() {
								sdb.RevertToSnapshot(revision2)
								Expect(sdb.GetState(alice, key)).To(Equal(common.Hash{2}))
							})
							It("revision 1", func() {
								sdb.RevertToSnapshot(revision)
								Expect(sdb.GetState(alice, key)).To(Equal(value))
							})
						}))
					})
				})
				When("we revert to an invalid snapshot", func() {
					It("should panic", func() {
						Expect(func() {
							sdb.RevertToSnapshot(100)
						}).To(Panic())
					})
				})
			})
		})
	})
})
