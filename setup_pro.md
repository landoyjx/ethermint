

./halled init node0   --chain-id   0812
./halled init node1   --chain-id   0812
./halled init node2   --chain-id   0812
./halled init node3   --chain-id   0812
./halled init node4   --chain-id   0812
./halled init node5   --chain-id   0812
./halled init node6   --chain-id   0812



  sed -i 's/"max_gas": "-1"/"max_gas": "1000000000"/g'   ~/.halled/config/genesis.json



  ./hallecli config chain-id  0812
  ./hallecli config output json
  ./hallecli config indent true


node0
  ./hallecli keys add  node0
  ./hallecli keys add  escrow
  ./hallecli keys add  halleasset  
node1
  ./hallecli keys add  node1  
node2
  ./hallecli keys add  node2  
node3
  ./hallecli keys add  node3  
node4
  ./hallecli keys add  node4  
node5
  ./hallecli keys add  node5
node6
  ./hallecli keys add  node6  





node0
  ./halled add-genesis-account halle1qknv02r5ef20tw5qvuzegm7lclpmej2mwg3aer   20000000000000uhale
  ./halled add-genesis-account halle168ev4ypzzy5g5fnwcgrjyth8vm7kgqfhmvkwj0   20000000000000uhale
  ./halled add-genesis-account halle1ufxlanum8zcd6wsrmsjcrfgkq6j7qh70zrtasj   20000000000000uhale
  ./halled add-genesis-account halle1j0ex7yfm6p2glcdhm0zgx3zytxcfayu2hl28en   20000000000000uhale
  ./halled add-genesis-account halle1x0q76kx3g9r44v5q6wmtd0e60y5z6zcq7gf6fc   20000000000000uhale
  ./halled add-genesis-account halle1hkan7nw0gcsxsumzmh2gg2q0p5nszrf8u8zaul   20000000000000uhale
  ./halled add-genesis-account halle1uftxyygasnmf2gv4l2yyxze5smhu383ga6lj0f   20000000000000uhale

escrow:
  ./halled add-genesis-account  halle1fmwntc5aathr4ykj35q7nyutr68arq36j6h220  200000000000000uhale
halleasset
1460000000
  ./halled add-genesis-account  halle1fyupz6yv4jq8y5x6mv0lrshc2m8cr3e9kpajgj  1460000000000000uhale


node1
  ./halled add-genesis-account halle168ev4ypzzy5g5fnwcgrjyth8vm7kgqfhmvkwj0   20000000000000uhale
node2
  ./halled add-genesis-account halle1ufxlanum8zcd6wsrmsjcrfgkq6j7qh70zrtasj   20000000000000uhale
node3
  ./halled add-genesis-account halle1j0ex7yfm6p2glcdhm0zgx3zytxcfayu2hl28en   20000000000000uhale
node4
  ./halled add-genesis-account halle1x0q76kx3g9r44v5q6wmtd0e60y5z6zcq7gf6fc   20000000000000uhale
node5
  ./halled add-genesis-account halle1hkan7nw0gcsxsumzmh2gg2q0p5nszrf8u8zaul   20000000000000uhale
node6
  ./halled add-genesis-account halle1uftxyygasnmf2gv4l2yyxze5smhu383ga6lj0f   20000000000000uhale





20000000000000uhale
1000000000000uhale
1000000000000uhale


node0
./halled gentx --name node0  --amount  1000000000000uhale    --ip 47.95.248.152

node1
./halled gentx --name node1  --amount  1000000000000uhale    --ip 39.106.43.205

node2
./halled gentx --name node2  --amount  1000000000000uhale    --ip  121.196.28.153

node3
./halled gentx --name node3  --amount  1000000000000uhale    --ip  47.108.133.148

node4
./halled gentx --name node4  --amount  1000000000000uhale    --ip  39.99.196.182

node5
./halled gentx --name node5  --amount  1000000000000uhale    --ip  120.24.108.48

node6
./halled gentx --name node6  --amount  1000000000000uhale    --ip 120.24.148.221




node0
./halled collect-gentxs
./halled validate-genesis



cp -r .halled/  .halled.origin
cp -r .hallecli/ .hallecli.orgin




p2p: 26656
rest-server: 8545
docker swarm:  2377
docker ui: 9000


start or update Command
```bash
docker stack deploy -c docker-compose.pro.yml  0812  
```



mv .halled  .halled.chainIdStr.ethTx.failed
mv .hallecli .hallecli.chainIdStr.ethTx.failed

cp -r .hallecli.orgin   .hallecli
cp -r .halled.origin    .halled


  sed -i 's/"chain_id": "8"/"chain_id":"halle-1"/g'   ~/.halled/config/genesis.json


sed -i 's/"chain_id": "halle-1"/"chain_id": "0812"/g'   ~/.halled/config/genesis.json
