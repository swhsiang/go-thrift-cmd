# swhsiang swh@hsiang.io

all: idl/computing.thrift glide.yaml config
	# Build Server
	cd src/server && go build -o ../../bin/thrift-ex-server
	# Build Client
	cd src/cmd && go build -o ../../bin/thrift-ex-client

install:
	# Instasll glide
	curl https://glide.sh/get | sh

config:
	# Install Go package
	glide install;
	# Generated Thrift package.
	thrift --gen go -out vendor -r -I idl idl/computing.thrift;


test:
	# Test Client
	cd src/cmd && go test -timeout 60s -v ./... ;
	# Test Server
	cd src/server && go test -timeout 60s -v ./... ;

errcheck:
	go get -v -u github.com/kisielk/errcheck ;
	#------------------------
	# Err check
	#------------------------
	errcheck $$(go list ./... | grep -v /vendor/);

regen:
	rm -rf vendor/swhsiang/computing/;
	glide rebuild;
	thrift -version || true # thrift -version needs to be 1.0.0-dev !
	thrift --gen go -out vendor -r -I idl idl/tutorial.thrift;

clean:
	rm -f bin/* *~;
	rm -rf gen-go;
	rm -rf vendor/swhsiang/computing;

run:
	bin/thrift-ex-server &
	sleep 1 && bin/thrift-ex-client
