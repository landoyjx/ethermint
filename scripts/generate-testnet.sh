# 1. hallecli init
echot "hallecli init"
rm -rf ~/.halle*
hallecli config keyring-backend test
hallecli config chain-id 8
hallecli config output json
hallecli config indent true
hallecli config trust-node true

# 2. init
echo "halled init"
rm -rf testnet/*
halled init node0 --chain-id 8 --home testnet/node0
halled init node1 --chain-id 8 --home testnet/node1
halled init node2 --chain-id 8 --home testnet/node2
halled init node3 --chain-id 8 --home testnet/node3
halled init node4 --chain-id 8 --home testnet/node4
halled init node5 --chain-id 8 --home testnet/node5
halled init node6 --chain-id 8 --home testnet/node6
halled init node7 --chain-id 8 --home testnet/node7
halled init node8 --chain-id 8 --home testnet/node8
halled init node9 --chain-id 8 --home testnet/node9
halled init node10 --chain-id 8 --home testnet/node10
halled init node11 --chain-id 8 --home testnet/node11
halled init node12 --chain-id 8 --home testnet/node12
halled init node13 --chain-id 8 --home testnet/node13
halled init node14 --chain-id 8 --home testnet/node14
halled init node15 --chain-id 8 --home testnet/node15
halled init node16 --chain-id 8 --home testnet/node16
halled init node17 --chain-id 8 --home testnet/node17
halled init node18 --chain-id 8 --home testnet/node18
halled init node19 --chain-id 8 --home testnet/node19


# 3. create genesis accounts
echo "hallecli keys add"
hallecli keys add mykey0
hallecli keys add mykey1
hallecli keys add mykey2
hallecli keys add mykey3
hallecli keys add mykey4
hallecli keys add mykey5
hallecli keys add mykey6
hallecli keys add mykey7
hallecli keys add mykey8
hallecli keys add mykey9
hallecli keys add mykey10
hallecli keys add mykey11
hallecli keys add mykey12
hallecli keys add mykey13
hallecli keys add mykey14
hallecli keys add mykey15
hallecli keys add mykey16
hallecli keys add mykey17
hallecli keys add mykey18
hallecli keys add mykey19

# 4. add genesis accounts to genesis.json
echo "halled add-genesis-account"
halled add-genesis-account $(hallecli keys show mykey0 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey1 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey2 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey3 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey4 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey5 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey6 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey7 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey8 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey9 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey10 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey11 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey12 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey13 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey14 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey15 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey16 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey17 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey18 -a) 1000000000000000000hale --home testnet/node0
halled add-genesis-account $(hallecli keys show mykey19 -a) 1000000000000000000hale --home testnet/node0
#
# # halled add-genesis-account $(hallecli keys show mykey1 -a) 1000000000000000000hale --home testnet/node1
# # halled add-genesis-account $(hallecli keys show mykey2 -a) 1000000000000000000hale --home testnet/node2
halled add-genesis-account $(hallecli keys show mykey1 -a) 1000000000000000000hale --home testnet/node1
halled add-genesis-account $(hallecli keys show mykey2 -a) 1000000000000000000hale --home testnet/node2
halled add-genesis-account $(hallecli keys show mykey3 -a) 1000000000000000000hale --home testnet/node3
halled add-genesis-account $(hallecli keys show mykey4 -a) 1000000000000000000hale --home testnet/node4
halled add-genesis-account $(hallecli keys show mykey5 -a) 1000000000000000000hale --home testnet/node5
halled add-genesis-account $(hallecli keys show mykey6 -a) 1000000000000000000hale --home testnet/node6
halled add-genesis-account $(hallecli keys show mykey7 -a) 1000000000000000000hale --home testnet/node7
halled add-genesis-account $(hallecli keys show mykey8 -a) 1000000000000000000hale --home testnet/node8
halled add-genesis-account $(hallecli keys show mykey9 -a) 1000000000000000000hale --home testnet/node9
halled add-genesis-account $(hallecli keys show mykey10 -a) 1000000000000000000hale --home testnet/node10
halled add-genesis-account $(hallecli keys show mykey11 -a) 1000000000000000000hale --home testnet/node11
halled add-genesis-account $(hallecli keys show mykey12 -a) 1000000000000000000hale --home testnet/node12
halled add-genesis-account $(hallecli keys show mykey13 -a) 1000000000000000000hale --home testnet/node13
halled add-genesis-account $(hallecli keys show mykey14 -a) 1000000000000000000hale --home testnet/node14
halled add-genesis-account $(hallecli keys show mykey15 -a) 1000000000000000000hale --home testnet/node15
halled add-genesis-account $(hallecli keys show mykey16 -a) 1000000000000000000hale --home testnet/node16
halled add-genesis-account $(hallecli keys show mykey17 -a) 1000000000000000000hale --home testnet/node17
halled add-genesis-account $(hallecli keys show mykey18 -a) 1000000000000000000hale --home testnet/node18
halled add-genesis-account $(hallecli keys show mykey19 -a) 1000000000000000000hale --home testnet/node19

#
# # 5. create gentxs
echo "create gentxs-----------------------------"
halled gentx --name mykey0 --home testnet/node0 --ip 192.168.20.2 --node-id $(halled tendermint show-node-id --home testnet/node0) --keyring-backend test
halled gentx --name mykey1 --home testnet/node1 --ip 192.168.20.3 --node-id $(halled tendermint show-node-id --home testnet/node1) --keyring-backend test
halled gentx --name mykey2 --home testnet/node2 --ip 192.168.20.4 --node-id $(halled tendermint show-node-id --home testnet/node2) --keyring-backend test
halled gentx --name mykey3 --home testnet/node3 --ip 192.168.20.5 --node-id $(halled tendermint show-node-id --home testnet/node3) --keyring-backend test
halled gentx --name mykey4 --home testnet/node4 --ip 192.168.20.6 --node-id $(halled tendermint show-node-id --home testnet/node4) --keyring-backend test
halled gentx --name mykey5 --home testnet/node5 --ip 192.168.20.7 --node-id $(halled tendermint show-node-id --home testnet/node5) --keyring-backend test
halled gentx --name mykey6 --home testnet/node6 --ip 192.168.20.8 --node-id $(halled tendermint show-node-id --home testnet/node6) --keyring-backend test
halled gentx --name mykey7 --home testnet/node7 --ip 192.168.20.9 --node-id $(halled tendermint show-node-id --home testnet/node7) --keyring-backend test
halled gentx --name mykey8 --home testnet/node8 --ip 192.168.20.10 --node-id $(halled tendermint show-node-id --home testnet/node8) --keyring-backend test
halled gentx --name mykey9 --home testnet/node9 --ip 192.168.20.11 --node-id $(halled tendermint show-node-id --home testnet/node9) --keyring-backend test
halled gentx --name mykey10 --home testnet/node10 --ip 192.168.20.12 --node-id $(halled tendermint show-node-id --home testnet/node10) --keyring-backend test
halled gentx --name mykey11 --home testnet/node11 --ip 192.168.20.13 --node-id $(halled tendermint show-node-id --home testnet/node11) --keyring-backend test
halled gentx --name mykey12 --home testnet/node12 --ip 192.168.20.14 --node-id $(halled tendermint show-node-id --home testnet/node12) --keyring-backend test
halled gentx --name mykey13 --home testnet/node13 --ip 192.168.20.15 --node-id $(halled tendermint show-node-id --home testnet/node13) --keyring-backend test
halled gentx --name mykey14 --home testnet/node14 --ip 192.168.20.16 --node-id $(halled tendermint show-node-id --home testnet/node14) --keyring-backend test
halled gentx --name mykey15 --home testnet/node15 --ip 192.168.20.17 --node-id $(halled tendermint show-node-id --home testnet/node15) --keyring-backend test
halled gentx --name mykey16 --home testnet/node16 --ip 192.168.20.18 --node-id $(halled tendermint show-node-id --home testnet/node16) --keyring-backend test
halled gentx --name mykey17 --home testnet/node17 --ip 192.168.20.19 --node-id $(halled tendermint show-node-id --home testnet/node17) --keyring-backend test
halled gentx --name mykey18 --home testnet/node18 --ip 192.168.20.20 --node-id $(halled tendermint show-node-id --home testnet/node18) --keyring-backend test
halled gentx --name mykey19 --home testnet/node19 --ip 192.168.20.21 --node-id $(halled tendermint show-node-id --home testnet/node19) --keyring-backend test
#
#
# # 6. collect-gentxs to genesis.json
echo "collect-gentxs to genesis.json"
cp testnet/node1/config/gentx/* testnet/node0/config/gentx/
cp testnet/node2/config/gentx/* testnet/node0/config/gentx/
cp testnet/node3/config/gentx/* testnet/node0/config/gentx/
cp testnet/node4/config/gentx/* testnet/node0/config/gentx/
cp testnet/node5/config/gentx/* testnet/node0/config/gentx/
cp testnet/node6/config/gentx/* testnet/node0/config/gentx/
cp testnet/node7/config/gentx/* testnet/node0/config/gentx/
cp testnet/node8/config/gentx/* testnet/node0/config/gentx/
cp testnet/node9/config/gentx/* testnet/node0/config/gentx/
cp testnet/node10/config/gentx/* testnet/node0/config/gentx/
cp testnet/node11/config/gentx/* testnet/node0/config/gentx/
cp testnet/node12/config/gentx/* testnet/node0/config/gentx/
cp testnet/node13/config/gentx/* testnet/node0/config/gentx/
cp testnet/node14/config/gentx/* testnet/node0/config/gentx/
cp testnet/node15/config/gentx/* testnet/node0/config/gentx/
cp testnet/node16/config/gentx/* testnet/node0/config/gentx/
cp testnet/node17/config/gentx/* testnet/node0/config/gentx/
cp testnet/node18/config/gentx/* testnet/node0/config/gentx/
cp testnet/node19/config/gentx/* testnet/node0/config/gentx/

halled collect-gentxs --home testnet/node0
#
#
# # 7. collect node1 and node2 genesis.json gentxs, copy to node0 genesis.json gentxs, copy node0 genesis.json to replace others
echo "collect node1 and node2 genesis.json gentxs, copy to node0 genesis.json gentxs, copy node0 genesis.json to replace others"
rm -f testnet/node1/config/genesis.json
rm -f testnet/node2/config/genesis.json
rm -f testnet/node3/config/genesis.json
rm -f testnet/node4/config/genesis.json
rm -f testnet/node5/config/genesis.json
rm -f testnet/node6/config/genesis.json
rm -f testnet/node7/config/genesis.json
rm -f testnet/node8/config/genesis.json
rm -f testnet/node9/config/genesis.json
rm -f testnet/node10/config/genesis.json
rm -f testnet/node11/config/genesis.json
rm -f testnet/node12/config/genesis.json
rm -f testnet/node13/config/genesis.json
rm -f testnet/node14/config/genesis.json
rm -f testnet/node15/config/genesis.json
rm -f testnet/node16/config/genesis.json
rm -f testnet/node17/config/genesis.json
rm -f testnet/node18/config/genesis.json
rm -f testnet/node19/config/genesis.json

cp testnet/node0/config/genesis.json testnet/node1/config/
cp testnet/node0/config/genesis.json testnet/node2/config/
cp testnet/node0/config/genesis.json testnet/node3/config/
cp testnet/node0/config/genesis.json testnet/node4/config/
cp testnet/node0/config/genesis.json testnet/node5/config/
cp testnet/node0/config/genesis.json testnet/node6/config/
cp testnet/node0/config/genesis.json testnet/node7/config/
cp testnet/node0/config/genesis.json testnet/node8/config/
cp testnet/node0/config/genesis.json testnet/node9/config/
cp testnet/node0/config/genesis.json testnet/node10/config/
cp testnet/node0/config/genesis.json testnet/node11/config/
cp testnet/node0/config/genesis.json testnet/node12/config/
cp testnet/node0/config/genesis.json testnet/node13/config/
cp testnet/node0/config/genesis.json testnet/node14/config/
cp testnet/node0/config/genesis.json testnet/node15/config/
cp testnet/node0/config/genesis.json testnet/node16/config/
cp testnet/node0/config/genesis.json testnet/node17/config/
cp testnet/node0/config/genesis.json testnet/node18/config/
cp testnet/node0/config/genesis.json testnet/node19/config/

# # 8. config each node's config.toml persistent_peers to the other two node's node-id@node-ip:26656
os=`uname -a`
mac='Darwin'
peers=`halled tendermint show-node-id --home testnet/node0`@192.168.20.2:26656,`halled tendermint show-node-id --home testnet/node1`@192.168.20.3:26656,`halled tendermint show-node-id --home testnet/node2`@192.168.20.4:26656
if [[ $os =~ $mac ]];then
    gsed -i '175,175d' testnet/node0/config/config.toml
    gsed -i "174a persistent_peers = \"$peers\"" testnet/node0/config/config.toml
    gsed -i '175,175d' testnet/node1/config/config.toml
    gsed -i "174a persistent_peers = \"$peers\"" testnet/node1/config/config.toml
    gsed -i '175,175d' testnet/node2/config/config.toml
    gsed -i "174a persistent_peers = \"$peers\"" testnet/node2/config/config.toml
else
    sed -i '175,175d' testnet/node0/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node0/config/config.toml
    sed -i '175,175d' testnet/node1/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node1/config/config.toml
    sed -i '175,175d' testnet/node2/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node2/config/config.toml
    sed -i '175,175d' testnet/node3/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node3/config/config.toml
    sed -i '175,175d' testnet/node4/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node4/config/config.toml
    sed -i '175,175d' testnet/node5/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node5/config/config.toml
    sed -i '175,175d' testnet/node6/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node6/config/config.toml
    sed -i '175,175d' testnet/node7/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node7/config/config.toml
    sed -i '175,175d' testnet/node8/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node8/config/config.toml
    sed -i '175,175d' testnet/node9/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node9/config/config.toml
    sed -i '175,175d' testnet/node10/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node10/config/config.toml
    sed -i '175,175d' testnet/node11/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node11/config/config.toml
    sed -i '175,175d' testnet/node12/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node12/config/config.toml
    sed -i '175,175d' testnet/node13/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node13/config/config.toml
    sed -i '175,175d' testnet/node14/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node14/config/config.toml
    sed -i '175,175d' testnet/node15/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node15/config/config.toml
    sed -i '175,175d' testnet/node16/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node16/config/config.toml
    sed -i '175,175d' testnet/node17/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node17/config/config.toml
    sed -i '175,175d' testnet/node18/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node18/config/config.toml
    sed -i '175,175d' testnet/node19/config/config.toml
    sed -i "174a persistent_peers = \"$peers\"" testnet/node19/config/config.toml
fi


# 9. start each node, halled start --home node* --rpc.unsafe --log_level "main:info,state:info,mempool:info"
echo -e "\n------Enjoy it!------"
