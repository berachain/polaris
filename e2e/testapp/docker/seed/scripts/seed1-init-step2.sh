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

if [ -z "$CHAINID" ]; then
    CHAINID="brickchain-666"
fi
if [ -z "$KEYRING" ]; then
    KEYRING="test"
fi
if [ -z "$HOMEDIR" ]; then
    HOMEDIR="/.polard"
fi

KEY="$1"

polard genesis add-genesis-account $KEY 100000000000000000000000000abera,100000000000000000000000000stgusdc --keyring-backend $KEYRING --home "$HOMEDIR"

polard genesis gentx $KEY 1000000000000000000000abera --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR" \
    --moniker="$KEY" \
    --identity="identity of $KEY" \
    --details="This is $KEY" \
    --security-contact="brick@berachain.com" \
    --website="https://quantumwn.org/"
