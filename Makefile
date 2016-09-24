# swhsiang swh@hsiang.io

all: idl/computing.thrift glide.yaml
	# Install Go package
	glide install;
	# Generated Thrift package.
	thrift --gen go -out vendor -r -I idl idl/computing.thrift;
	# Build Client
	cd src/cmd && go build -o ../bin/thrift-ex-client
	# Build Server
	cd src/server && go build -o ../bin/thrift-ex-server


test:
	# Test Client
	cd src/cmd && go test -timeout 60s -v ./... ;
	# Test Server
	cd src/server && go test -timeout 60s -v ./... ;

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
	bin/thrift-ex-server -secure &
	sleep 1 && bin/thrift-ex-client -secure
	sleep 1 && pkill -9 thrift-ex-server
