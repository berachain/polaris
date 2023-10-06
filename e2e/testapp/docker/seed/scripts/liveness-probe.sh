#!/bin/bash
# SPDX-License-Identifier: BUSL-1.1
#
# Copyright (C) 2023, Berachain Foundation. All rights reserved.
# Use of this software is govered by the Business Source License included
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


# Execute the cURL command and capture the response
response=$(curl -s -X POST -H "Content-Type: application/json" -d '{
    "jsonrpc":"2.0",
    "method": "eth_blockNumber",
    "params": [],
    "id": 1
}' http://localhost:8545)

height=$(echo "$response" | jq -r '.result')

file="last_height.txt"

# Check if the file exists
if [ -e "$file" ]; then
  # Read the contents of the file into the result variable
  last_height=$(cat "$file")
else
  # File does not exist, set result to an empty string
  last_height=""
fi

rm $file
echo "$height" >> $file

# Check if the two input strings are equal
if [ "$height" == "$last_height" ]; then
  # Strings are equal, return 1
  exit 1
else
  # Strings are not equal, return 0
  exit 0
fi
