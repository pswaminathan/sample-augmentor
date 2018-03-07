.PHONY: build vendor proto vendor_dep_ensure test artifact

all: build

build:
	go build $(FLAGS) -o augmentor

test:
	go test -v $(FLAGS) ./...

artifact:
	GOOS=linux GOARCH=amd64 go build $(FLAGS) -o augmentor_linux_amd64

proto:
	protoc -I=proto/ --gogofast_out=vendor proto/beeswax/base/*.proto
	protoc -I=proto/ --gogofast_out=vendor proto/beeswax/openrtb/*.proto
	protoc -I=proto/ --gogofast_out=vendor proto/beeswax/augment/*.proto

vendor: vendor_dep_ensure proto

vendor_dep_ensure:
	dep ensure
