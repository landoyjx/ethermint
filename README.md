
## Introduction

Halle Chain is a scalable, high-throughput Proof-of-Stake blockchain that is fully compatible and
interoperable with Ethereum. It's build using the the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.


> **WARNING:** 1hale = 1uhale*10^6


Note: Requires Go 1.14+


## Useful links
- [Take a look at the Explorer](http://explorer.hallechain.cn/)


### Building halled

To build, execute the following commands:

```bash
# To build the project and install it in $GOBIN
make install

# To build the binary and put the resulting binary in ./build
make build
```

### Starting a Halled daemon (node)

The following config steps can be performed all at once by executing the `init.sh` file located in the root directory like this:
```bash
./init.sh
```
> This bash file removes previous blockchain data from `~/.halled` and `~/.hallecli`. It uses the `keyring-backend` called `test` that should prevent you from needing to enter a passkey. The `keyring-backend` `test` is unsecured and should not be used in production.



### Starting Halled Web3 RPC API

After the daemon is started, run (in another process):

```bash
hallecli  rest-server --laddr "tcp://localhost:8545" --unlock-key mykey
```

and to make sure the server has started correctly, try querying the current block number:

```bash
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545
```

or point any dev tooling at `http://localhost:8545` or whatever port is chosen just as you would with an Ethereum node

#### Clearing data from chain

Data for the CLI and Daemon should be stored at `~/.halled` and `~/.hallecli` by default, to start the node with a fresh state, run:

```bash
rm -rf ~/.halle*
```

To clear all data except key storage (if keyring backend chosen) and then you can rerun the commands to start the node again.

#### Keyring backend options

The instructions above include commands to use `test` as the `keyring-backend`. This is an unsecured keyring that doesn't require entering a password and should not be used in production. Otherwise, Halle supports using a file or OS keyring backend for key storage. To create and use a file stored key instead of defaulting to the OS keyring, add the flag `--keyring-backend file` to any relevant command and the password prompt will occur through the command line. This can also be saved as a CLI config option with:

```bash
hallecli config keyring-backend file
```

### Exporting Halle private key from Halle Chain

To export the private key from Halle to something like Metamask, run:

```bash
hallecli keys unsafe-export-eth-key mykey
```

Import account through private key, and to verify that the Ethereum address is correct with:

```bash
hallecli keys parse $(hallecli keys show  -a)
```
