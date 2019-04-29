echo First let’s run the cryptogen tool.

../bin/cryptogen generate --config=./crypto-config.yaml

echo We need to tell theconfigtxgentool where to look for theconfigtx.yamlfile that it needs to ingest. We will tell it look in our present working directory:

export FABRIC_CFG_PATH=$PWD

echo Then, we’ll invoke the configtxgen tool to create the orderer genesis block:

../bin/configtxgen -profile TwoOrgsOrdererGenesis -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block

echo Next, we need to create the channel transaction artifact.
echo The channel.tx artifact contains the definitions for our sample channel

export CHANNEL_NAME=mychannel  && ../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME

echo Next, we will define the anchor peer for Org1 on the channel that we are constructing. 

../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP

echo Now, we will define the anchor peer for Org2 on the same channel:

../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP

echo Now, we will define the anchor peer for Org3 on the same channel:

../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org3MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org3MSP

echo First let’s start our network:

docker-compose -f docker-compose-cli.yaml up -d

echo Environment variables for PEER0

CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
CORE_PEER_ADDRESS=peer0.org1.example.com:7051
CORE_PEER_LOCALMSPID="Org1MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

echo We will enter the CLI container using the docker exec command:

docker exec -it cli bash