GOCMD=go
GOBUILD=$(GOCMD) build

.PHONY: all

all: keys sign verify

verify:
	$(GOBUILD) -o build/verify.exe cmd/verify/main.go

sign:
	$(GOBUILD) -o build/sign.exe cmd/sign/main.go

keys:
	$(GOBUILD) -o build/keys.exe cmd/keys/main.go

clean:
	rm -rf build/*.exe