build:
	go build -o bin/goway-sidecar cmd/agent/main.go

dep:
	glide install

docker:
	docker run --rm -v $(shell pwd)/bin:/go/src/github.com/andrepinto/goway-sidecar/bin $(shell docker build -f Dockerfile.build -q .) go build  -o bin/goway-sidecar cmd/agent/main.go
	docker build -f Dockerfile.dist -t andrepinto/goway-sidecar:1.0.3 .

.PHONY: build dep docker