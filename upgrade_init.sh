#!/bin/bash


wade="wade"
jack="jack"
CHAINID=8
MONIKER="mymoniker"


echo "===> Deleting previous data"

rm -rf  ~/.halle*

make install

echo "===> Setting up CLI"
hallecli config chain-id $CHAINID
hallecli config output json
hallecli config indent true
hallecli config trust-node true
hallecli config keyring-backend test

# if mykey exists it should be deleted
hallecli keys add $wade
hallecli keys add $jack

echo "===> Initializing chain"
# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
halled init $MONIKER --chain-id $CHAINID

echo "===> Editing genesis.json"
# sed -i 's/desmos/udaric/g' $HOME/.desmosd/config/genesis.json

os=`uname -a`
mac='Darwin'
if [[ $os =~ $mac ]];then
  sed -i ''  's/"voting_period": "172800000000000"/"voting_period": "120000000000"/g'  ~/.halled/config/genesis.json
else
  sed -i 's/"voting_period": "172800000000000"/"voting_period": "120000000000"/g'  ~/.halled/config/genesis.json
fi

echo "===> Creating genesis accounts"
halled add-genesis-account $(hallecli keys show $wade -a)  100000000000000000000000000uhale
halled add-genesis-account $(hallecli keys show $jack -a)  100000000000000000000000000uhale



echo "===> Collecting genesis trasanctions"
#halled gentx --amount 1000000000uhale --name jack
halled gentx --name $jack --keyring-backend test


halled collect-gentxs

# For handling the upgrades
# go build -o $GOBIN/cosmosd


echo -e '\n\nRun this rest-server command in a different terminal/window:'
echo -e "hallecli rest-server --laddr \"tcp://localhost:8545\" --unlock-key $KEY --chain-id $CHAINID\n\n"


echo "===> shutdown the chain"
echo "make install new version with app.upgradeKeeper.SetUpgradeHandler(\"test\", func(ctx sdk.Context, plan upgrade.Plan) {})"
echo "run again cmd: halled start   --pruning nothing"

echo "===> Submitting governance proposal"
echo "hallecli tx gov submit-proposal software-upgrade \"test\"  --title \"test software upgrade proposal\"  --description \"something about the proposal here\"   --deposit 1000000000000uhale   --upgrade-height=30 --from jack   --yes --trace   -b block -y"
echo "===> query proposal"
echo " hallecli  query gov  proposals"
echo "===> Waiting for transaction to be effective"
echo "===> Voting governance proposal"
echo "hallecli tx gov vote 1 yes --from jack --yes --trace  -b block -y"


echo "===> Starting chain"
halled start   --pruning nothing
