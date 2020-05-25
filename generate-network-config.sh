# 1. emintcli init
rm -rf ~/.emint*
emintcli config keyring-backend test
emintcli config chain-id 8
emintcli config output json
emintcli config indent true
emintcli config trust-node true

# 2. init
rm -rf testnet/*
emintd init node0 --chain-id 8 --home testnet/node0
emintd init node1 --chain-id 8 --home testnet/node1
emintd init node2 --chain-id 8 --home testnet/node2

# 3. create genesis accounts
emintcli keys add mykey0
emintcli keys add mykey1
emintcli keys add mykey2

# 4. add genesis accounts to genesis.json
emintd add-genesis-account $(emintcli keys show mykey0 -a) 1000000000000000000hale --home testnet/node0
emintd add-genesis-account $(emintcli keys show mykey1 -a) 1000000000000000000hale --home testnet/node0
emintd add-genesis-account $(emintcli keys show mykey2 -a) 1000000000000000000hale --home testnet/node0
emintd add-genesis-account $(emintcli keys show mykey1 -a) 1000000000000000000hale --home testnet/node1
emintd add-genesis-account $(emintcli keys show mykey2 -a) 1000000000000000000hale --home testnet/node2

# 5. create gentxs
emintd gentx --name mykey0 --home testnet/node0 --ip 192.168.20.2 --node-id $(emintd tendermint show-node-id --home testnet/node0) --keyring-backend test
emintd gentx --name mykey1 --home testnet/node1 --ip 192.168.20.3 --node-id $(emintd tendermint show-node-id --home testnet/node1) --keyring-backend test
emintd gentx --name mykey2 --home testnet/node2 --ip 192.168.20.4 --node-id $(emintd tendermint show-node-id --home testnet/node2) --keyring-backend test

# 6. collect-gentxs to genesis.json
cp testnet/node1/config/gentx/* testnet/node0/config/gentx/
cp testnet/node2/config/gentx/* testnet/node0/config/gentx/
emintd collect-gentxs --home testnet/node0
emintd collect-gentxs --home testnet/node1
emintd collect-gentxs --home testnet/node2

# 7. collect node1 and node2 genesis.json gentxs, copy to node0 genesis.json gentxs, copy node0 genesis.json to replace others
rm -f testnet/node1/config/genesis.json
rm -f testnet/node2/config/genesis.json
cp testnet/node0/config/genesis.json testnet/node1/config/
cp testnet/node0/config/genesis.json testnet/node2/config/

# 8. config each node's config.toml persistent_peers to the other two node's node-id@node-ip:26656
os=`uname -a`
mac='Darwin'
peers=`emintd tendermint show-node-id --home testnet/node0`@192.168.20.2:26656,`emintd tendermint show-node-id --home testnet/node1`@192.168.20.3:26656,`emintd tendermint show-node-id --home testnet/node2`@192.168.20.4:26656
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
fi


# 9. start each node, emintd start --home node* --rpc.unsafe --log_level "main:info,state:info,mempool:info"
echo -e "\n------Enjoy it!------"
