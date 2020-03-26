peer0.org1.kevin.kongyixueyuan.com:
  image: hyperledger/fabric-peer
  container_name: peer0.org1.kevin.kongyixueyuan.com
  environment:
    - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
    - CORE_VM_DOCKER_ATTACHSTDOUT=true
    - CORE_LOGGING_LEVEL=DEBUG
    - CORE_PEER_NETWORKID=kongyixueyuan
    - CORE_PEER_PROFILE_ENABLED=true
    - CORE_PEER_TLS_ENABLED=true
    - CORE_PEER_TLS_CERT_FILE=/var/hyperledger/tls/server.crt
    - CORE_PEER_TLS_KEY_FILE=/var/hyperledger/tls/server.key
    - CORE_PEER_TLS_ROOTCERT_FILE=/var/hyperledger/tls/ca.crt
    - CORE_PEER_ID=peer0.org1.kevin.kongyixueyuan.com
    - CORE_PEER_ADDRESSAUTODETECT=true
    - CORE_PEER_ADDRESS=peer0.org1.kevin.kongyixueyuan.com:7051
    - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.org1.kevin.kongyixueyuan.com:7051
    - CORE_PEER_GOSSIP_USELEADERELECTION=true
    - CORE_PEER_GOSSIP_ORGLEADER=false
    - CORE_PEER_GOSSIP_SKIPHANDSHAKE=true
    - CORE_PEER_LOCALMSPID=org1.kevin.kongyixueyuan.com
    - CORE_PEER_MSPCONFIGPATH=/var/hyperledger/msp
    - CORE_PEER_TLS_SERVERHOSTOVERRIDE=peer0.org1.kevin.kongyixueyuan.com
    - CORE_LEDGER_STATE_STATEDATABASE=CouchDB
    - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb:5984
    - CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=
    - CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD=
  working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
  command: peer node start
  volumes:
    - /var/run/:/host/var/run/
    - ./crypto-config/peerOrganizations/org1.kevin.kongyixueyuan.com/peers/peer0.org1.kevin.kongyixueyuan.com/msp:/var/hyperledger/msp
    - ./crypto-config/peerOrganizations/org1.kevin.kongyixueyuan.com/peers/peer0.org1.kevin.kongyixueyuan.com/tls:/var/hyperledger/tls
  ports:
    - 7051:7051
    - 7053:7053
  depends_on:
    - orderer.kevin.kongyixueyuan.com
    - couchdb
  networks:
    default:
      aliases:
        - peer0.org1.kevin.kongyixueyuan.com