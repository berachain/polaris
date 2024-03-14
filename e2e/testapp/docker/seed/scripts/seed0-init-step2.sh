# SPDX-License-Identifier: BUSL-1.1
#
# Copyright (C) 2023, Berachain Foundation. All rights reserved.
# Use of this software is governed by the Business Source License included
# in the LICENSE file of this repository and at www.mariadb.com/bsl11.
#
# ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
# TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
# VERSIONS OF THE LICENSED WORK.
#
# THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
# LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
# LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
#
# TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
# AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
# EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
# TITLE.

if [ -z "$HOMEDIR" ]; then
    HOMEDIR="/.polard"
fi
if [ -z "$KEYRING" ]; then
    KEYRING="test"
fi
if [ -z "$KEYALGO" ]; then
    KEYALGO="eth_secp256k1"
fi

polard genesis collect-gentxs --home "$HOMEDIR"

polard genesis validate-genesis --home "$HOMEDIR"

# # faucet
# polard keys add faucet --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
# polard genesis add-genesis-account faucet 1000000000000000000000000000abera,1000000000000000000000000000stgusdc --keyring-backend $KEYRING --home "$HOMEDIR"

# # # Test Account
# # absurd surge gather author blanket acquire proof struggle runway attract cereal quiz tattoo shed almost sudden survey boring film memory picnic favorite verb tank
# # 0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306
# polard genesis add-genesis-account cosmos1yrene6g2zwjttemf0c65fscg8w8c55w58yh8rl 1000000000000000000000000000abera,1000000000000000000000000000stgusdc --keyring-backend $KEYRING --home "$HOMEDIR"
