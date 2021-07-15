CHAINID=tokenchain-1
DEMON=onp
DAEMON=tokend
HOME=node1
rm -rf $HOME
$DAEMON init node1 --chain-id $CHAINID --home $HOME
sed -i "" 's#stake#onp#g' $HOME/config/genesis.json
$DAEMON keys add validator --home $HOME
$DAEMON add-genesis-account validator "200000000000000 $DEMON" --home $HOME
$DAEMON gentx validator "20000000000$DEMON" --min-self-delegation "1000000" --chain-id $CHAINID --home $HOME
$DAEMON collect-gentxs --home $HOME
$DAEMON start --home $HOME