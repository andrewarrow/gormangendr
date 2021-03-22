# gormangendr
golang cardano jormungandr (ADA)

./gormangendr node --genesis-block-hash=123

./gormangendr generate --proto-file=network/node.proto 

protoc --go_out="/Users/aa/" service.proto 
protoc --go-grpc_out="/Users/aa/" service.proto


./jcli key generate --type=Ed25519Extended > receiver_secret.key
cat receiver_secret.key | ./jcli key to-public > receiver_public.key

./jcli address account --testing --prefix addr $(cat receiver_public.key)

