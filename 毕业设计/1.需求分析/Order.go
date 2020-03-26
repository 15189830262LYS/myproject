 orderer.kevin.kongyixueyuan.com:
   image: hyperledger/fabric-orderer
   container_name: orderer.kevin.kongyixueyuan.com
   environment:
     - ORDERER_GENERAL_LOGLEVEL=debug
     - ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
     - ORDERER_GENERAL_LISTENPORT=7050
     - ORDERER_GENERAL_GENESISPROFILE=kongyixueyuan
     - ORDERER_GENERAL_GENESISMETHOD=file
     - ORDERER_GENERAL_GENESISFILE=/var/hyperledger/orderer/genesis.block
     - ORDERER_GENERAL_LOCALMSPID=kevin.kongyixueyuan.com
     - ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp
     - ORDERER_GENERAL_TLS_ENABLED=true
     - ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key
     - ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt
     - ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
   working_dir: /opt/gopath/src/github.com/hyperledger/fabric
   command: orderer
   volumes:
     - ./artifacts/genesis.block:/var/hyperledger/orderer/genesis.block
     - ./crypto-config/ordererOrganizations/kevin.kongyixueyuan.com/orderers/orderer.kevin.kongyixueyuan.com/msp:/var/hyperledger/orderer/msp
     - ./crypto-config/ordererOrganizations/kevin.kongyixueyuan.com/orderers/orderer.kevin.kongyixueyuan.com/tls:/var/hyperledger/orderer/tls
   ports:
     - 7050:7050
   networks:
     default:
       aliases:
         - orderer.kevin.kongyixueyuan.com