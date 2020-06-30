#!/bin/bash

KEY=$2
CHAINID=8
MONIKER=$1

# remove existing chain environment, data and 
rm -rf ~/.emint*

make install

hallecli config keyring-backend test

# if mykey exists it should be deleted
hallecli keys add $KEY

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
halled init $MONIKER --chain-id $CHAINID

# Set up config for CLI
hallecli config chain-id $CHAINID
hallecli config output json
hallecli config indent true
hallecli config trust-node true

# Allocate genesis accounts (cosmos formatted addresses)
halled add-genesis-account $(hallecli keys show $KEY -a) 1000000000000000000hale

# Sign genesis transaction
halled gentx --name $KEY --keyring-backend test

