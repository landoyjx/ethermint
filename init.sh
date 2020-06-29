#!/bin/bash

KEY="mykey"
KEY2="wade"
CHAINID=8
MONIKER="localtestnet"

# remove existing chain environment, data and
rm -rf  ~/.halle*

make install

hallecli config keyring-backend test

# if mykey exists it should be deleted
hallecli keys add $KEY
hallecli keys add $KEY2

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
halled init $MONIKER --chain-id $CHAINID

# Set up config for CLI
hallecli config chain-id $CHAINID
hallecli config output json
hallecli config indent true
hallecli config trust-node true

# if $KEY exists it should be deleted
emintcli keys add $KEY

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
emintd init $MONIKER --chain-id $CHAINID

# Allocate genesis accounts (cosmos formatted addresses)
halled add-genesis-account $(hallecli keys show $KEY -a) 1000000000000000000hale
halled add-genesis-account $(hallecli keys show $KEY2 -a) 1000hale    --vesting-amount 600hale  --vesting-start-time  1591781100   --vesting-end-time  1591781700

# Sign genesis transaction
halled gentx --name $KEY --keyring-backend test

# Collect genesis tx
halled collect-gentxs

# Enable faucet
cat  $HOME/.emintd/config/genesis.json | jq '.app_state["faucet"]["enable_faucet"]=true' >  $HOME/.emintd/config/tmp_genesis.json && mv $HOME/.emintd/config/tmp_genesis.json $HOME/.emintd/config/genesis.json

echo -e '\n\ntestnet faucet enabled'
echo -e 'to transfer tokens to your account address use:'
echo -e "emintcli tx faucet request 100photon --from $KEY\n"


# Run this to ensure everything worked and that the genesis file is setup correctly
halled validate-genesis

# Command to run the rest server in a different terminal/window
echo -e '\n\nRun this rest-server command in a different terminal/window:'
echo -e "hallecli rest-server --laddr \"tcp://localhost:8545\" --unlock-key $KEY --chain-id $CHAINID\n\n"

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
halled start  --minimum-gas-prices  0.000001hale  --pruning=nothing --rpc.unsafe --log_level "main:info,state:info,mempool:info"
