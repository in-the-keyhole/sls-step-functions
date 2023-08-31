.PHONY: build clean deploy gomodgen

GOFLAGS=-ldflags="-s -w"
GOOS=linux
ifeq ($(shell uname -m), aarch64)
GOARCH=arm64
else ifeq ($(shell uname -m), x86_64)
GOARCH=amd64
else
	@echo "Unsupported architecture. Please build manually."
	@exit 1
endif
GO_ENV=GOARCH=${GOARCH} GOOS=${GOOS} CGO_ENABLED=0

build: gomodgen
	export GO111MODULE=on
	go mod tidy
	env ${GO_ENV} go build ${GOFLAGS} -o bin/assembleMedia assembleMedia/main.go
	env ${GO_ENV} go build ${GOFLAGS} -o bin/distribution distribution/main.go

clean:
	rm -rf ./bin

offline: clean build
	sls offline --noTimeout

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
