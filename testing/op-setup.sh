#!/bin/bash

# Stage 1: Clone repos
cd ~/
mkdir op-stack-deployment
cd op-stack-deployment
git clone https://github.com/ethereum-optimism/optimism.git --depth=1
git clone https://github.com/ethereum-optimism/op-geth.git --depth=1

# Stage 2: Setup nvm
brew install nvm

# Check if ~/.nvm directory doesn't exist
if [ ! -d "$nvm_dir" ]; then
    mkdir "$nvm_dir"
    echo "Created ~/.nvm directory."
fi
# hardhat needs ^16.0.0
. ~/.nvm/nvm.sh
. ~/.zshrc
. $(brew --prefix nvm)/nvm.sh  # if installed via Brew
nvm install 16
nvm use 16

# Stage 3: Install op-node op-batcher op-proposer
cd optimism
yarn install
make op-node op-batcher op-proposer
yarn build
cd ..

# Stage 4: Install op-geth
cd op-geth
make geth
cd ..

# Stage 5: Install direnv
brew install direnv
direnv_hook='eval "$(direnv hook zsh)"'
zsh_config="$HOME/.zshrc"

# Check if the direnv hook already exists in the file
if ! grep -qF "$direnv_hook" "$zsh_config"; then
    # Append the direnv hook to the file
    echo "$direnv_hook" >> "$zsh_config"
    echo "direnv hook added to $zsh_config"
else
    echo "direnv hook already exists in $zsh_config"
fi
source $zsh_config