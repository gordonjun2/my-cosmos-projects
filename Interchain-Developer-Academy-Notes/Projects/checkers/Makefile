mock-expected-keepers:
	mockgen -source=x/checkers/types/expected_keepers.go \
		-package testutil \
		-destination=x/checkers/testutil/expected_keepers_mocks.go 

install-protoc-gen-ts:
	mkdir -p scripts/protoc
	cd scripts/protoc
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-osx-aarch_64.zip -o scripts/protoc/protoc.zip
	cd scripts/protoc && unzip -o protoc.zip
	rm scripts/protoc/protoc.zip
	sudo ln -s $(PWD)/scripts/protoc/bin/protoc /usr/local/bin/protoc
	cd scripts && npm install ts-proto@1.121.6 --save-dev --save-exact

cosmos-version = v0.45.4

download-cosmos-proto:
	mkdir -p proto/cosmos/base/query/v1beta1
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/proto/cosmos/base/query/v1beta1/pagination.proto -o proto/cosmos/base/query/v1beta1/pagination.proto
	mkdir -p proto/google/api
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/google/api/annotations.proto -o proto/google/api/annotations.proto
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/google/api/http.proto -o proto/google/api/http.proto
	mkdir -p proto/gogoproto
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/gogoproto/gogo.proto -o proto/gogoproto/gogo.proto

gen-protoc-ts: 
	mkdir -p ./client/src/types/generated/
	ls proto/checkers | xargs -I {} protoc \
		--plugin="./scripts/node_modules/.bin/protoc-gen-ts_proto" \
		--ts_proto_out="./client/src/types/generated" \
		--proto_path="./proto" \
		--ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
		checkers/{}
	
install-and-gen-protoc-ts: install-protoc-gen-ts download-cosmos-proto gen-protoc-ts

build-all:
    GOOS=linux GOARCH=amd64 go build -o ./build/myproject-linux-amd64 ./cmd/myprojectd/main.go
    GOOS=linux GOARCH=arm64 go build -o ./build/myproject-linux-arm64 ./cmd/myprojectd/main.go
    GOOS=darwin GOARCH=amd64 go build -o ./build/myproject-darwin-amd64 ./cmd/myprojectd/main.go

do-checksum:
    cd build && sha256sum myproject-linux-amd64 myproject-linux-arm64 myproject-darwin-amd64 > myproject_checksum

build-with-checksum: build-all do-checksum
