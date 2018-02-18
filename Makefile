all: bin/stellarctl.linux.amd64 \
	bin/stellarctl.darwin.amd64 \
	bin/stellarctl.linux.arm \
	bin/stellarctl.linux.arm64 \


SRC=$(shell find ./cmd ./transaction) main.go

vendor: Gopkg.toml $(SRC)
	dep ensure -v

bin/stellarctl.linux.amd64: $(SRC) vendor
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags '-extldflags "-static"' -o $@ .

bin/stellarctl.darwin.amd64: $(SRC) vendor
	mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -a -ldflags '-extldflags "-static"' -o $@ .

bin/stellarctl.linux.arm: $(SRC) vendor
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -v -a -ldflags '-extldflags "-static"' -o $@ .

bin/stellarctl.linux.arm64: $(SRC) vendor
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -a -ldflags '-extldflags "-static"' -o $@ .

clean:
	rm -rf bin
