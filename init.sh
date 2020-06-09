#!/bin/bash

KEY="mykey"
KEY2="wade"
CHAINID=8
MONIKER="mymoniker"

# remove existing chain environment, data and
rm -rf ~/.emint*

make install

emintcli config keyring-backend test

# if mykey exists it should be deleted
emintcli keys add $KEY
emintcli keys add $KEY2

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
emintd init $MONIKER --chain-id $CHAINID

# Set up config for CLI
emintcli config chain-id $CHAINID
emintcli config output json
emintcli config indent true
emintcli config trust-node true

# Allocate genesis accounts (cosmos formatted addresses)
# emintd add-genesis-account $(emintcli keys show $KEY -a) 800000000hale
# emintd add-genesis-account $(emintcli keys show wade1 -a)  10000hale  --vesting-end-time 1591334071  --vesting-amount  1000hale

#emintd add-genesis-account $(emintcli keys show $KEY -a)   800000000hale  --vesting-end-time 1591334671  --vesting-amount  1000hale
emintd add-genesis-account $(emintcli keys show $KEY -a) 800000000hale --vesting-end-time 1594149150     --vesting-amount  10hale
emintd add-genesis-account $(emintcli keys show $KEY2 -a) 1000hale  --vesting-end-time 1591557150     --vesting-amount  500hale

# Sign genesis transaction
emintd gentx --name $KEY --keyring-backend test

# Collect genesis tx
emintd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
emintd validate-genesis

# Command to run the rest server in a different terminal/window
echo -e '\n\nRun this rest-server command in a different terminal/window:'
echo -e "emintcli rest-server --laddr \"tcp://localhost:8545\" --unlock-key $KEY --chain-id $CHAINID\n\n"

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
emintd start --pruning=nothing --rpc.unsafe --log_level "main:info,state:info,mempool:info"
