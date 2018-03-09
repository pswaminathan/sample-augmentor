.PHONY: build vendor proto vendor_dep_ensure test artifact update_subtree

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

update_subtree:
	# This is an advanced maneuver.
	# I stored the proto definitions as a subtree so that we can periodically update them.
	# But I did some fancy stuff to make it apply a subdirectory.
	# Go read https://stackoverflow.com/questions/23937436/add-subdirectory-of-remote-repo-with-git-subtree
	# and decide if this is still the thing you want to do.
	git checkout proto_remote/master
	git pull proto_remote master
	git checkout master
	git rm -rf proto
	git read-tree --prefix=proto/beeswax -u proto_remote/master:beeswax
