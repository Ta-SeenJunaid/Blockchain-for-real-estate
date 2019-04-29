#15. Applications interact with the blockchain ledger through chaincode. As such we need to install the chaincode on every peer that will execute and endorse our transactions, and then instantiate the chaincode on the channel.

# this installs the Go chaincode. For go chaincode -p takes the relative path from $GOPATH/src
#peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/flatfinal1/


#16. Modify the following four environment variables to issue the install command against peer0 in Org2:

export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=peer0.org2.example.com:7051
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

#17. Now install the sample Go, Node.js or Java chaincode onto a peer0 in Org2. 

#peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
peer chaincode install -n mycc -v 1.0 -p  github.com/chaincode/flatfinal1/

#16.1 Modify the following four environment variables to issue the install command against peer0 in Org3:

export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
export CORE_PEER_ADDRESS=peer0.org3.example.com:7051
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt

#17.1 Now install the sample Go, Node.js or Java chaincode onto a peer0 in Org3. 

#peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/
peer chaincode install -n mycc -v 1.0 -p  github.com/chaincode/flatfinal1/
