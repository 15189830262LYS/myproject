 couchdb:
   container_name: couchdb
   image: hyperledger/fabric-couchdb
   # Populate the COUCHDB_USER and COUCHDB_PASSWORD to set an admin user and password
   # for CouchDB.  This will prevent CouchDB from operating in an "Admin Party" mode.
   environment:
     - COUCHDB_USER=
     - COUCHDB_PASSWORD=
   # Comment/Uncomment the port mapping if you want to hide/expose the CouchDB service,
   # for example map it to utilize Fauxton User Interface in dev environments.
   ports:
     - "5984:5984"