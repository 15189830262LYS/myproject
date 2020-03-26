 ca.org1.kevin.kongyixueyuan.com:
   image: hyperledger/fabric-ca
   container_name: ca.org1.kevin.kongyixueyuan.com
   environment:
     - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
     - FABRIC_CA_SERVER_CA_NAME=ca.org1.kevin.kongyixueyuan.com
     - FABRIC_CA_SERVER_CA_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.org1.kevin.kongyixueyuan.com-cert.pem
     - FABRIC_CA_SERVER_CA_KEYFILE=/etc/hyperledger/fabric-ca-server-config/727e69ed4a01a204cd53bf4a97c2c1cb947419504f82851f6ae563c3c96dea3a_sk
     - FABRIC_CA_SERVER_TLS_ENABLED=true
     - FABRIC_CA_SERVER_TLS_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.org1.kevin.kongyixueyuan.com-cert.pem
     - FABRIC_CA_SERVER_TLS_KEYFILE=/etc/hyperledger/fabric-ca-server-config/727e69ed4a01a204cd53bf4a97c2c1cb947419504f82851f6ae563c3c96dea3a_sk
   ports:
     - 7054:7054
   command: sh -c 'fabric-ca-server start -b admin:adminpw -d'
   volumes:
     - ./crypto-config/peerOrganizations/org1.kevin.kongyixueyuan.com/ca/:/etc/hyperledger/fabric-ca-server-config
   networks:
     default:
       aliases:
         - ca.org1.kevin.kongyixueyuan.com