#!/bin/bash

KEY=$2
CHAINID=8
MONIKER=$1

# remove existing chain environment, data and 
rm -rf ~/.emint*

make install

emintcli config keyring-backend test

# if mykey exists it should be deleted
emintcli keys add $KEY

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
emintd init $MONIKER --chain-id $CHAINID

# Set up config for CLI
emintcli config chain-id $CHAINID
emintcli config output json
emintcli config indent true
emintcli config trust-node true

# Allocate genesis accounts (cosmos formatted addresses)
emintd add-genesis-account $(emintcli keys show $KEY -a) 1000000000000000000stake

# Sign genesis transaction
emintd gentx --name $KEY --keyring-backend test

