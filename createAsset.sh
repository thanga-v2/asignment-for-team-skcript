#!/bin/bash

# Exit on first error, print all commands.
set -ev
CHAINCODE_NAME="mycc"
CHANNEL_NAME="mychannel"
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

docker exec cli peer chaincode invoke -o orderer.example.com:7050  --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CHAINCODE_NAME -c '{"Args":["createAsset","DD1145", "The Blue V1 winston Diamond", "04-24-1988", "350000", "PENDING","","OW115","Thangaraj"]}'
    
